package admin_api

import (
	"net/http"
	"testing"
)

func TestExpireSucceeds(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"

	if err := ensureTestStoreExists(apiClient, authContext, storeId); err != nil {
		t.Fatalf("ensureTestStoreExists failed: %v", err)
	}

	// Create initial upload

	uploadId, err := upload(apiClient, authContext, storeId, &exampleUpload1)
	if err != nil {
		t.Fatalf("upload failed failed: %v", err)
	}

	_, err = apiClient.DefaultApi.ExpireStoreUpload(authContext, uploadId, storeId).Execute()
	if err != nil {
		t.Fatalf("Expire failed: %v", err)
	}
}

func TestExpireFailsIfUploadDoesNotExist(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"
	invalidUploadId := "invalidUploadId"

	if err := ensureTestStoreExists(apiClient, authContext, storeId); err != nil {
		t.Fatalf("ensureTestStoreExists failed: %v", err)
	}

	r, err := apiClient.DefaultApi.ExpireStoreUpload(authContext, invalidUploadId, storeId).Execute()
	desiredStatusCode := http.StatusNotFound
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("Expire with invalid upload ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}
