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

func TestExpireRemovesOnlyFilesWithZeroReferences(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"

	if err := ensureTestStoreExists(apiClient, authContext, storeId); err != nil {
		t.Fatalf("ensureTestStoreExists failed: %v", err)
	}

	// Create uploads 1 & 2

	uploadId1, err := upload(apiClient, authContext, storeId, &exampleUpload1)
	if err != nil {
		t.Fatalf("upload 1 failed: %v", err)
	}

	uploadId2, err := upload(apiClient, authContext, storeId, &exampleUpload2)
	if err != nil {
		t.Fatalf("upload 2 failed: %v", err)
	}

	// Validate files from uploads 1 & 2 exist

	{
		getStoreFileIdsResponse, _, err := apiClient.DefaultApi.GetStoreFileIds(authContext, storeId).Execute()
		if err != nil {
			t.Fatalf("Get file ids failed: %v", err)
		}

		expectedNumFiles := 3
		if expectedNumFiles != len(getStoreFileIdsResponse.Items) {
			t.Fatalf("After upload 1 + 2, there should be %d files present, but only %d found", expectedNumFiles, len(getStoreFileIdsResponse.Items))
		}
	}

	// Expire upload 1

	_, err = apiClient.DefaultApi.ExpireStoreUpload(authContext, uploadId1, storeId).Execute()
	if err != nil {
		t.Fatalf("Expire failed: %v", err)
	}

	// Validate only files from upload 2 exist

	{
		getStoreFileIdsResponse, _, err := apiClient.DefaultApi.GetStoreFileIds(authContext, storeId).Execute()
		if err != nil {
			t.Fatalf("Get file ids failed: %v", err)
		}

		expectedNumFiles := 2
		if expectedNumFiles != len(getStoreFileIdsResponse.Items) {
			t.Fatalf("After upload 1 + 2 and expire 1, there should be %d files present, but only %d found", expectedNumFiles, len(getStoreFileIdsResponse.Items))
		}
	}

	// Expire upload 2

	_, err = apiClient.DefaultApi.ExpireStoreUpload(authContext, uploadId2, storeId).Execute()
	if err != nil {
		t.Fatalf("Expire failed: %v", err)
	}

	// Validate no files exist

	{
		getStoreFileIdsResponse, _, err := apiClient.DefaultApi.GetStoreFileIds(authContext, storeId).Execute()
		if err != nil {
			t.Fatalf("Get file ids failed: %v", err)
		}

		expectedNumFiles := 0
		if expectedNumFiles != len(getStoreFileIdsResponse.Items) {
			t.Fatalf("After upload 1 + 2 and expire 1 + 2, there should be %d files present, but only %d found", expectedNumFiles, len(getStoreFileIdsResponse.Items))
		}
	}
}
