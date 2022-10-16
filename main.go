package main

import (
	"context"
	"encoding/json"
	"fmt"
	fb "gocrud/firebase"
	"gocrud/json/body"
	"gocrud/layer"
	"gocrud/storage"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	r.POST("/user", layer.BasicAuthHandler(UserSignupHandler))
	r.POST("/resetPassword", layer.BasicAuthHandler(resetPasswordHandler))

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}

func UserSignupHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {

	defer r.Body.Close()
	rw.Header().Add("X-Request-ID", r.Header.Get("X-Request-ID"))
	rw.Header().Add("Content-Type", "application/json")

	var request body.UserRequestData
	var response body.UserResponseData

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(reqBody, &request)

	if err != nil {
		http.Error(rw, "Unable to unmarshal request", http.StatusInternalServerError)
		return
	}

	dao := request.MapObjects()

	err = storage.Save(&dao)
	if err != nil {
		http.Error(rw, "Unable to save user in database", http.StatusInternalServerError)
		return
	}

	response.MapResponse(&dao)

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(response)
}

func resetPasswordHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rw.Header().Add("X-Request-ID", r.Header.Get("X-Request-ID"))
	rw.Header().Add("Content-Type", "application/json")

	defer r.Body.Close()

	var request body.PasswordRequestData
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		http.Error(rw, "Unable to unmarshal request", http.StatusInternalServerError)
		return
	}

	fbClient := fb.SetupFirebase()

	ctx := context.Background()

	user, err := fbClient.GetUserByEmail(ctx, request.Data.Email)

	if err != nil {
		fmt.Println(err)
		http.Error(rw, "no user exists with the email "+request.Data.Email, http.StatusNotFound)
		return
	}
	domain := "example.com"
	_, err = url.ParseRequestURI(domain)
	if err != nil {
		domain = "https://" + domain
	}

	resetPasswordLink, err := fb.GetPasswordResetLink(user.Email, domain, "user")

	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(resetPasswordLink)
}
