package admin_api

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	helpers "github.com/falldamagestudio/cloud-symbol-server/admin-api/helpers"
)

type ApiService struct {
}

var handler http.Handler
var apiService openapi.DefaultApiServicer

func init() {

	helpers.InitSQL()

	apiService = &ApiService{}
	DefaultApiController := openapi.NewDefaultApiController(apiService)

	router := openapi.NewRouter(DefaultApiController)

	patAM := &patAuthenticationMiddleware{}
	router.Use(patAM.Middleware)

	handler = cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "HEAD", "GET", "PUT", "POST", "PATCH", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(router)
}

func AdminAPI(w http.ResponseWriter, r *http.Request) {
	log.Print("Path called: " + r.URL.Path)

	handler.ServeHTTP(w, r)
}
