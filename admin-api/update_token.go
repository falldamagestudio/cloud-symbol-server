package admin_api

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/sql-db-models"
)

func (s *ApiService) UpdateToken(ctx context.Context, token string, updateTokenRequest openapi.UpdateTokenRequest) (openapi.ImplResponse, error) {

	log.Printf("Updating PAT %v", token)

	db := GetDB()
	if db == nil {
		log.Printf("No DB")
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	// TODO: fetch owner from auth
	owner := "hello"

	// update PAT
	numRowsAffected, err := models.PersonalAccessTokens(
		qm.Where(models.PersonalAccessTokenTableColumns.Owner+" = ?", owner),
		qm.And(models.PersonalAccessTokenTableColumns.Token+" = ?", token),
	).UpdateAll(ctx, db, models.M{
		models.PersonalAccessTokenColumns.Description: updateTokenRequest.Description,
	})

	if (err == nil) && (numRowsAffected == 0) {
		log.Printf("PAT %v not found for owner %v", token, owner)
		return openapi.Response(http.StatusInternalServerError, nil), err
	} else if (err != nil) || (numRowsAffected != 1) {
		log.Printf("Error while updating token %v for owner %v: %v", token, owner, err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	log.Printf("PAT %v updated for owner %v", token, owner)

	return openapi.Response(http.StatusOK, nil), nil
}
