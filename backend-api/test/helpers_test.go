package backend_api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	retry "github.com/hashicorp/go-retryablehttp"

	openapi_client "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-client"
)

func getTestEmailAndPat() (string, string) {
	email := os.Getenv("TEST_EMAIL")
	pat := os.Getenv("TEST_PAT")
	return email, pat
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getBackendAPIEndpoint() string {

	backendAPIEndpoint := os.Getenv("BACKEND_API_ENDPOINT")
	if backendAPIEndpoint == "" {
		backendAPIEndpoint = "http://localhost:8080"
	}
	return backendAPIEndpoint
}

func getDownloadAPIServiceURL(email string, pat string) string {

	backendAPIEndpoint := getBackendAPIEndpoint()

	downloadAPIEndpoint := fmt.Sprintf("%v/download", backendAPIEndpoint)

	serviceUrl := ""
	if email != "" || pat != "" {
		parts := strings.Split(downloadAPIEndpoint, "://")
		serviceUrl = fmt.Sprintf("%s://%s:%s@%s", parts[0], email, pat, parts[1])
	} else {
		serviceUrl = downloadAPIEndpoint
	}

	return serviceUrl
}

func getAPIClient(email string, pat string) (context.Context, *openapi_client.APIClient) {

	authContext := context.WithValue(context.Background(), openapi_client.ContextBasicAuth, openapi_client.BasicAuth{
		UserName: email,
		Password: pat,
	})

	configuration := openapi_client.NewConfiguration()
	configuration.Servers[0].URL = getBackendAPIEndpoint()
	api_client := openapi_client.NewAPIClient(configuration)

	return authContext, api_client
}

func getStores(apiClient *openapi_client.APIClient, authContext context.Context) ([]string, error) {
	result, _, err := apiClient.DefaultApi.GetStores(authContext).Execute()
	if err != nil {
		return nil, err
	} else {
		return result, err
	}
}

func createStore(apiClient *openapi_client.APIClient, authContext context.Context, storeId string, acceptStoreAlreadyExists bool) error {
	r, err := apiClient.DefaultApi.CreateStore(authContext, storeId).Execute()
	if err != nil {
		if !acceptStoreAlreadyExists && r.StatusCode != http.StatusConflict {
			return err
		}
	}
	return nil
}

func deleteStore(apiClient *openapi_client.APIClient, authContext context.Context, storeId string, acceptStoreDoesNotExist bool) error {
	r, err := apiClient.DefaultApi.DeleteStore(authContext, storeId).Execute()
	if err != nil {
		if !acceptStoreDoesNotExist && r.StatusCode != http.StatusNotFound {
			return err
		}
	}
	return nil
}

func ensureTestStoreExists(apiClient *openapi_client.APIClient, authContext context.Context, storeId string) error {

	err := deleteStore(apiClient, authContext, storeId, true)
	if err != nil {
		return err
	}

	stores1, err := getStores(apiClient, authContext)
	if err != nil {
		return err
	}
	if stringInSlice(storeId, stores1) {
		return errors.New(fmt.Sprintf("Store %v should not be among stores, but is: %v", storeId, stores1))
	}

	err = createStore(apiClient, authContext, storeId, false)
	if err != nil {
		return err
	}

	stores2, err := getStores(apiClient, authContext)
	if err != nil {
		return err
	}
	if !stringInSlice(storeId, stores2) {
		return errors.New(fmt.Sprintf("Store %v should be among stores, but is not: %v", storeId, stores2))
	}

	return nil
}

func ensureTestStoreDoesNotExist(apiClient *openapi_client.APIClient, authContext context.Context, storeId string) error {

	err := deleteStore(apiClient, authContext, storeId, true)
	if err != nil {
		return err
	}

	stores1, err := getStores(apiClient, authContext)
	if err != nil {
		return err
	}
	if stringInSlice(storeId, stores1) {
		return errors.New(fmt.Sprintf("Store %v should not be among stores, but is: %v", storeId, stores1))
	}

	return nil
}

func populateTestStore(adminAPIClient *openapi_client.APIClient, authContext context.Context, storeId string) error {

	description := "test upload"
	buildId := "test build id"

	fileName1 := "file1"
	hash1 := "hash1"
	content1 := "content1"
	fileName2 := "file2"
	hash2 := "hash2"
	content2 := "content2"

	contentUploads := map[string]string{
		fileName1: content1,
		fileName2: content2,
	}

	files := []openapi_client.CreateStoreUploadFileRequest{
		{
			FileName: fileName1,
			Hash:     hash1,
		},
		{
			FileName: fileName2,
			Hash:     hash2,
		},
	}

	createStoreUploadRequest := openapi_client.CreateStoreUploadRequest{
		Description: &description,
		BuildId:     &buildId,
		Files:       files,
	}

	createStoreUploadResponse, r, err := adminAPIClient.DefaultApi.CreateStoreUpload(authContext, storeId).CreateStoreUploadRequest(createStoreUploadRequest).Execute()
	desiredStatusCode := http.StatusOK
	if err != nil {
		return err
	} else if desiredStatusCode != r.StatusCode {
		return errors.New(fmt.Sprintf("CreateStoreUpload is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err))
	}

	retryClient := retry.NewClient()

	for _, fileToUpload := range createStoreUploadResponse.Files {
		content := contentUploads[fileToUpload.FileName]

		request, err := retry.NewRequest(http.MethodPut, *fileToUpload.Url, []byte(content))
		if err != nil {
			return err
		}
		_, err = retryClient.Do(request)
		if err != nil {
			return err
		}
	}

	return nil
}

func ensureTestStoreExistsAndIsPopulated(apiClient *openapi_client.APIClient, authContext context.Context, storeId string) error {

	if err := ensureTestStoreExists(apiClient, authContext, storeId); err != nil {
		return err
	}
	if err := populateTestStore(apiClient, authContext, storeId); err != nil {
		return err
	}

	return nil
}

type TestUpload struct {
	BuildId     string
	Description string
	Files       []TestUploadFile
}

type TestUploadFile struct {
	FileName string
	Hash     string
	Content  string
}

var (
	exampleUpload1 = TestUpload{
		BuildId:     "example upload Build ID 1",
		Description: "example upload description 1",
		Files: []TestUploadFile{
			{
				FileName: "file1",
				Hash:     "hash1",
				Content:  "content1",
			},
			{
				FileName: "file2",
				Hash:     "hash2",
				Content:  "content2",
			},
		},
	}

	exampleUpload2 = TestUpload{
		BuildId:     "example upload Build ID 2",
		Description: "example upload description 2",
		Files: []TestUploadFile{
			{
				FileName: "file1",
				Hash:     "hash1",
				Content:  "content1",
			},
			{
				FileName: "file3",
				Hash:     "hash3",
				Content:  "content3",
			},
		},
	}
)

func upload(apiClient *openapi_client.APIClient, authContext context.Context, storeId string, testUpload *TestUpload) (string, error) {

	files := make([]openapi_client.CreateStoreUploadFileRequest, len(testUpload.Files))

	useProgressApi := true
	createStoreUploadRequest := &openapi_client.CreateStoreUploadRequest{
		BuildId:        &testUpload.BuildId,
		Description:    &testUpload.Description,
		Files:          files,
		UseProgressApi: &useProgressApi,
	}

	for fileIndex := range testUpload.Files {
		sourceFile := &testUpload.Files[fileIndex]
		targetFile := &((*createStoreUploadRequest).Files)[fileIndex]

		targetFile.FileName = sourceFile.FileName
		targetFile.Hash = sourceFile.Hash
	}

	createStoreUploadResponse, _, err := apiClient.DefaultApi.CreateStoreUpload(authContext, storeId).CreateStoreUploadRequest(*createStoreUploadRequest).Execute()
	if err != nil {
		return "", err
	}

	// Upload individual files, and mark them as uploaded

	storeUploadId := createStoreUploadResponse.Id

	for fileIndex := range createStoreUploadRequest.Files {

		uploadUrlPtr := (createStoreUploadResponse.Files)[fileIndex].Url
		if uploadUrlPtr != nil {

			// File did not already exist; server has supplied an upload URL

			uploadUrl := *uploadUrlPtr
			content := []byte(testUpload.Files[fileIndex].Content)

			// Upload file to GCS using server-supplied upload URL

			uploadRequest, err := http.NewRequest(http.MethodPut, uploadUrl, bytes.NewReader(content))
			if err != nil {
				return "", err
			}

			client := http.Client{}
			uploadResponse, err := client.Do(uploadRequest)
			if err != nil {
				return "", err
			}

			defer uploadResponse.Body.Close()
			_, err = io.ReadAll(uploadResponse.Body)
			if err != nil {
				return "", err
			}

			// Mark file as uploaded

			_, err = apiClient.DefaultApi.MarkStoreUploadFileUploaded(authContext, storeUploadId, storeId, int32(fileIndex)).Execute()
			if err != nil {
				return "", err
			}
		}
	}

	// Complete upload

	{
		_, err = apiClient.DefaultApi.MarkStoreUploadCompleted(authContext, storeUploadId, storeId).Execute()
		if err != nil {
			return "", err
		}
	}

	return storeUploadId, nil
}
