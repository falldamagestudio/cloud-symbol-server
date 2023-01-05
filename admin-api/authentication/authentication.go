package authentication

import (
	"encoding/json"
	"log"
	"net/http"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	helpers "github.com/falldamagestudio/cloud-symbol-server/admin-api/helpers"
)

type Validator interface {
	Validate(r *http.Request) (userIdentity string, statusCode int, errorMessage string)
}

type AuthenticationMiddleware struct {
	validators []Validator
}

func CreateAuthenticationMiddleware(validators []Validator) (*AuthenticationMiddleware) {

	authenticationMiddleware := &AuthenticationMiddleware{
		validators: validators,
	}
	return authenticationMiddleware
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

func (authenticationMiddleware *AuthenticationMiddleware) Handler(next http.Handler) http.Handler {

	log.Printf("Validating auth")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userIdentity := ""

		// Try validators, one by one, until there is a valid identity, an invalid identity, or we have tried all validators

		for _, validator := range(authenticationMiddleware.validators) {
			userIdentity2, statusCode, errorMessage := validator.Validate(r)
			if statusCode != 0 {
				if statusCode == http.StatusUnauthorized {
					w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
					_ = writeAuthenticationHttpError(w, http.StatusUnauthorized, errorMessage)
				} else {
					_ = writeAuthenticationHttpError(w, http.StatusInternalServerError, "Internal server error")
				}
				return
			}

			if userIdentity2 != "" {
				// We have found a valid identity; do not try subsequent validators
				userIdentity = userIdentity2
				break
			}
		}

		if userIdentity == "" {

			// No valid authentication found

			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			_ = writeAuthenticationHttpError(w, http.StatusUnauthorized, "Unauthorized; please provide some form of credentials")
			return
		}

		// Chain to following handlers, and include user identity in context

		ctxWithUserIdentity := helpers.AddUserIdentity(r.Context(), userIdentity)

		next.ServeHTTP(w, r.WithContext(ctxWithUserIdentity))
	})
}
