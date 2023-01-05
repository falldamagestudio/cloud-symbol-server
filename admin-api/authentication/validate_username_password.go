package authentication

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/sql-db-models"
	postgres "github.com/falldamagestudio/cloud-symbol-server/admin-api/postgres"
)

type UsernamePasswordValidator struct {
}

func CreateUsernamePasswordValidator() (*UsernamePasswordValidator) {
	return &UsernamePasswordValidator{}
}

func (*UsernamePasswordValidator) Validate(r *http.Request) (string, int, string) {

	// Fetch email + PAT from Basic Authentication header of WWW request

	email, pat, basicAuthPresent := r.BasicAuth()

	if !basicAuthPresent {

		log.Print("Basic auth header (with email/token) not provided")
		return "", 0, ""
	}

	log.Printf("Basic Auth header present, email: %v, PAT: %v", email, pat)

	// Validate that email + PAT exists in database

	success, err := doesOwnerAndPatExistInDB(r.Context(), email, pat)

	if err != nil {
		log.Printf("Error while looking up email/PAT pair %v, %v in database: %v", email, pat, err)
		return "", http.StatusInternalServerError, ""
	}

	if !success {
		log.Printf("Email/PAT pair %v, %v does not exist in database", email, pat)
		return "", http.StatusUnauthorized, "Unauthorized; please provide valid email + personal access token"
	} else {

		// email + PAT are valid

		log.Printf("Email/PAT pair %v, %v exists in database; authentication successful", email, pat)
		return email, 0, ""
	}
}

func doesOwnerAndPatExistInDB(ctx context.Context, email string, pat string) (bool, error) {
	db := postgres.GetDB()
	if db == nil {
		return false, errors.New("no DB")
	}

	numHits, err := models.PersonalAccessTokens(
		qm.Where(models.PersonalAccessTokenColumns.Owner+" = ?", email),
		qm.And(models.PersonalAccessTokenColumns.Token+" = ?", pat),
	).Count(ctx, db)

	if err != nil {
		return false, err
	} else {
		return numHits != 0, nil
	}
}