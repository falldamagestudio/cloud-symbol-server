package admin_api

import (
	"context"
	"os"

	openapi_client "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-client"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getAdminAPIEndpoint() string {

	adminAPIEndpoint := os.Getenv("ADMIN_API_ENDPOINT")
	if adminAPIEndpoint == "" {
		adminAPIEndpoint = "http://localhost:8080"
	}
	return adminAPIEndpoint
}

func getAPIClient(email string, pat string) (context.Context, *openapi_client.APIClient) {

	authContext := context.WithValue(context.Background(), openapi_client.ContextBasicAuth, openapi_client.BasicAuth{
		UserName: email,
		Password: pat,
	})

	configuration := openapi_client.NewConfiguration()
	configuration.Servers[0].URL = getAdminAPIEndpoint()
	api_client := openapi_client.NewAPIClient(configuration)

	return authContext, api_client
}
