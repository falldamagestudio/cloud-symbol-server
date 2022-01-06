package upload_api

import (
	"log"
	"net/http"
)

type patAuthenticationMiddleware struct{}

func (patAM *patAuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Fetch email + PAT from Basic Authentication header of WWW request

		email, pat, basicAuthPresent := r.BasicAuth()

		if !basicAuthPresent {

			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			log.Print("Basic auth header (with email/token) not provided")
			http.Error(w, "Unauthorized; please provide email + personal access token using Basic Authentication", http.StatusUnauthorized)
			return
		}

		log.Printf("Basic Auth header present, email: %v, PAT: %v", email, pat)

		// Validate that email + PAT exists in database

		firestoreClient, err := firestoreClient(r.Context())
		if err != nil {
			log.Printf("Unable to talk to database: %v", err)
			http.Error(w, "Unable to talk to database", http.StatusInternalServerError)
			return
		}

		patDocRef := firestoreClient.Collection("users").Doc(email).Collection("pats").Doc(pat)

		if _, err := patDocRef.Get(r.Context()); err != nil {

			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			log.Printf("Unable to find email %v, pat %v combination in database: %v", email, pat, err)
			http.Error(w, "Unauthorized; unable to find email / pat combination in database", http.StatusUnauthorized)
			return
		}

		log.Printf("Email/PAT pair %v, %v exist in database; authentication successful", email, pat)

		next.ServeHTTP(w, r)
	})
}
