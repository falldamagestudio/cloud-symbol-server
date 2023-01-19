package backend_api

import (
	"net/http"
	"testing"

	openapi_client "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-client"
)

func TestGetStoreUploadWithInvalidCredentialsFails(t *testing.T) {

	email := "invalidemail"
	pat := "invalidpat"

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"
	uploadId := int32(999999999)

	_, r, err := apiClient.DefaultApi.GetStoreUpload(authContext, uploadId, storeId).Execute()
	desiredStatusCode := http.StatusUnauthorized
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreUpload with invalid email/PAT is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreUploadsForStoreThatDoesNotExistFails(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	invalidStoreId := "invalidstoreid"

	_, r, err := apiClient.DefaultApi.GetStoreUploads(authContext, invalidStoreId).Execute()
	desiredStatusCode := http.StatusNotFound
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreUploads with invalid store ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreUploadIdsForStoreExistsSucceeds(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"

	if err := ensureTestStoreExists(apiClient, authContext, storeId); err != nil {
		t.Fatalf("ensureTestStoreExists failed: %v", err)
	}

	_, r, err := apiClient.DefaultApi.GetStoreUploads(authContext, storeId).Execute()
	desiredStatusCode := http.StatusOK
	if err != nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreUploads with valid store ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreUploadThatDoesNotExistFails(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"
	uploadId := int32(999999999)

	if err := ensureTestStoreExists(apiClient, authContext, storeId); err != nil {
		t.Fatalf("ensureTestStoreExists failed: %v", err)
	}

	_, r, err := apiClient.DefaultApi.GetStoreUpload(authContext, uploadId, storeId).Execute()
	desiredStatusCode := http.StatusNotFound
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreUpload with invalid upload ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestCreateStoreUploadWithoutProgressSucceeds(t *testing.T) {

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
	fileName2 := "file2"
	blobIdentifier2 := "blobIdentifier2"

	files := []openapi_client.CreateStoreUploadFileRequest{
		{
			FileName:       fileName1,
			BlobIdentifier: blobIdentifier1,
		},
		{
			FileName:       fileName2,
			BlobIdentifier: blobIdentifier2,
		},
	}

	useProgressApi := false

	createStoreUploadRequest := openapi_client.CreateStoreUploadRequest{
		UseProgressApi: &useProgressApi,
		Description:    &description,
		BuildId:        &buildId,
		Files:          files,
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

	// Ensure upload is in "completed" status

	expectedUploadStatus := openapi_client.STOREUPLOADSTATUS_COMPLETED
	if getStoreUploadsResponse.Uploads[0].Status != expectedUploadStatus {
		t.Fatalf("GetStoreUpload should return an upload with status %v, but it has status %v", expectedUploadStatus, getStoreUploadsResponse.Uploads[0].Status)
	}

	// Ensure file is in "completed" status

	expectedUploadFileStatus := openapi_client.STOREUPLOADFILESTATUS_COMPLETED
	if getStoreUploadsResponse.Uploads[0].Files[0].Status != expectedUploadFileStatus {
		t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, getStoreUploadsResponse.Uploads[0].Files[0].Status)
	}
}

func TestCreateStoreUploadWithProgressSucceeds(t *testing.T) {

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
	fileName2 := "file2"
	blobIdentifier2 := "blobIdentifier2"
	type2 := openapi_client.StoreFileBlobType("pdb")
	size2 := int64(9)
	contentHash2 := "35c6a7f16428d39c386bd4ebb4c9e0d256bae81634acdf8e65e21bc0abebd0d5" // SHA256 hash of "content2_"

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
		storeUploadId := int32(0)
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
		if getStoreUploadResponse.Files[0].Status != expectedUploadFileStatus {
			t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, getStoreUploadResponse.Files[0].Status)
		}
	}

	// Complete upload of second file

	{
		storeUploadId := int32(0)
		fileId := int32(1)

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

		// Ensure file has transitioned to "completed" status

		expectedUploadFileStatus := openapi_client.STOREUPLOADFILESTATUS_COMPLETED
		if getStoreUploadResponse.Files[0].Status != expectedUploadFileStatus {
			t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, getStoreUploadResponse.Files[0].Status)
		}
	}

	// Complete upload

	{
		storeUploadId := int32(0)

		r, err = apiClient.DefaultApi.MarkStoreUploadCompleted(authContext, storeUploadId, storeId).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("MarkStoreUploadCompleted is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		// Ensure upload has transitioned to "completed" status

		getStoreUploadResponse, r, err := apiClient.DefaultApi.GetStoreUpload(authContext, storeUploadId, storeId).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("GetStoreUpload is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		expectedUploadStatus := openapi_client.STOREUPLOADSTATUS_COMPLETED
		if getStoreUploadResponse.Status != expectedUploadStatus {
			t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadStatus, getStoreUploadResponse.Status)
		}
	}
}

func TestCreateStoreUploadWithProgressAndAbortSucceeds(t *testing.T) {

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
	fileName2 := "file2"
	blobIdentifier2 := "blobIdentifier2"
	type2 := openapi_client.StoreFileBlobType("pdb")
	size2 := int64(9)
	contentHash2 := "35c6a7f16428d39c386bd4ebb4c9e0d256bae81634acdf8e65e21bc0abebd0d5" // SHA256 hash of "content2_"

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
		storeUploadId := int32(0)
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
		if getStoreUploadResponse.Files[0].Status != expectedUploadFileStatus {
			t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, getStoreUploadResponse.Files[0].Status)
		}
	}

	// Abort upload

	{
		storeUploadId := int32(0)

		r, err = apiClient.DefaultApi.MarkStoreUploadAborted(authContext, storeUploadId, storeId).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("MarkStoreUploadAborted is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		// Ensure upload has transitioned to "aborted" status

		getStoreUploadResponse, r, err := apiClient.DefaultApi.GetStoreUpload(authContext, storeUploadId, storeId).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("GetStoreUpload is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		expectedUploadStatus := openapi_client.STOREUPLOADSTATUS_ABORTED
		if getStoreUploadResponse.Status != expectedUploadStatus {
			t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadStatus, getStoreUploadResponse.Status)
		}

		// Ensure first file has transitioned to "aborted" status

		{
			expectedUploadFileStatus := openapi_client.STOREUPLOADFILESTATUS_ABORTED
			if getStoreUploadResponse.Files[0].Status != expectedUploadFileStatus {
				t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, getStoreUploadResponse.Files[0].Status)
			}
		}

		// Ensure second file has transitioned to "aborted" status

		{
			expectedUploadFileStatus := openapi_client.STOREUPLOADFILESTATUS_ABORTED
			if getStoreUploadResponse.Files[1].Status != expectedUploadFileStatus {
				t.Fatalf("GetStoreUpload should return that the second file has status %v, but it has status %v", expectedUploadFileStatus, getStoreUploadResponse.Files[1].Status)
			}
		}
	}
}
