package download_api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	admin_api_client "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-client"
	retry "github.com/hashicorp/go-retryablehttp"
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

func getAdminAPIEndpoint() string {

	adminAPIEndpoint := os.Getenv("ADMIN_API_ENDPOINT")
	if adminAPIEndpoint == "" {
		adminAPIEndpoint = "http://localhost:8080"
	}
	return adminAPIEndpoint
}

func getDownloadAPIEndpoint() string {

	adminAPIEndpoint := os.Getenv("DOWNLOAD_API_ENDPOINT")
	if adminAPIEndpoint == "" {
		adminAPIEndpoint = "http://localhost:8080"
	}
	return adminAPIEndpoint
}

func getDownloadAPIServiceURL(email string, pat string) string {

	downloadAPIEndpoint := getDownloadAPIEndpoint()

	serviceUrl := ""
	if email != "" || pat != "" {
		parts := strings.Split(downloadAPIEndpoint, "://")
		serviceUrl = fmt.Sprintf("%s://%s:%s@%s", parts[0], email, pat, parts[1])
	} else {
		serviceUrl = downloadAPIEndpoint
	}

	return serviceUrl
}

func getAdminAPIClient(email string, pat string) (context.Context, *admin_api_client.APIClient) {

	authContext := context.WithValue(context.Background(), admin_api_client.ContextBasicAuth, admin_api_client.BasicAuth{
		UserName: email,
		Password: pat,
	})

	configuration := admin_api_client.NewConfiguration()
	configuration.Servers[0].URL = getAdminAPIEndpoint()
	api_client := admin_api_client.NewAPIClient(configuration)

	return authContext, api_client
}

func getStores(adminAPIClient *admin_api_client.APIClient, authContext context.Context) ([]string, error) {
	result, _, err := adminAPIClient.DefaultApi.GetStores(authContext).Execute()
	if err != nil {
		return nil, err
	} else {
		return result.Items, err
	}
}

func createStore(adminAPIClient *admin_api_client.APIClient, authContext context.Context, storeId string, acceptStoreAlreadyExists bool) error {
	r, err := adminAPIClient.DefaultApi.CreateStore(authContext, storeId).Execute()
	if err != nil {
		if !acceptStoreAlreadyExists && r.StatusCode != http.StatusConflict {
			return err
		}
	}
	return nil
}

func deleteStore(adminAPIClient *admin_api_client.APIClient, authContext context.Context, storeId string, acceptStoreDoesNotExist bool) error {
	r, err := adminAPIClient.DefaultApi.DeleteStore(authContext, storeId).Execute()
	if err != nil {
		if !acceptStoreDoesNotExist && r.StatusCode != http.StatusNotFound {
			return err
		}
	}
	return nil
}

func ensureTestStoreExists(adminAPIClient *admin_api_client.APIClient, authContext context.Context, storeId string) error {

	err := deleteStore(adminAPIClient, authContext, storeId, true)
	if err != nil {
		return err
	}

	stores1, err := getStores(adminAPIClient, authContext)
	if err != nil {
		return err
	}
	if stringInSlice(storeId, stores1) {
		return errors.New(fmt.Sprintf("Store %v should not be among stores, but is: %v", storeId, stores1))
	}

	err = createStore(adminAPIClient, authContext, storeId, false)
	if err != nil {
		return err
	}

	stores2, err := getStores(adminAPIClient, authContext)
	if err != nil {
		return err
	}
	if !stringInSlice(storeId, stores2) {
		return errors.New(fmt.Sprintf("Store %v should be among stores, but is not: %v", storeId, stores2))
	}

	return nil
}

func populateTestStore(adminAPIClient *admin_api_client.APIClient, authContext context.Context, storeId string) error {

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

	files := []admin_api_client.UploadFileRequest{
		{
			FileName: &fileName1,
			Hash:     &hash1,
		},
		{
			FileName: &fileName2,
			Hash:     &hash2,
		},
	}

	createStoreUploadRequest := admin_api_client.CreateStoreUploadRequest{
		Description: &description,
		BuildId:     &buildId,
		Files:       &files,
	}

	createStoreUploadResponse, r, err := adminAPIClient.DefaultApi.CreateStoreUpload(authContext, storeId).CreateStoreUploadRequest(createStoreUploadRequest).Execute()
	desiredStatusCode := http.StatusOK
	if err != nil {
		return err
	} else if desiredStatusCode != r.StatusCode {
		return errors.New(fmt.Sprintf("CreateStoreUpload is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err))
	}

	retryClient := retry.NewClient()

	for _, fileToUpload := range *createStoreUploadResponse.Files {
		content := contentUploads[*fileToUpload.FileName]

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

func ensureTestStoreExistsAndIsPopulated(email string, pat string, storeId string) error {

	authContext, adminAPIClient := getAdminAPIClient(email, pat)

	if err := ensureTestStoreExists(adminAPIClient, authContext, storeId); err != nil {
		return err
	}
	if err := populateTestStore(adminAPIClient, authContext, storeId); err != nil {
		return err
	}

	return nil
}
