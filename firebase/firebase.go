package firebase

import (
	"context"
	"fmt"
	"log"

	fbase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// SetupFirebase sets up firebase client for authentication
func SetupFirebase() *auth.Client {

	opt := option.WithCredentialsFile("myproject.json")
	//Firebase admin SDK initialization
	app, err := fbase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error %v\n", err)
	}

	//Firebase Auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error %v\n", err)
	}

	return auth
}

// GetPasswordResetLink will return passowrd reset link
func GetPasswordResetLink(email string, baseURL string, roles string) (string, error) {
	var actionCodeSettings *auth.ActionCodeSettings

	if baseURL != "" {
		actionCodeSettings = &auth.ActionCodeSettings{
			URL: baseURL + "/auth/login?app=" + roles,
		}
	}

	f := SetupFirebase()

	link, err := f.PasswordResetLinkWithSettings(context.Background(), email, actionCodeSettings)
	if err != nil {
		fmt.Println("46 ", err)
		return "", err
	}

	link = link + "&email=" + email

	return link, nil
}
