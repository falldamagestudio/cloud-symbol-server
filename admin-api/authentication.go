package admin_api

import (
	"encoding/json"
	"log"
	"net/http"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
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

func (patAM *patAuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Fetch email + PAT from Basic Authentication header of WWW request

		email, pat, basicAuthPresent := r.BasicAuth()

		if !basicAuthPresent {

			log.Print("Basic auth header (with email/token) not provided")

			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			_ = writePatHttpError(w, http.StatusUnauthorized, "Unauthorized; please provide email + personal access token using Basic Authentication")
			return
		}

		log.Printf("Basic Auth header present, email: %v, PAT: %v", email, pat)

		// Validate that email + PAT exists in database

		firestoreClient, err := firestoreClient(r.Context())
		if err != nil {
			log.Printf("Unable to talk to database: %v", err)
			_ = writePatHttpError(w, http.StatusInternalServerError, "Unable to talk to database")
			return
		}

		patDocRef := firestoreClient.Collection("users").Doc(email).Collection("pats").Doc(pat)

		if _, err := patDocRef.Get(r.Context()); err != nil {

			log.Printf("Unable to find email %v, pat %v combination in database: %v", email, pat, err)

			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			_ = writePatHttpError(w, http.StatusUnauthorized, "Unauthorized; unable to find email / pat combination in database")
			return
		}

		log.Printf("Email/PAT pair %v, %v exist in database; authentication successful", email, pat)

		next.ServeHTTP(w, r)
	})
}
