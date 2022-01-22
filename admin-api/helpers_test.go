package admin_api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
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

func getStores(apiClient *openapi_client.APIClient, authContext context.Context) ([]string, error) {
	result, _, err := apiClient.DefaultApi.GetStores(authContext).Execute()
	if err != nil {
		return nil, err
	} else {
		return result.Items, err
	}
}

func createStore(apiClient *openapi_client.APIClient, authContext context.Context, storeId string, acceptStoreAlreadyExists bool) error {
	r, err := apiClient.DefaultApi.CreateStore(authContext, storeId).Execute()
	if err != nil {
		if !acceptStoreAlreadyExists && r.StatusCode != http.StatusConflict {
			return err
		}
	}
	return nil
}

func deleteStore(apiClient *openapi_client.APIClient, authContext context.Context, storeId string, acceptStoreDoesNotExist bool) error {
	r, err := apiClient.DefaultApi.DeleteStore(authContext, storeId).Execute()
	if err != nil {
		if !acceptStoreDoesNotExist && r.StatusCode != http.StatusNotFound {
			return err
		}
	}
	return nil
}

func ensureTestStoreExists(apiClient *openapi_client.APIClient, authContext context.Context, storeId string) error {

	err := deleteStore(apiClient, authContext, storeId, true)
	if err != nil {
		return err
	}

	stores1, err := getStores(apiClient, authContext)
	if err != nil {
		return err
	}
	if stringInSlice(storeId, stores1) {
		return errors.New(fmt.Sprintf("Store %v should not be among stores, but is: %v", storeId, stores1))
	}

	err = createStore(apiClient, authContext, storeId, false)
	if err != nil {
		return err
	}

	stores2, err := getStores(apiClient, authContext)
	if err != nil {
		return err
	}
	if !stringInSlice(storeId, stores2) {
		return errors.New(fmt.Sprintf("Store %v should be among stores, but is not: %v", storeId, stores2))
	}

	return nil
}
