package upload_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

func getServiceUrl(email string, pat string, path string) string {

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

func TestUploadTransaction(t *testing.T) {

	email := "testuser"
	pat := "testpat"
	path := "pingme.txt"

	serviceUrl := getServiceUrl(email, pat, path)

	description := "test upload"
	buildId := "test build id"
	files := []UploadFileRequest{
		{
			FileName: "file1",
			Hash:     "hash1",
		},
		{
			FileName: "file2",
			Hash:     "hash2",
		},
	}

	uploadTransaction := &UploadTransactionRequest{
		Description: description,
		BuildId:     buildId,
		Files:       files,
	}

	requestBody, err := json.Marshal(uploadTransaction)
	if err != nil {
		t.Fatalf("Error when marshalling json to text: %v", err)
	}

	response, err := retryablehttp.Post(serviceUrl, "application/json", requestBody)
	defer response.Body.Close()

	if err != nil {
		t.Fatalf("Error in request: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		responseBody, _ := ioutil.ReadAll(response.Body)
		t.Fatalf("HTTP POST request failed: StatusCode = %v, Body = %v", response.StatusCode, string(responseBody))
	}

	uploadTransactionResponse := UploadTransactionResponse{}
	json.NewDecoder(response.Body).Decode(uploadTransactionResponse)

}
