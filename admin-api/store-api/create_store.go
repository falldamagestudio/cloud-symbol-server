package store_api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/volatiletech/sqlboiler/v4/boil"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/sql-db-models"
	postgres "github.com/falldamagestudio/cloud-symbol-server/admin-api/postgres"
)

func CreateStore(ctx context.Context, storeId string) (openapi.ImplResponse, error) {

	log.Printf("Creating store")

	db := postgres.GetDB()
	if db == nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	var store = models.Store{
		Name: storeId,
	}
	if err := store.Insert(ctx, db, boil.Infer()); err != nil {

		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			if pgError.Code == pgerrcode.UniqueViolation {
				log.Printf("Store %v already exists; err = %v", storeId, err)
				return openapi.Response(http.StatusConflict, openapi.MessageResponse{Message: fmt.Sprintf("Store %v already exists", storeId)}), err
			}
		}

		log.Printf("Create store error: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	log.Printf("Store created OK")
	return openapi.Response(http.StatusOK, nil), nil
}
