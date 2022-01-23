package admin_api_test

import (
	"net/http"
	"reflect"
	"testing"

	openapi_client "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-client"
)

func TestGetStoreUploadWithInvalidCredentialsFails(t *testing.T) {

	email := "invalidemail"
	pat := "invalidpat"

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"
	uploadId := "invaliduploadid"

	_, r, err := apiClient.DefaultApi.GetStoreUpload(authContext, uploadId, storeId).Execute()
	desiredStatusCode := http.StatusUnauthorized
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreUpload with invalid email/PAT is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreUploadThatDoesNotExistFails(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"
	uploadId := "invaliduploadid"

	if err := ensureTestStoreExists(apiClient, authContext, storeId); err != nil {
		t.Fatalf("ensureTestStoreExists failed: %v", err)
	}

	_, r, err := apiClient.DefaultApi.GetStoreUpload(authContext, uploadId, storeId).Execute()
	desiredStatusCode := http.StatusNotFound
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreUpload with invalid upload ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestCreateStoreUploadSucceeds(t *testing.T) {

	email, pat := getTestEmailAndPat()

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

	createStoreUploadRequest := openapi_client.CreateStoreUploadRequest{
		Description: &description,
		BuildId:     &buildId,
		Files:       &files,
	}

	_, r, err := apiClient.DefaultApi.CreateStoreUpload(authContext, storeId).CreateStoreUploadRequest(createStoreUploadRequest).Execute()
	desiredStatusCode := http.StatusOK
	if err != nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("CreateStoreUpload is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}

	getStoreUploadIdsResponse, r, err := apiClient.DefaultApi.GetStoreUploadIds(authContext, storeId).Execute()
	desiredStatusCode = http.StatusOK
	if err != nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreUploads is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}

	expectedStoreUploads := []string{"0"}
	if !reflect.DeepEqual(expectedStoreUploads, getStoreUploadIdsResponse.Items) {
		t.Fatalf("Expected GetStoreUploads is expected to return %v, but returned %v", expectedStoreUploads, getStoreUploadIdsResponse.Items)
	}
}
