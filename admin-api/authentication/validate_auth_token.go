package authentication

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

type AuthTokenValidator struct {
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

func CreateAuthTokenValidator(clientID string) (*AuthTokenValidator, error) {

	ar := jwk.NewAutoRefresh(context.Background())
	ar.Configure(jwksEndpoint, jwk.WithMinRefreshInterval(jwksRefreshIntervalMinutes*time.Minute))
	ks, err := ar.Refresh(context.Background(), jwksEndpoint)
	if err != nil {
		return nil, err
	}

	authTokenValidator := &AuthTokenValidator{
		Ks:       ks,
		ClientID: clientID,
	}
	return authTokenValidator, nil	
}

func (authTokenValidator *AuthTokenValidator) Validate(r *http.Request) (string, int, string) {

	// Check for token in Authorization header in WWW request

	if authorizationHeaderValues, ok := r.Header["Authorization"]; !ok || (len(authorizationHeaderValues) != 1) || !strings.HasPrefix(authorizationHeaderValues[0], "Bearer ") {
		log.Print("No JWT auth token present")
		return "", 0, ""
	} else {

		// Fetch JWT, and validate that it originates from Google

		token, err := jwt.ParseRequest(r, jwt.WithKeySet(authTokenValidator.Ks))
		if err != nil {
			log.Printf("Error when parsing JWT: %v", err)
			return "", 0, ""
		}

		// Validate that JWT is intended for this particular GCP project

		if err := jwt.Validate(token, jwt.WithAudience(authTokenValidator.ClientID)); err != nil {
			log.Printf("JWT fails validation: %v", err)
			return "", http.StatusUnauthorized, "Unauthorized; please provide valid OIDC token when using OIDC authentication"
		}
	
		// JWT is valid

		email := token.PrivateClaims()["email"].(string)
	
		log.Printf("JWT auth token for %v validated", email)
	
		return email, 0, ""
	}
}

