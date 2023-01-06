package token_api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/sql-db-models"
	helpers "github.com/falldamagestudio/cloud-symbol-server/backend-api/helpers"
	postgres "github.com/falldamagestudio/cloud-symbol-server/backend-api/postgres"
)

func GetToken(ctx context.Context, token string) (openapi.ImplResponse, error) {

	log.Printf("Fetching PAT %v", token)

	db := postgres.GetDB()
	if db == nil {
		log.Printf("No DB")
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	owner := helpers.GetUserIdentity(ctx)

	// fetch PAT
	pat, err := models.PersonalAccessTokens(
		qm.Where(models.PersonalAccessTokenTableColumns.Owner+" = ?", owner),
		qm.And(models.PersonalAccessTokenTableColumns.Token+" = ?", token),
	).One(ctx, db)

	if err != nil {
		log.Printf("Error while fetching token %v for owner %v: %v", token, owner, err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	getTokenResponse := openapi.GetTokenResponse{
		Token:             pat.Token,
		Description:       pat.Description,
		CreationTimestamp: pat.CreationTimestamp.Format(time.RFC3339),
	}

	return openapi.Response(http.StatusOK, getTokenResponse), nil
}