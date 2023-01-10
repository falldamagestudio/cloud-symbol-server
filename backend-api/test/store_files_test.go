package backend_api

import (
	"net/http"
	"reflect"
	"testing"

	openapi_client "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-client"
)

func TestGetStoreFilesWithInvalidCredentialsFails(t *testing.T) {

	email := "invalidemail"
	pat := "invalidpat"

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"

	_, r, err := apiClient.DefaultApi.GetStoreFiles(authContext, storeId).Execute()
	desiredStatusCode := http.StatusUnauthorized
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreFiles with invalid email/PAT is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreFilesForStoreThatDoesNotExistFails(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	invalidStoreId := "invalidstoreid"

	_, r, err := apiClient.DefaultApi.GetStoreFiles(authContext, invalidStoreId).Execute()
	desiredStatusCode := http.StatusNotFound
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreFiles with invalid store ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreFilesForStoreExistsSucceeds(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"

	if err := ensureTestStoreExists(apiClient, authContext, storeId); err != nil {
		t.Fatalf("ensureTestStoreExists failed: %v", err)
	}

	_, r, err := apiClient.DefaultApi.GetStoreFiles(authContext, storeId).Execute()
	desiredStatusCode := http.StatusOK
	if err != nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreFiles with valid store ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreFilesWithPaginationSucceeds(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"

	if err := ensureTestStoreExists(apiClient, authContext, storeId); err != nil {
		t.Fatalf("ensureTestStoreExists failed: %v", err)
	}

	// Create initial upload

	description := "test upload"
	buildId := "test build id"

	fileName1 := "file1"
	hash1 := "hash1"
	fileName2 := "file2"
	hash2 := "hash2"

	files := []openapi_client.UploadFileRequest{
		{
			FileName: fileName1,
			Hash:     hash1,
		},
		{
			FileName: fileName2,
			Hash:     hash2,
		},
	}

	useProgressApi := true
	createStoreUploadRequest := openapi_client.CreateStoreUploadRequest{
		Description:    &description,
		BuildId:        &buildId,
		Files:          files,
		UseProgressApi: &useProgressApi,
	}

	_, r, err := apiClient.DefaultApi.CreateStoreUpload(authContext, storeId).CreateStoreUploadRequest(createStoreUploadRequest).Execute()
	desiredStatusCode := http.StatusOK
	if err != nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("CreateStoreUpload is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}

	// Ensure this upload exists, and has been assigned ID "0"

	getStoreUploadIdsResponse, r, err := apiClient.DefaultApi.GetStoreUploadIds(authContext, storeId).Execute()
	desiredStatusCode = http.StatusOK
	if err != nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreUploadIds is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}

	expectedStoreUploadIds := []string{"0"}
	if !reflect.DeepEqual(expectedStoreUploadIds, getStoreUploadIdsResponse) {
		t.Fatalf("Expected GetStoreUploadIds is expected to return %v, but returned %v", expectedStoreUploadIds, getStoreUploadIdsResponse)
	}

	// Ensure upload is in "in_progress" status

	storeUploadId := "0"

	getStoreUploadResponse, r, err := apiClient.DefaultApi.GetStoreUpload(authContext, storeUploadId, storeId).Execute()
	desiredStatusCode = http.StatusOK
	if err != nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreUpload is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}

	expectedUploadStatus := "in_progress"
	if getStoreUploadResponse.Status != expectedUploadStatus {
		t.Fatalf("GetStoreUpload should return an upload with status %v, but it has status %v", expectedUploadStatus, getStoreUploadResponse.Status)
	}

	// Ensure file is in "pending" status

	expectedUploadFileStatus := "pending"
	if getStoreUploadResponse.Files[0].Status != expectedUploadFileStatus {
		t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, getStoreUploadResponse.Files[0].Status)
	}

	// Complete upload of first file

	{
		fileId := int32(0)

		r, err = apiClient.DefaultApi.MarkStoreUploadFileUploaded(authContext, storeUploadId, storeId, fileId).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("MarkStoreUploadFileUploaded is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		// Ensure upload is in "in_progress" status

		storeUploadId := "0"

		getStoreUploadResponse, r, err := apiClient.DefaultApi.GetStoreUpload(authContext, storeUploadId, storeId).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("GetStoreUpload is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		expectedUploadStatus := "in_progress"
		if getStoreUploadResponse.Status != expectedUploadStatus {
			t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadStatus, getStoreUploadResponse.Status)
		}

		// Ensure file has transitioned to "uploaded" status

		expectedUploadFileStatus := "completed"
		if getStoreUploadResponse.Files[0].Status != expectedUploadFileStatus {
			t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, getStoreUploadResponse.Files[0].Status)
		}

		// Fetch store-file info

		getStoreFilesResponse, r, err := apiClient.DefaultApi.GetStoreFiles(authContext, storeId).Offset(0).Limit(100).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("GetStoreFiles is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		// Ensure there are exactly two store-files

		expectedNumResults := 2
		expectedTotalResults := int32(2)
		if (len(getStoreFilesResponse.Files) != expectedNumResults) || (getStoreFilesResponse.Pagination.Total != expectedTotalResults) || (getStoreFilesResponse.Files[0] != fileName1) || (getStoreFilesResponse.Files[1] != fileName2) {
			t.Fatalf("GetStoreFiles should show %v file with name %v & %v, and %v total, but shows the following files: %v, and total: %v", expectedNumResults, fileName1, fileName2, expectedTotalResults, getStoreFilesResponse.Files, getStoreFilesResponse.Pagination.Total)
		}
	}

	// Try some pagination
	{
		// Fetch store-file info with offset = 0, limit = 1

		getStoreFilesResponse, r, err := apiClient.DefaultApi.GetStoreFiles(authContext, storeId).Offset(0).Limit(1).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("GetStoreFiles is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		// Ensure this returned just the first entry

		expectedNumResults := 1
		expectedTotalResults := int32(2)
		if (len(getStoreFilesResponse.Files) != expectedNumResults) || (getStoreFilesResponse.Pagination.Total != expectedTotalResults) || (getStoreFilesResponse.Files[0] != fileName1) {
			t.Fatalf("GetStoreFiles should show %v file with name %v, and %v total, but shows the following files: %v, and total: %v", expectedNumResults, fileName1, expectedTotalResults, getStoreFilesResponse.Files, getStoreFilesResponse.Pagination.Total)
		}
	}

	{
		// Fetch store-file info with offset = 1, limit = 1

		getStoreFilesResponse, r, err := apiClient.DefaultApi.GetStoreFiles(authContext, storeId).Offset(1).Limit(1).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("GetStoreFiles is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		// Ensure this returned just the second entry

		expectedNumResults := 1
		expectedTotalResults := int32(2)
		if (len(getStoreFilesResponse.Files) != expectedNumResults) || (getStoreFilesResponse.Pagination.Total != expectedTotalResults) || (getStoreFilesResponse.Files[0] != fileName2) {
			t.Fatalf("GetStoreFiles should show %v file with name %v, and %v total, but shows the following files: %v, and total: %v", expectedNumResults, fileName2, expectedTotalResults, getStoreFilesResponse.Files, getStoreFilesResponse.Pagination.Total)
		}
	}
}
