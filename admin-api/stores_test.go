package admin_api

import (
	"context"
	"net/http"
	"testing"

	openapi_client "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-client"
)

func getStores(apiClient *openapi_client.APIClient, authContext context.Context) ([]string, error) {
	result, _, err := apiClient.DefaultApi.GetStores(authContext).Execute()
	if err != nil {
		return nil, err
	} else {
		return result, err
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

func TestCreateAndDestroyStore(t *testing.T) {

	email := "testuser"
	pat := "testpat"

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "test-store"

	err := deleteStore(apiClient, authContext, storeId, true)
	if err != nil {
		t.Fatalf("DeleteStore failed: %v", err)
	}

	stores1, err := getStores(apiClient, authContext)
	if err != nil {
		t.Fatalf("GetStores failed: %v", err)
	}
	if stringInSlice(storeId, stores1) {
		t.Fatalf("Store %v should not be among stores, but is: %v", storeId, stores1)
	}

	err = createStore(apiClient, authContext, storeId, false)
	if err != nil {
		t.Fatalf("CreateStore failed: %v", err)
	}

	stores2, err := getStores(apiClient, authContext)
	if err != nil {
		t.Fatalf("GetStores failed: %v", err)
	}
	if !stringInSlice(storeId, stores2) {
		t.Fatalf("Store %v should be among stores, but is not: %v", storeId, stores2)
	}

	err = deleteStore(apiClient, authContext, storeId, false)
	if err != nil {
		t.Fatalf("DeleteStore failed: %v", err)
	}

	stores3, err := getStores(apiClient, authContext)
	if err != nil {
		t.Fatalf("GetStores failed: %v", err)
	}
	if stringInSlice(storeId, stores3) {
		t.Fatalf("Store %v should not be among stores, but is: %v", storeId, stores3)
	}
}
