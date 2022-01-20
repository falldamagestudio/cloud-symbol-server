package admin_api

import (
	"net/http"
	"testing"

	openapi_client "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-client"
)

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

	if err := ensureTestStoreExists(apiClient, authContext, storeId); err != nil {
		t.Fatalf("ensureTestStoreExists failed: %v", err)
	}

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

	if err := ensureTestStoreExists(apiClient, authContext, storeId); err != nil {
		t.Fatalf("ensureTestStoreExists failed: %v", err)
	}

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
