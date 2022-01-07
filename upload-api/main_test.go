package upload_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	openapi "github.com/falldamagestudio/cloud-symbol-store/upload-api/generated/go"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

func getServiceUrl(email string, pat string) string {

	uploadAPIProtocol := os.Getenv("UPLOAD_API_PROTOCOL")
	if uploadAPIProtocol == "" {
		uploadAPIProtocol = "http"
	}

	uploadAPIHost := os.Getenv("UPLOAD_API_HOST")
	if uploadAPIHost == "" {
		uploadAPIHost = "localhost:8080"
	}

	serviceUrl := ""
	if email != "" || pat != "" {
		serviceUrl = fmt.Sprintf("%s://%s:%s@%s", uploadAPIProtocol, email, pat, uploadAPIHost)
	} else {
		serviceUrl = fmt.Sprintf("%s://%s", uploadAPIProtocol, uploadAPIHost)
	}

	return serviceUrl
}

func TestGetTransactionWithInvalidCredentialsFails(t *testing.T) {

	email := "invalidemail"
	pat := "invalidpat"

	serviceUrl := getServiceUrl(email, pat)

	path := "/transactions/nonexistentid"

	response, err := retryablehttp.Get(serviceUrl + path)
	defer response.Body.Close()

	if err != nil {
		t.Fatalf("Error in request: %v", err)
	}

	if response.StatusCode != http.StatusUnauthorized {
		responseBody, _ := ioutil.ReadAll(response.Body)
		t.Fatalf("HTTP GET request failed: StatusCode expected %v, was %v, Body = %v", http.StatusUnauthorized, response.StatusCode, string(responseBody))
	}
}

func TestGetTransactionThatDoesNotExistFails(t *testing.T) {

	email := "testuser"
	pat := "testpat"

	serviceUrl := getServiceUrl(email, pat)

	path := "/transactions/nonexistentid"

	response, err := retryablehttp.Get(serviceUrl + path)
	defer response.Body.Close()

	if err != nil {
		t.Fatalf("Error in request: %v", err)
	}

	if response.StatusCode != http.StatusNotFound {
		responseBody, _ := ioutil.ReadAll(response.Body)
		t.Fatalf("HTTP GET request failed: StatusCode expected %v, was %v, Body = %v", http.StatusNotFound, response.StatusCode, string(responseBody))
	}
}

func TestUploadTransactionSucceeds(t *testing.T) {

	email := "testuser"
	pat := "testpat"

	path := "/transactions"

	serviceUrl := getServiceUrl(email, pat)

	description := "test upload"
	buildId := "test build id"
	files := []openapi.UploadFileRequest{
		{
			FileName: "file1",
			Hash:     "hash1",
		},
		{
			FileName: "file2",
			Hash:     "hash2",
		},
	}

	uploadTransaction := &openapi.UploadTransactionRequest{
		Description: description,
		BuildId:     buildId,
		Files:       files,
	}

	requestBody, err := json.Marshal(uploadTransaction)
	if err != nil {
		t.Fatalf("Error when marshalling json to text: %v", err)
	}

	response, err := retryablehttp.Post(serviceUrl+path, "application/json", requestBody)
	defer response.Body.Close()

	if err != nil {
		t.Fatalf("Error in request: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		responseBody, _ := ioutil.ReadAll(response.Body)
		t.Fatalf("HTTP POST request failed: StatusCode = %v, Body = %v", response.StatusCode, string(responseBody))
	}

	uploadTransactionResponse := openapi.UploadTransactionResponse{}
	json.NewDecoder(response.Body).Decode(uploadTransactionResponse)

}
