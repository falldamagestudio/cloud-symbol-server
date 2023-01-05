package admin_api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	helpers "github.com/falldamagestudio/cloud-symbol-server/admin-api/helpers"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

type AuthenticationMiddleware struct {
	// JSON Web Key Set used to validate authenticity of OpenID Connect ID Tokens
	Ks jwk.Set

	// Expected Client ID when authenticating with OpenID Connect ID Token
	ClientID string
}

const (
	// We expect OpenID Connect ID Tokens from Firebase's auth flow
	//
	// We use an undocumented API endpoint to retrieve the JWKS
	// since the official endpoint (at https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com)
	// requires to be converted to JWKS format for use with the jwx toolkit
	//
	// Reference: https://stackoverflow.com/a/71988314
	jwksEndpoint = "https://www.googleapis.com/service_accounts/v1/jwk/securetoken@system.gserviceaccount.com"

	jwksRefreshIntervalMinutes = 15
)

func createAuthenticationMiddleware(clientID string) (*AuthenticationMiddleware, error) {

	ar := jwk.NewAutoRefresh(context.Background())
	ar.Configure(jwksEndpoint, jwk.WithMinRefreshInterval(jwksRefreshIntervalMinutes*time.Minute))
	ks, err := ar.Refresh(context.Background(), jwksEndpoint)
	if err != nil {
		fmt.Printf("failed to refresh JWKS: %v\n", err)
		return nil, err
	}
	authenticationMiddleware := &AuthenticationMiddleware{
		Ks:       ks,
		ClientID: clientID,
	}
	return authenticationMiddleware, nil
}

func writeAuthenticationHttpError(w http.ResponseWriter, status int, message string) error {

	messageResponse := openapi.MessageResponse{Message: message}
	messageJsonBytes, err := json.Marshal(messageResponse)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(messageJsonBytes)
	log.Print(message)
	return nil
}

func validateUsernamePassword(r *http.Request) (string, int) {

	// Fetch email + PAT from Basic Authentication header of WWW request

	email, pat, basicAuthPresent := r.BasicAuth()

	if !basicAuthPresent {

		log.Print("Basic auth header (with email/token) not provided")
		return "", 0
	}

	log.Printf("Basic Auth header present, email: %v, PAT: %v", email, pat)

	// Validate that email + PAT exists in database

	success, err := helpers.ValidateEmailAndPAT(r.Context(), email, pat)

	if err != nil {
		log.Printf("Error while looking up email/PAT pair %v, %v in database: %v", email, pat, err)
		return "", http.StatusInternalServerError
	}

	if !success {
		log.Printf("Email/PAT pair %v, %v does not exist in database", email, pat)
		return "", 0
	} else {

		// email + PAT are valid

		log.Printf("Email/PAT pair %v, %v exists in database; authentication successful", email, pat)
		return email, 0
	}
}

func (authenticationMiddleware *AuthenticationMiddleware) validateAuthToken(r *http.Request) (string, int) {

	// Check for token in Authorization header in WWW request

	if authorizationHeaderValues, ok := r.Header["Authorization"]; !ok || (len(authorizationHeaderValues) != 1) || !strings.HasPrefix(authorizationHeaderValues[0], "Bearer ") {
		log.Print("No JWT auth token present")
		return "", 0
	} else {

		// Fetch JWT, and validate that it originates from Google

		token, err := jwt.ParseRequest(r, jwt.WithKeySet(authenticationMiddleware.Ks))
		if err != nil {
			log.Printf("Error when parsing JWT: %v", err)
			return "", 0
		}

		// Validate that JWT is intended for this particular GCP project

		if err := jwt.Validate(token, jwt.WithAudience(authenticationMiddleware.ClientID)); err != nil {
			log.Printf("JWT fails validation: %v", err)
			return "", http.StatusUnauthorized
		}
	
		// JWT is valid

		email := token.PrivateClaims()["email"].(string)
	
		log.Printf("JWT auth token for %v validated", email)
	
		return email, 0
	}
}

func (authenticationMiddleware *AuthenticationMiddleware) Handler(next http.Handler) http.Handler {

	log.Printf("Validating auth")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Authenticate via username + password combination

		userIdentity, statusCode := validateUsernamePassword(r)
		if statusCode != 0 {

			if statusCode == http.StatusUnauthorized {
				w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
				_ = writeAuthenticationHttpError(w, http.StatusUnauthorized, "Unauthorized; please provide valid email + personal access token when using Basic Authentication")
				return
			} else {
				_ = writeAuthenticationHttpError(w, http.StatusInternalServerError, "Internal server error")
				return
			}
		}

		if userIdentity == "" {

			// Authenticate via Bearer token

			userIdentity, statusCode = authenticationMiddleware.validateAuthToken(r)
			if statusCode != 0 {

				if statusCode == http.StatusUnauthorized {
					w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
					_ = writeAuthenticationHttpError(w, http.StatusUnauthorized, "Unauthorized; please provide valid OIDC token when using OIDC authentication")
					return
				} else {
					_ = writeAuthenticationHttpError(w, http.StatusInternalServerError, "Internal server error")
					return
				}
			}
		}

		if userIdentity == "" {

			// No valid authentication found

			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			_ = writeAuthenticationHttpError(w, http.StatusUnauthorized, "Unauthorized; please provide valid email + personal access token, or a valid OIDC token")
			return
		}

		// Chain to following handlers, and include user identity in context

		ctxWithUserIdentity := helpers.AddUserIdentity(r.Context(), userIdentity)

		next.ServeHTTP(w, r.WithContext(ctxWithUserIdentity))
	})
}
