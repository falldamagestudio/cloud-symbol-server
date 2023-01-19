package backend_api

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	openapi_client "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-client"
)

func TestGetStoreFileBlobsWithInvalidCredentialsFails(t *testing.T) {

	email := "invalidemail"
	pat := "invalidpat"

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"
	fileId := "file1"

	_, r, err := apiClient.DefaultApi.GetStoreFileBlobs(authContext, storeId, fileId).Execute()
	desiredStatusCode := http.StatusUnauthorized
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreFileBlobs with invalid email/PAT is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreFileBlobsForStoreThatDoesNotExistFails(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	invalidStoreId := "invalidstoreid"
	fileId := "file1"

	_, r, err := apiClient.DefaultApi.GetStoreFileBlobs(authContext, invalidStoreId, fileId).Execute()
	desiredStatusCode := http.StatusNotFound
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreFileBlobs with invalid store ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreFileBlobsForStoreFileThatDoesNotExistFails(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"

	if err := ensureTestStoreExists(apiClient, authContext, storeId); err != nil {
		t.Fatalf("ensureTestStoreExists failed: %v", err)
	}

	invalidFileId := "file1"

	_, r, err := apiClient.DefaultApi.GetStoreFileBlobs(authContext, storeId, invalidFileId).Execute()
	desiredStatusCode := http.StatusNotFound
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreFileBlobs with invalid file ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreFileBlobsForStoreFileExistsSucceeds(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"

	if err := ensureTestStoreExistsAndIsPopulated(apiClient, authContext, storeId); err != nil {
		t.Fatalf("ensureTestStoreExistsAndIsPopulated failed: %v", err)
	}

	fileId := "file1"

	_, r, err := apiClient.DefaultApi.GetStoreFileBlobs(authContext, storeId, fileId).Execute()
	desiredStatusCode := http.StatusOK
	if err != nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreFileBlobs with valid file ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestDownloadStoreFileBlobForStoreFileBlobExistsSucceeds(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"

	if err := ensureTestStoreExistsAndIsPopulated(apiClient, authContext, storeId); err != nil {
		t.Fatalf("ensureTestStoreExistsAndIsPopulated failed: %v", err)
	}

	fileId := "file1"
	blobId := "blobIdentifier1"
	content1 := "content1"

	getStoreFileBlobDownloadUrlResponse, r, err := apiClient.DefaultApi.GetStoreFileBlobDownloadUrl(authContext, storeId, fileId, blobId).Execute()
	desiredStatusCode := http.StatusOK
	if err != nil && desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreFileBlobDownloadUrl with valid file-blob ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}

	const desiredMethod = "GET"
	if getStoreFileBlobDownloadUrlResponse.Method != desiredMethod {
		t.Fatalf("GetStoreFileBlobDownloadUrl is expected to return a URL that should be used with %v method, but gave %v method", desiredMethod, getStoreFileBlobDownloadUrlResponse.Method)
	}

	downloadFileResponse, err := http.Get(getStoreFileBlobDownloadUrlResponse.Url)
	if err != nil {
		t.Fatalf("Error when downloading content: %v", err)
	}

	binaryContent, err := ioutil.ReadAll(downloadFileResponse.Body)
	if err != nil {
		t.Fatalf("Error while reading response body: %v", err)
	}
	content := string(binaryContent)
	if content1 != content {
		t.Fatalf("Downloaded file should contain the content '%v', but contains '%v'", content1, content)
	}
}

func TestGetStoreFileBlobsWithPaginationSucceeds(t *testing.T) {

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
	blobIdentifier1 := "blobIdentifier1"
	type1 := openapi_client.StoreFileBlobType("pe")
	size1 := int64(8)
	contentHash1 := "d0b425e00e15a0d36b9b361f02bab63563aed6cb4665083905386c55d5b679fa" // SHA256 hash of "content1"
	fileName2 := "file1"
	blobIdentifier2 := "blobIdentifier2"
	type2 := openapi_client.StoreFileBlobType("pdb")
	size2 := int64(9)
	contentHash2 := "35c6a7f16428d39c386bd4ebb4c9e0d256bae81634acdf8e65e21bc0abebd0d5" // SHA256 hash of "content2_"
	fileName3 := "file2"
	blobIdentifier3 := "blobIdentifier3"
	type3 := openapi_client.StoreFileBlobType("pdb")
	size3 := int64(9)
	contentHash3 := "a722ad33c435a5d321f89b00e2bc019fcb758384fddc896877cfb19eca65b6d0" // SHA256 hash of "content3_"

	files := []openapi_client.CreateStoreUploadFileRequest{
		{
			FileName:       fileName1,
			BlobIdentifier: blobIdentifier1,
			Type:           &type1,
			Size:           &size1,
			ContentHash:    &contentHash1,
		},
		{
			FileName:       fileName2,
			BlobIdentifier: blobIdentifier2,
			Type:           &type2,
			Size:           &size2,
			ContentHash:    &contentHash2,
		},
		{
			FileName:       fileName3,
			BlobIdentifier: blobIdentifier3,
			Type:           &type3,
			Size:           &size3,
			ContentHash:    &contentHash3,
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

	// Ensure this upload exists

	getStoreUploadsResponse, r, err := apiClient.DefaultApi.GetStoreUploads(authContext, storeId).Offset(0).Limit(100).Execute()
	desiredStatusCode = http.StatusOK
	if err != nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreUploads is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}

	// Ensure there is a single upload

	expectedNumStoreUploads := 1
	if len(getStoreUploadsResponse.Uploads) != expectedNumStoreUploads {
		t.Fatalf("Expected GetStoreUploads returns %v results, but returned %v results: %v", expectedNumStoreUploads, len(getStoreUploadsResponse.Uploads), getStoreUploadsResponse)
	}

	// Ensure upload is in "in_progress" status

	expectedUploadStatus := openapi_client.STOREUPLOADSTATUS_IN_PROGRESS
	if getStoreUploadsResponse.Uploads[0].Status != expectedUploadStatus {
		t.Fatalf("GetStoreUpload should return an upload with status %v, but it has status %v", expectedUploadStatus, getStoreUploadsResponse.Uploads[0].Status)
	}

	// Ensure file is in "pending" status

	expectedUploadFileStatus := openapi_client.STOREUPLOADFILESTATUS_PENDING
	if getStoreUploadsResponse.Uploads[0].Files[0].Status != expectedUploadFileStatus {
		t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, getStoreUploadsResponse.Uploads[0].Files[0].Status)
	}

	// Complete upload of first file

	{
		storeUploadId := "0"
		fileId := int32(0)

		r, err = apiClient.DefaultApi.MarkStoreUploadFileUploaded(authContext, storeUploadId, storeId, fileId).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("MarkStoreUploadFileUploaded is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		// Ensure upload is in "in_progress" status

		getStoreUploadResponse, r, err := apiClient.DefaultApi.GetStoreUpload(authContext, storeUploadId, storeId).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("GetStoreUpload is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		expectedUploadStatus := openapi_client.STOREUPLOADSTATUS_IN_PROGRESS
		if getStoreUploadResponse.Status != expectedUploadStatus {
			t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadStatus, getStoreUploadResponse.Status)
		}

		// Ensure file has transitioned to "uploaded" status

		expectedUploadFileStatus := openapi_client.STOREUPLOADFILESTATUS_COMPLETED
		if (getStoreUploadResponse.Files)[0].Status != expectedUploadFileStatus {
			t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, (getStoreUploadResponse.Files)[0].Status)
		}

		// Fetch store-file info

		getStoreFileBlobsResponse, r, err := apiClient.DefaultApi.GetStoreFileBlobs(authContext, storeId, fileName1).Offset(0).Limit(100).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("GetStoreFileBlobs is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		// Ensure there are exactly two file-blobs for file 1

		expectedNumResults := 2
		expectedTotalResults := int32(2)
		if (len(getStoreFileBlobsResponse.Blobs) != expectedNumResults) || (getStoreFileBlobsResponse.Pagination.Total != expectedTotalResults) || (getStoreFileBlobsResponse.Blobs[0].BlobIdentifier != blobIdentifier1) || ((getStoreFileBlobsResponse.Blobs)[1].BlobIdentifier != blobIdentifier2) {
			t.Fatalf("GetStoreFileBlobs should show %v results with blobs %v & %v, and %v total, but shows the following blobs: %v, and total: %v", expectedNumResults, blobIdentifier1, blobIdentifier2, expectedTotalResults, getStoreFileBlobsResponse.Blobs, getStoreFileBlobsResponse.Pagination.Total)
		}

		// Ensure blob metadata is correct

		require.Equal(t, blobIdentifier1, getStoreFileBlobsResponse.Blobs[0].BlobIdentifier)
		require.Equal(t, type1, *getStoreFileBlobsResponse.Blobs[0].Type)
		require.Equal(t, size1, *getStoreFileBlobsResponse.Blobs[0].Size)
		require.Equal(t, contentHash1, *getStoreFileBlobsResponse.Blobs[0].ContentHash)

		require.Equal(t, blobIdentifier2, getStoreFileBlobsResponse.Blobs[1].BlobIdentifier)
		require.Equal(t, type2, *getStoreFileBlobsResponse.Blobs[1].Type)
		require.Equal(t, size2, *getStoreFileBlobsResponse.Blobs[1].Size)
		require.Equal(t, contentHash2, *getStoreFileBlobsResponse.Blobs[1].ContentHash)

	}

	// Try some pagination
	{
		// Fetch store-file info with offset = 0, limit = 1

		getStoreFileBlobsResponse, r, err := apiClient.DefaultApi.GetStoreFileBlobs(authContext, storeId, fileName1).Offset(0).Limit(1).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("GetStoreFileBlobs is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		// Ensure this returned just the first entry

		expectedNumResults := 1
		expectedTotalResults := int32(2)
		if (len(getStoreFileBlobsResponse.Blobs) != expectedNumResults) || (getStoreFileBlobsResponse.Pagination.Total != expectedTotalResults) || (getStoreFileBlobsResponse.Blobs[0].BlobIdentifier != blobIdentifier1) {
			t.Fatalf("GetStoreFileBlobs should show %v result with blob %v, and %v total, but shows the following files: %v, and total: %v", expectedNumResults, blobIdentifier1, expectedTotalResults, getStoreFileBlobsResponse.Blobs, getStoreFileBlobsResponse.Pagination.Total)
		}
	}

	{
		// Fetch store-file info with offset = 1, limit = 1

		getStoreFileBlobsResponse, r, err := apiClient.DefaultApi.GetStoreFileBlobs(authContext, storeId, fileName1).Offset(1).Limit(1).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("GetStoreFileBlobs is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		// Ensure this returned just the second entry

		expectedNumResults := 1
		expectedTotalResults := int32(2)
		if (len(getStoreFileBlobsResponse.Blobs) != expectedNumResults) || (getStoreFileBlobsResponse.Pagination.Total != expectedTotalResults) || (getStoreFileBlobsResponse.Blobs[0].BlobIdentifier != blobIdentifier2) {
			t.Fatalf("GetStoreFileBlobs should show %v results with blob %v, and %v total, but shows the following files: %v, and total: %v", expectedNumResults, blobIdentifier2, expectedTotalResults, getStoreFileBlobsResponse.Blobs, getStoreFileBlobsResponse.Pagination.Total)
		}
	}
}
