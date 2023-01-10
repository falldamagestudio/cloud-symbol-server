package backend_api

import (
	"net/http"
	"reflect"
	"testing"

	openapi_client "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-client"
)

func TestGetStoreFileHashesWithInvalidCredentialsFails(t *testing.T) {

	email := "invalidemail"
	pat := "invalidpat"

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"
	fileId := "file1"

	_, r, err := apiClient.DefaultApi.GetStoreFileHashes(authContext, storeId, fileId).Execute()
	desiredStatusCode := http.StatusUnauthorized
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreFileHashes with invalid email/PAT is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreFileHashesForStoreThatDoesNotExistFails(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	invalidStoreId := "invalidstoreid"
	fileId := "file1"

	_, r, err := apiClient.DefaultApi.GetStoreFileHashes(authContext, invalidStoreId, fileId).Execute()
	desiredStatusCode := http.StatusNotFound
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreFileHashes with invalid store ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreFileHashesForStoreFileThatDoesNotExistFails(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"

	if err := ensureTestStoreExists(apiClient, authContext, storeId); err != nil {
		t.Fatalf("ensureTestStoreExists failed: %v", err)
	}

	invalidFileId := "file1"

	_, r, err := apiClient.DefaultApi.GetStoreFileHashes(authContext, storeId, invalidFileId).Execute()
	desiredStatusCode := http.StatusNotFound
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreFileHashes with invalid file ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreFileHashesForStoreFileExistsSucceeds(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"

	if err := ensureTestStoreExistsAndIsPopulated(apiClient, authContext, storeId); err != nil {
		t.Fatalf("ensureTestStoreExistsAndIsPopulated failed: %v", err)
	}

	fileId := "file1"

	_, r, err := apiClient.DefaultApi.GetStoreFileHashes(authContext, storeId, fileId).Execute()
	desiredStatusCode := http.StatusOK
	if err != nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreFileHashes with valid file ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreFileHashesWithPaginationSucceeds(t *testing.T) {

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
	fileName2 := "file1"
	hash2 := "hash2"
	fileName3 := "file2"
	hash3 := "hash3"

	files := []openapi_client.UploadFileRequest{
		{
			FileName: &fileName1,
			Hash:     &hash1,
		},
		{
			FileName: &fileName2,
			Hash:     &hash2,
		},
		{
			FileName: &fileName3,
			Hash:     &hash3,
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
	if *getStoreUploadResponse.Status != expectedUploadStatus {
		t.Fatalf("GetStoreUpload should return an upload with status %v, but it has status %v", expectedUploadStatus, *getStoreUploadResponse.Status)
	}

	// Ensure file is in "pending" status

	expectedUploadFileStatus := "pending"
	if *(getStoreUploadResponse.Files)[0].Status != expectedUploadFileStatus {
		t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, *(getStoreUploadResponse.Files)[0].Status)
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
		if *getStoreUploadResponse.Status != expectedUploadStatus {
			t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadStatus, *getStoreUploadResponse.Status)
		}

		// Ensure file has transitioned to "uploaded" status

		expectedUploadFileStatus := "completed"
		if *(getStoreUploadResponse.Files)[0].Status != expectedUploadFileStatus {
			t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, *(getStoreUploadResponse.Files)[0].Status)
		}

		// Fetch store-file info

		getStoreFileHashesResponse, r, err := apiClient.DefaultApi.GetStoreFileHashes(authContext, storeId, fileName1).Offset(0).Limit(100).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("GetStoreFileHashes is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		// Ensure there are exactly two file-hashes for file 1

		expectedNumResults := 2
		expectedTotalResults := int32(2)
		if (len(getStoreFileHashesResponse.Hashes) != expectedNumResults) || (*getStoreFileHashesResponse.Pagination.Total != expectedTotalResults) || (*(getStoreFileHashesResponse.Hashes)[0].Hash != hash1) || (*(getStoreFileHashesResponse.Hashes)[1].Hash != hash2) {
			t.Fatalf("GetStoreFileHashes should show %v results with hashes %v & %v, and %v total, but shows the following hashes: %v, and total: %v", expectedNumResults, hash1, hash2, expectedTotalResults, getStoreFileHashesResponse.Hashes, *getStoreFileHashesResponse.Pagination.Total)
		}
	}

	// Try some pagination
	{
		// Fetch store-file info with offset = 0, limit = 1

		getStoreFileHashesResponse, r, err := apiClient.DefaultApi.GetStoreFileHashes(authContext, storeId, fileName1).Offset(0).Limit(1).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("GetStoreFileHashes is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		// Ensure this returned just the first entry

		expectedNumResults := 1
		expectedTotalResults := int32(2)
		if (len(getStoreFileHashesResponse.Hashes) != expectedNumResults) || (*getStoreFileHashesResponse.Pagination.Total != expectedTotalResults) || (*(getStoreFileHashesResponse.Hashes)[0].Hash != hash1) {
			t.Fatalf("GetStoreFileHashes should show %v result with hash %v, and %v total, but shows the following files: %v, and total: %v", expectedNumResults, hash1, expectedTotalResults, getStoreFileHashesResponse.Hashes, *getStoreFileHashesResponse.Pagination.Total)
		}
	}

	{
		// Fetch store-file info with offset = 1, limit = 1

		getStoreFileHashesResponse, r, err := apiClient.DefaultApi.GetStoreFileHashes(authContext, storeId, fileName1).Offset(1).Limit(1).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("GetStoreFileHashes is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		// Ensure this returned just the second entry

		expectedNumResults := 1
		expectedTotalResults := int32(2)
		if (len(getStoreFileHashesResponse.Hashes) != expectedNumResults) || (*getStoreFileHashesResponse.Pagination.Total != expectedTotalResults) || (*(getStoreFileHashesResponse.Hashes)[0].Hash != hash2) {
			t.Fatalf("GetStoreFileHashes should show %v results with hash %v, and %v total, but shows the following files: %v, and total: %v", expectedNumResults, hash2, expectedTotalResults, getStoreFileHashesResponse.Hashes, *getStoreFileHashesResponse.Pagination.Total)
		}
	}
}
