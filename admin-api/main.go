package admin_api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	helpers "github.com/falldamagestudio/cloud-symbol-server/admin-api/helpers"
)

type ApiService struct {
}

var router *mux.Router
var apiService openapi.DefaultApiServicer

func init() {

	helpers.InitSQL()

	apiService = &ApiService{}
	DefaultApiController := openapi.NewDefaultApiController(apiService)

	router = openapi.NewRouter(DefaultApiController)

	patAM := &patAuthenticationMiddleware{}
	router.Use(patAM.Middleware)
}

func AdminAPI(w http.ResponseWriter, r *http.Request) {
	log.Print("Path called: " + r.URL.Path)

	router.ServeHTTP(w, r)
}
