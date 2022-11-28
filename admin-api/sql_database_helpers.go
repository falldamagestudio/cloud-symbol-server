package admin_api

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
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

func BeginDBTransaction(ctx context.Context) (*sql.Tx, error) {
	if db == nil {
		return nil, errors.New("no DB")
	}

	return db.BeginTx(ctx, nil)
}
