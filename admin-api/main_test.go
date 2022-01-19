package admin_api

import (
	"context"
	"net/http"
	"os"
	"testing"

	openapi_client "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-client"
)

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

func TestGetTransactionWithInvalidCredentialsFails(t *testing.T) {

	email := "invalidemail"
	pat := "invalidpat"

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"
	transactionId := "invalidtransactionId"

	_, r, err := apiClient.DefaultApi.GetTransaction(authContext, transactionId, storeId).Execute()
	desiredStatusCode := http.StatusUnauthorized
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetTransaction with invalid email/PAT is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetTransactionThatDoesNotExistFails(t *testing.T) {

	email := "testuser"
	pat := "testpat"

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"
	transactionId := "invalidtransactionId"

	_, r, err := apiClient.DefaultApi.GetTransaction(authContext, transactionId, storeId).Execute()
	desiredStatusCode := http.StatusNotFound
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetTransaction with invalid transaction ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestUploadTransactionSucceeds(t *testing.T) {

	email := "testuser"
	pat := "testpat"

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"

	description := "test upload"
	buildId := "test build id"

	fileName1 := "file1"
	hash1 := "hash1"
	fileName2 := "file2"
	hash2 := "hash2"

	files := []openapi_client.UploadFileRequest{
		{
			FileName: &fileName1,
			Hash:     &hash1,
		},
		{
			FileName: &fileName2,
			Hash:     &hash2,
		},
	}

	uploadTransactionRequest := openapi_client.UploadTransactionRequest{
		Description: &description,
		BuildId:     &buildId,
		Files:       &files,
	}

	_, r, err := apiClient.DefaultApi.CreateTransaction(authContext, storeId).UploadTransactionRequest(uploadTransactionRequest).Execute()
	desiredStatusCode := http.StatusOK
	if err != nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("CreateTransaction is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}

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

func createStore(apiClient *openapi_client.APIClient, authContext context.Context, storeId string, acceptStoreAlreadyExists bool) error {
	r, err := apiClient.DefaultApi.CreateStore(authContext, storeId).Execute()
	if err != nil {
		if !acceptStoreAlreadyExists && r.StatusCode != http.StatusConflict {
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
		t.Fatalf("DeleteeStore failed: %v", err)
	}

	err = createStore(apiClient, authContext, storeId, false)
	if err != nil {
		t.Fatalf("CreateStore failed: %v", err)
	}

	err = deleteStore(apiClient, authContext, storeId, false)
	if err != nil {
		t.Fatalf("DeleteStore failed: %v", err)
	}
}
