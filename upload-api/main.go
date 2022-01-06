package upload_api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	openapi "github.com/falldamagestudio/cloud-symbol-store/upload-api/generated/go"
	"github.com/gorilla/mux"
)

type ApiService struct {
}

var router *mux.Router
var apiService openapi.DefaultApiServicer

func init() {
	apiService = &ApiService{}
	DefaultApiController := openapi.NewDefaultApiController(apiService)

	router = openapi.NewRouter(DefaultApiController)

	patAM := &patAuthenticationMiddleware{}
	router.Use(patAM.Middleware)
}

var localFirestoreClient *firestore.Client

func firestoreClient(context context.Context) (*firestore.Client, error) {

	if localFirestoreClient == nil {

		gcpProjectId := os.Getenv("GCP_PROJECT_ID")
		if gcpProjectId == "" {
			log.Print("No GCP Project ID configured")
			return nil, errors.New("no GCP Project ID configured")
		}

		err := (error)(nil)
		localFirestoreClient, err = firestore.NewClient(context, gcpProjectId)
		if err != nil {
			log.Printf("Unable to create firestoreClient: %v", err)
			return nil, errors.New("unable to create firestoreClient")
		}
	}

	return localFirestoreClient, nil
}

func UploadAPI(w http.ResponseWriter, r *http.Request) {

	router.ServeHTTP(w, r)
}
