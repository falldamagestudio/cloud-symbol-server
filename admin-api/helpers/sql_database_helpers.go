package helpers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
)

var db *sql.DB

type ErrCloudSQLInstance struct {
}

func (err ErrCloudSQLInstance) Error() string {
	return "No cloud SQL instance configured"
}

func getCloudSQLInstance() (string, error) {

	cloudSQLInstance := os.Getenv("CLOUD_SQL_INSTANCE")
	if cloudSQLInstance == "" {
		return "", &ErrCloudSQLInstance{}
	}

	return cloudSQLInstance, nil
}

type ErrCloudSQLUser struct {
}

func (err ErrCloudSQLUser) Error() string {
	return "No cloud SQL user configured"
}

func getCloudSQLUser() (string, error) {

	cloudSQLUser := os.Getenv("CLOUD_SQL_USER")
	if cloudSQLUser == "" {
		return "", &ErrCloudSQLUser{}
	}

	return cloudSQLUser, nil
}

func InitSQL() {

	cloudSQLInstance, err := getCloudSQLInstance()
	if err != nil {
		log.Printf("Err: %v", err)
		return
	}

	cloudSQLUser, err := getCloudSQLUser()
	if err != nil {
		log.Printf("Err: %v", err)
		return
	}

	dbDriver := "cloudsql-postgres"

	log.Printf("Registering cloudsql-postgres driver")
	cleanup, err := pgxv4.RegisterDriver(dbDriver, cloudsqlconn.WithIAMAuthN())
	if err != nil {
		log.Printf("Err: %v", err)
		return
	}
	defer cleanup()

	dbName := "cloud_symbol_server"

	log.Printf("Establishing connection to cloud SQL / DB")
	db, err = sql.Open(
		dbDriver,
		fmt.Sprintf("host=%v user=%v dbname=%v sslmode=disable", cloudSQLInstance, cloudSQLUser, dbName),
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