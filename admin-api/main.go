package admin_api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	authentication "github.com/falldamagestudio/cloud-symbol-server/admin-api/authentication"
	postgres "github.com/falldamagestudio/cloud-symbol-server/admin-api/postgres"
)

type ApiService struct {
}

var corsHandler *cors.Cors
var router *mux.Router

const (
	env_GCP_PROJECT_ID = "GCP_PROJECT_ID"
)

func init() {

	postgres.InitSQL()

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

	clientID := os.Getenv(env_GCP_PROJECT_ID)
	if clientID == "" {
		panic(fmt.Sprintf("%v must be set", env_GCP_PROJECT_ID))
	}

	usernamePasswordValidator := authentication.CreateUsernamePasswordValidator()
	authTokenValidator, err := authentication.CreateAuthTokenValidator(clientID)
	if err != nil {
		panic(err)
	}
	authenticationMiddleware := authentication.CreateAuthenticationMiddleware([]authentication.Validator{
		usernamePasswordValidator,
		authTokenValidator,
	})

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
