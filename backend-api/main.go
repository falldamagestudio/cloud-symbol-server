package backend_api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	authentication "github.com/falldamagestudio/cloud-symbol-server/backend-api/authentication"
	openapi "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-server/go"
	http_symbol_store "github.com/falldamagestudio/cloud-symbol-server/backend-api/http-symbol-store"
	postgres "github.com/falldamagestudio/cloud-symbol-server/backend-api/postgres"
)

type ApiService struct {
}

var corsHandler *cors.Cors
var router *mux.Router

const (
	env_GCP_PROJECT_ID    = "GCP_PROJECT_ID"
	httpSymbolStorePrefix = "/httpSymbolStore/"
)

func init() {

	postgres.InitSQL()

	apiService := &ApiService{}
	DefaultApiController := openapi.NewDefaultApiController(apiService)

	router = openapi.NewRouter(DefaultApiController)

	// Set up CORS handler

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

	// Set up authentication

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

	// Set up HTTP Symbol Store handler

	httpSymbolStoreHandler := http_symbol_store.CreateHttpSymbolStoreHandler()
	router.PathPrefix(httpSymbolStorePrefix).Handler(http.StripPrefix(httpSymbolStorePrefix, httpSymbolStoreHandler))
}

func BackendAPI(w http.ResponseWriter, r *http.Request) {
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
