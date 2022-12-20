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

func (s *ApiService) DeleteToken(ctx context.Context, token string) (openapi.ImplResponse, error) {

	log.Printf("Deleting PAT %v", token)

	tx, err := BeginDBTransaction(ctx)
	if err != nil {
		log.Printf("Err: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	// TODO: fetch owner from auth
	owner := "hello"

	// Delete PAT
	numRowsAffected, err := models.PersonalAccessTokens(
		qm.Where(models.PersonalAccessTokenTableColumns.Owner+" = ?", owner),
		qm.And(models.PersonalAccessTokenTableColumns.Token+" = ?", token),
	).DeleteAll(ctx, tx)

	if (err == nil) && (numRowsAffected == 0) {
		log.Printf("Token %v not found for owner %v", token, owner)
		tx.Rollback()
		return openapi.Response(http.StatusNotFound, nil), err
	} else if err != nil {
		log.Printf("Error while deleting Token %v for owner %v: %v", token, owner, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("error while committing txn: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}
