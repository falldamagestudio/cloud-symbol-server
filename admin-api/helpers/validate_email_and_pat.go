package helpers

import (
	"context"
	"errors"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/sql-db-models"
)

func ValidateEmailAndPAT(ctx context.Context, email string, pat string) (bool, error) {
	db := GetDB()
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