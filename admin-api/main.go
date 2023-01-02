package admin_api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	helpers "github.com/falldamagestudio/cloud-symbol-server/admin-api/helpers"
)

type ApiService struct {
}

var corsHandler *cors.Cors
var router *mux.Router

func init() {

	helpers.InitSQL()

	apiService := &ApiService{}
	DefaultApiController := openapi.NewDefaultApiController(apiService)

	router = openapi.NewRouter(DefaultApiController)

	corsHandler = cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:     []string{"*"},
		MaxAge:             3600,
		OptionsPassthrough: false,
		Debug:              false,
	})

	clientID := os.Getenv("GCP_PROJECT_ID")
	if clientID == "" {
		panic("GCP_PROJECT_ID must be set")
	}

	authenticationMiddleware, err := createAuthenticationMiddleware(clientID)
	if err != nil {
		panic("Failed creating authentication handler")
	}
	router.Use(authenticationMiddleware.Handler)
}

func AdminAPI(w http.ResponseWriter, r *http.Request) {
	log.Printf("API invoked: Path: %v, Method: %v ", r.URL.Path, r.Method)

	// Invoke CORS handler, and forward request to router in case it is a non-preflight request
	//
	// While the router has a mechanism for chaining to middlewares, the router will only do so for routes
	//  that it knows about. We don't list any of the OPTIONS-method routes in the OpenAPI spec,
	//  thus the router doesn't know about them, and thus the router will automatically return HTTP 405 Method Not Allowed
	//  for those routes, without giving the CORS middleware any chance to handle them.
	// We work around this by invoking the CORS middleware explicitly, before the router.
	corsHandler.ServeHTTP(w, r, router.ServeHTTP)
}
