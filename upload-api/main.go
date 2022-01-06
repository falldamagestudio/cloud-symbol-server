package upload_api

import (
	"net/http"

	openapi "github.com/falldamagestudio/cloud-symbol-store/upload-api/api/go"
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

func UploadAPI(w http.ResponseWriter, r *http.Request) {

	router.ServeHTTP(w, r)
}
