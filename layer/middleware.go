package layer

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// BasicAuthHandler verifies auth or basic auth credentials
func BasicAuthHandler(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		const basicAuthPrefix string = "Basic "

		// Get the value in header
		auth := r.Header.Get("Authorization")

		if strings.HasPrefix(auth, basicAuthPrefix) {
			user := []byte(os.Getenv("BASICAUTH_USERNAME"))
			pass := []byte(os.Getenv("BASICAUTH_PASSWORD"))

			// Check credentials
			payload, err := base64.StdEncoding.DecodeString(auth[len(basicAuthPrefix):])
			if err == nil {
				pair := bytes.SplitN(payload, []byte(":"), 2)
				if len(pair) == 2 && bytes.Equal(pair[0], user) && bytes.Equal(pair[1], pass) {
					// Delegate request to the given handle
					next(w, r, ps)
					return
				}
			}
		}
		// Request Basic Authentication otherwise
		w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

	}
}
