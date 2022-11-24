package admin_api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
)

func (s *ApiService) CreateStore(ctx context.Context, storeId string) (openapi.ImplResponse, error) {

	log.Printf("Creating store")

	if err := sqlCreateStore(ctx, storeId); err != nil {

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
