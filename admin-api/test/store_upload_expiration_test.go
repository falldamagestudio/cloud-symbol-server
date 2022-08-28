package admin_api

import (
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
