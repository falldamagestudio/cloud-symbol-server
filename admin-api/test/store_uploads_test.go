package admin_api

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
	uploadId := "999999999"

	_, r, err := apiClient.DefaultApi.GetStoreUpload(authContext, uploadId, storeId).Execute()
	desiredStatusCode := http.StatusUnauthorized
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreUpload with invalid email/PAT is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreUploadIdsForStoreThatDoesNotExistFails(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	invalidStoreId := "invalidstoreid"

	_, r, err := apiClient.DefaultApi.GetStoreUploadIds(authContext, invalidStoreId).Execute()
	desiredStatusCode := http.StatusNotFound
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreUploadIds with invalid store ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreUploadIdsForStoreExistsSucceeds(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"

	if err := ensureTestStoreExists(apiClient, authContext, storeId); err != nil {
		t.Fatalf("ensureTestStoreExists failed: %v", err)
	}

	_, r, err := apiClient.DefaultApi.GetStoreUploadIds(authContext, storeId).Execute()
	desiredStatusCode := http.StatusOK
	if err != nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreUploadIds with valid store ID is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestGetStoreUploadThatDoesNotExistFails(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "example"
	uploadId := "999999999"

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

	useProgressApi := false

	createStoreUploadRequest := openapi_client.CreateStoreUploadRequest{
		UseProgressApi: &useProgressApi,
		Description:    &description,
		BuildId:        &buildId,
		Files:          &files,
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
	if !reflect.DeepEqual(expectedStoreUploadIds, getStoreUploadIdsResponse.Items) {
		t.Fatalf("Expected GetStoreUploadIds is expected to return %v, but returned %v", expectedStoreUploadIds, getStoreUploadIdsResponse.Items)
	}

	// Ensure upload is in "completed" status

	storeUploadId := "0"

	getStoreUploadResponse, r, err := apiClient.DefaultApi.GetStoreUpload(authContext, storeUploadId, storeId).Execute()
	desiredStatusCode = http.StatusOK
	if err != nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("GetStoreUpload is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}

	expectedUploadStatus := "completed"
	if *getStoreUploadResponse.Status != expectedUploadStatus {
		t.Fatalf("GetStoreUpload should return an upload with status %v, but it has status %v", expectedUploadStatus, *getStoreUploadResponse.Status)
	}

	// Ensure file is in "completed" status

	expectedUploadFileStatus := "completed"
	if *(*getStoreUploadResponse.Files)[0].Status != expectedUploadFileStatus {
		t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, *(*getStoreUploadResponse.Files)[0].Status)
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

	useProgressApi := true
	createStoreUploadRequest := openapi_client.CreateStoreUploadRequest{
		Description:    &description,
		BuildId:        &buildId,
		Files:          &files,
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
	if !reflect.DeepEqual(expectedStoreUploadIds, getStoreUploadIdsResponse.Items) {
		t.Fatalf("Expected GetStoreUploadIds is expected to return %v, but returned %v", expectedStoreUploadIds, getStoreUploadIdsResponse.Items)
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
	if *(*getStoreUploadResponse.Files)[0].Status != expectedUploadFileStatus {
		t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, *(*getStoreUploadResponse.Files)[0].Status)
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
		if *(*getStoreUploadResponse.Files)[0].Status != expectedUploadFileStatus {
			t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, *(*getStoreUploadResponse.Files)[0].Status)
		}
	}

	// Complete upload of second file

	{
		fileId := int32(1)

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

		// Ensure file has transitioned to "completed" status

		expectedUploadFileStatus := "completed"
		if *(*getStoreUploadResponse.Files)[0].Status != expectedUploadFileStatus {
			t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, *(*getStoreUploadResponse.Files)[0].Status)
		}
	}

	// Complete upload

	{
		r, err = apiClient.DefaultApi.MarkStoreUploadCompleted(authContext, storeUploadId, storeId).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("MarkStoreUploadCompleted is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		// Ensure upload has transitioned to "completed" status

		storeUploadId := "0"

		getStoreUploadResponse, r, err := apiClient.DefaultApi.GetStoreUpload(authContext, storeUploadId, storeId).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("GetStoreUpload is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		expectedUploadStatus := "completed"
		if *getStoreUploadResponse.Status != expectedUploadStatus {
			t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadStatus, *getStoreUploadResponse.Status)
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

	useProgressApi := true
	createStoreUploadRequest := openapi_client.CreateStoreUploadRequest{
		Description:    &description,
		BuildId:        &buildId,
		Files:          &files,
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
	if !reflect.DeepEqual(expectedStoreUploadIds, getStoreUploadIdsResponse.Items) {
		t.Fatalf("Expected GetStoreUploadIds is expected to return %v, but returned %v", expectedStoreUploadIds, getStoreUploadIdsResponse.Items)
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
	if *(*getStoreUploadResponse.Files)[0].Status != expectedUploadFileStatus {
		t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, *(*getStoreUploadResponse.Files)[0].Status)
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
		if *(*getStoreUploadResponse.Files)[0].Status != expectedUploadFileStatus {
			t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, *(*getStoreUploadResponse.Files)[0].Status)
		}
	}

	// Abort upload

	{
		r, err = apiClient.DefaultApi.MarkStoreUploadAborted(authContext, storeUploadId, storeId).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("MarkStoreUploadAborted is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		// Ensure upload has transitioned to "aborted" status

		storeUploadId := "0"

		getStoreUploadResponse, r, err := apiClient.DefaultApi.GetStoreUpload(authContext, storeUploadId, storeId).Execute()
		desiredStatusCode = http.StatusOK
		if err != nil || desiredStatusCode != r.StatusCode {
			t.Fatalf("GetStoreUpload is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
		}

		expectedUploadStatus := "aborted"
		if *getStoreUploadResponse.Status != expectedUploadStatus {
			t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadStatus, *getStoreUploadResponse.Status)
		}

		// Ensure first file remains in "uploaded" status

		{
			expectedUploadFileStatus := "uploaded"
			if *(*getStoreUploadResponse.Files)[0].Status != expectedUploadFileStatus {
				t.Fatalf("GetStoreUpload should return that the first file has status %v, but it has status %v", expectedUploadFileStatus, *(*getStoreUploadResponse.Files)[0].Status)
			}
		}

		// Ensure second file has transitioned to "aborted" status

		{
			expectedUploadFileStatus := "aborted"
			if *(*getStoreUploadResponse.Files)[1].Status != expectedUploadFileStatus {
				t.Fatalf("GetStoreUpload should return that the second file has status %v, but it has status %v", expectedUploadFileStatus, *(*getStoreUploadResponse.Files)[1].Status)
			}
		}
	}
}
