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

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		// Allow API calls from any web origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// The calling party is allowed to cache the results from a preflight request for this many seconds
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	router.ServeHTTP(w, r)
}
