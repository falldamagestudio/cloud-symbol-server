package admin_api

import (
	"encoding/json"
	"log"
	"net/http"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	helpers "github.com/falldamagestudio/cloud-symbol-server/admin-api/helpers"
)

type patAuthenticationMiddleware struct{}

func writePatHttpError(w http.ResponseWriter, status int, message string) error {

	messageResponse := openapi.MessageResponse{Message: message}
	messageJsonBytes, err := json.Marshal(messageResponse)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(messageJsonBytes)
	return nil
}

func validateUsernamePassword(r *http.Request) (bool, int) {

	// Fetch email + PAT from Basic Authentication header of WWW request

	email, pat, basicAuthPresent := r.BasicAuth()

	if !basicAuthPresent {

		log.Print("Basic auth header (with email/token) not provided")
		return false, 0
	}

	log.Printf("Basic Auth header present, email: %v, PAT: %v", email, pat)

	// Validate that email + PAT exists in database

	firestoreClient, err := helpers.FirestoreClient(r.Context())
	if err != nil {
		log.Printf("Unable to talk to database: %v", err)
		return false, http.StatusInternalServerError
	}

	patDocRef := firestoreClient.Collection("users").Doc(email).Collection("pats").Doc(pat)

	if _, err := patDocRef.Get(r.Context()); err != nil {

		log.Printf("Unable to find email %v, pat %v combination in database: %v", email, pat, err)
		return false, http.StatusUnauthorized
	}

	log.Printf("Email/PAT pair %v, %v exist in database; authentication successful", email, pat)
	return true, 0
}

func (patAM *patAuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Authenticate via username + password combination

		authenticated, statusCode := validateUsernamePassword(r)
		if statusCode != 0 {

			if statusCode == http.StatusUnauthorized {
				w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
				_ = writePatHttpError(w, http.StatusUnauthorized, "Unauthorized; please provide valid email + personal access token when using Basic Authentication")
				return
			} else {
				_ = writePatHttpError(w, http.StatusInternalServerError, "Internal server error")
				return
			}
		}

		if !authenticated {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			_ = writePatHttpError(w, http.StatusUnauthorized, "Unauthorized; please provide valid email + personal access token")
			return
		}

		next.ServeHTTP(w, r)
	})
}
