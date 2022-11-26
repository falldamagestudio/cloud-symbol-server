package admin_api

import (
	"context"
	"database/sql"
	"log"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/sql-db-models"
)

var db *sql.DB

func initSQL() {

	log.Printf("Registering cloudsql-postgres driver")
	cleanup, err := pgxv4.RegisterDriver("cloudsql-postgres", cloudsqlconn.WithIAMAuthN())
	if err != nil {
		log.Printf("Err: %v", err)
		return
	}
	defer cleanup()

	log.Printf("Establishing connection to cloud SQL / DB")
	db, err = sql.Open(
		"cloudsql-postgres",
		// TODO: change hardcoded params to dynamic ones
		"host=test-cloud-symbol-server:europe-west1:db user=admin-api@test-cloud-symbol-server.iam dbname=cloud_symbol_server sslmode=disable",
	)
	if err != nil {
		log.Printf("Err: %v", err)
		return
	}
	log.Printf("connection up!")
}

func GetDB() *sql.DB {
	return db
}

func sqlCreateStore(ctx context.Context, storeId string) error {
	var store = models.Store{
		Name: storeId,
	}
	err := store.Insert(ctx, db, boil.Infer())
	return err
}

func sqlGetStore(ctx context.Context, storeId string) (*models.Store, error) {
	store, err := models.Stores(qm.Where("name = ?", storeId)).One(ctx, db)
	return store, err
}

func sqlDeleteStore(ctx context.Context, storeId string) error {
	store, err := sqlGetStore(ctx, storeId)
	if err != nil {
		return err
	}
	_, err = store.Delete(ctx, db)
	return err
}
