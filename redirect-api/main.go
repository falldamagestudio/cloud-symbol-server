package redirect_api

import (
	"fmt"
	"log"
	"net/http"
    "net/http/httputil"
	"net/url"
	"os"
)

const (
	env_TARGET_URI = "TARGET_URI"
)

var targetURI = ""

func init() {

	targetURI = os.Getenv(env_TARGET_URI)
	if targetURI == "" {
		panic(fmt.Sprintf("%v must be set", env_TARGET_URI))
	}
}

// NewProxy takes target host and creates a reverse proxy
func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
    url, err := url.Parse(targetHost)
    if err != nil {
        return nil, err
    }
 
    return httputil.NewSingleHostReverseProxy(url), nil
}
 
func RedirectAPI(w http.ResponseWriter, r *http.Request) {
	log.Printf("RedirectAPI invoked: Path: %v, Method: %v ", r.URL.Path, r.Method)
	username, pass, basicAuthPresent := r.BasicAuth()
	log.Printf("BasicAuth: %v, %v, %v", username, pass, basicAuthPresent)

    // initialize a reverse proxy and pass the actual backend server url here
    proxy, err := NewProxy(targetURI)
	if err != nil {
		log.Printf("Error creating reverse proxy: %v", err)
		http.Error(w, fmt.Sprintf("Error creating reverse proxy: %v", err), http.StatusInternalServerError)
		return
	}

	log.Print("Passed URL to reverse proxy")

	proxy.ServeHTTP(w, r)
}
