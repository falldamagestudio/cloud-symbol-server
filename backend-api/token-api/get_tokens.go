package token_api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/sql-db-models"
	helpers "github.com/falldamagestudio/cloud-symbol-server/backend-api/helpers"
	postgres "github.com/falldamagestudio/cloud-symbol-server/backend-api/postgres"
)

var getTokensSortOptions = map[string]string{
	"":                   models.PersonalAccessTokenColumns.CreationTimestamp,
	"token":              models.PersonalAccessTokenColumns.Token,
	"-token":             models.PersonalAccessTokenColumns.Token + " desc",
	"description":        models.PersonalAccessTokenColumns.Description,
	"-description":       models.PersonalAccessTokenColumns.Description + " desc",
	"creationTimestamp":  models.PersonalAccessTokenColumns.CreationTimestamp,
	"-creationTimestamp": models.PersonalAccessTokenColumns.CreationTimestamp + " desc",
}

func GetTokens(ctx context.Context, sort string) (openapi.ImplResponse, error) {

	log.Printf("Fetching PATs, sort %v", sort)

	// Handle sorting
	orderByOption := ""
	if sortOption, ok := getTokensSortOptions[sort]; ok {
		orderByOption = sortOption
	} else {
		log.Printf("Unsupported sort option: %v", sort)
		return openapi.Response(http.StatusBadRequest, openapi.MessageResponse{Message: fmt.Sprintf("Unsupported sort option: %v", sort)}), nil
	}

	db := postgres.GetDB()
	if db == nil {
		log.Printf("No DB")
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	owner := helpers.GetUserIdentity(ctx)

	// fetch PAT
	pats, err := models.PersonalAccessTokens(
		qm.Where(models.PersonalAccessTokenTableColumns.Owner+" = ?", owner),
		qm.OrderBy(orderByOption),
	).All(ctx, db)

	if err != nil {
		log.Printf("Error while fetching tokens for owner %v: %v", owner, err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	tokensArray := make([]openapi.GetTokenResponse, len(pats))
	for tokenIndex, pat := range pats {
		tokensArray[tokenIndex] = openapi.GetTokenResponse{
			Token:             pat.Token,
			Description:       pat.Description,
			CreationTimestamp: pat.CreationTimestamp.Format(time.RFC3339),
		}
	}

	return openapi.Response(http.StatusOK, tokensArray), nil
}
