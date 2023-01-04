package token_api

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/volatiletech/sqlboiler/v4/boil"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/sql-db-models"
	helpers "github.com/falldamagestudio/cloud-symbol-server/admin-api/helpers"
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func CreateToken(ctx context.Context) (openapi.ImplResponse, error) {

	log.Printf("Creating PAT")

	db := helpers.GetDB()
	if db == nil {
		log.Printf("No DB")
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	owner := helpers.GetUserIdentity(ctx)

	tokenLength := 32

	token, err := randomHex(tokenLength)
	if err != nil {
		log.Printf("Error while generating new token: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), nil
	}

	pat := models.PersonalAccessToken{
		Owner:             owner,
		Token:             token,
		Description:       "",
		CreationTimestamp: time.Now(),
	}

	if err := pat.Insert(ctx, db, boil.Infer()); err != nil {

		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			if pgError.Code == pgerrcode.UniqueViolation {
				log.Printf("PAT %v already exists; err = %v", token, err)
				return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Duplicate PAT generated")}), nil
			}
		}

		log.Printf("Create PAT error: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	log.Printf("PAT created OK")
	createTokenResponse := openapi.CreateTokenResponse{
		Token: token,
	}
	return openapi.Response(http.StatusOK, createTokenResponse), nil
}
