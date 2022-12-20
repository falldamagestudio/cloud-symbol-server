package admin_api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/sql-db-models"
)

func (s *ApiService) GetTokens(ctx context.Context) (openapi.ImplResponse, error) {

	log.Printf("Fetching PATs")

	db := GetDB()
	if db == nil {
		log.Printf("No DB")
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	// TODO: fetch owner from auth
	owner := "hello"

	// fetch PAT
	pats, err := models.PersonalAccessTokens(
		qm.Where(models.PersonalAccessTokenTableColumns.Owner+" = ?", owner),
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

	getTokensResponse := openapi.GetTokensResponse{
		Items: tokensArray,
	}

	return openapi.Response(http.StatusOK, getTokensResponse), nil
}
