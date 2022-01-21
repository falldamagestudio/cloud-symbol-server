package download_api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/go-cleanhttp"
	retry "github.com/hashicorp/go-retryablehttp"
)

func getServiceUrl(email string, pat string) string {

	downloadAPIEndpoint := os.Getenv("DOWNLOAD_API_ENDPOINT")
	if downloadAPIEndpoint == "" {
		downloadAPIEndpoint = "http://localhost:8080"
	}

	serviceUrl := ""
	if email != "" || pat != "" {
		parts := strings.Split(downloadAPIEndpoint, "://")
		serviceUrl = fmt.Sprintf("%s://%s:%s@%s", parts[0], email, pat, parts[1])
	} else {
		serviceUrl = downloadAPIEndpoint
	}

	return serviceUrl
}

func apiRequest(email string, pat string, path string) (*http.Response, error) {

	serviceUrl := getServiceUrl(email, pat)

	httpClient := &http.Client{
		Transport: cleanhttp.DefaultPooledTransport(),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	retryClient := retry.NewClient()
	retryClient.HTTPClient = httpClient

	req, err := retry.NewRequest(http.MethodGet, serviceUrl+"/"+path, nil)
	if err != nil {
		//t.Fatalf("retry.NewRequest: %v", err)
		return nil, err
	}

	token := os.Getenv("TOKEN")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := retryClient.Do(req)
	if err != nil {
		//t.Fatalf("retryClient.Do: %v", err)
		return nil, err
	}

	return resp, nil
}

func fileRequest(url string) (*http.Response, error) {

	httpClient := &http.Client{
		Transport: cleanhttp.DefaultPooledTransport(),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	retryClient := retry.NewClient()
	retryClient.HTTPClient = httpClient

	req, err := retry.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		//t.Fatalf("retry.NewRequest: %v", err)
		return nil, err
	}

	resp, err := retryClient.Do(req)
	if err != nil {
		//t.Fatalf("retryClient.Do: %v", err)
		return nil, err
	}

	return resp, nil
}

func TestAccessFileThatExists(t *testing.T) {

	email := "testuser"
	pat := "testpat"
	path := "pingme.txt"

	// Request contents of a file that exists

	resp, err := apiRequest(email, pat, path)

	if err != nil {
		t.Errorf("Error in request: %v", err)
	}

	// API response should be a HTTP 307 redirect URL, to a download location for the actual file

	if statusCode := resp.StatusCode; statusCode != http.StatusTemporaryRedirect {
		t.Errorf("HTTP Response status: got %d, want %d", statusCode, http.StatusTemporaryRedirect)
	}

	// Unsigned URLs are on the format,
	// http://<host>:<port>/[storage/v1/]b/<bucketname>/o/<objectname>?alt=media
	desiredUnsignedURLSuffix := "pingme.txt?alt=media"

	// Signed URLs are on the format,
	// https://<host>/<bucketname>/<objectname>?<bunch of key-value pairs, including an 'Expires' header that is typically first>
	desiredSignedURLComponent := "?Expires="

	location, err := resp.Location()

	if err == http.ErrNoLocation {
		t.Error("HTTP Response should include a Location header")
	} else if !strings.HasSuffix(location.String(), desiredUnsignedURLSuffix) && !strings.Contains(location.String(), desiredSignedURLComponent) {
		t.Errorf("HTTP Response Location: got %v, want string to either have unsigned suffix %v or signed substring %v", location, desiredUnsignedURLSuffix, desiredSignedURLComponent)
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll: %v", err)
	}

	// Download actual file

	resp, err = fileRequest(location.String())
	if err != nil {
		t.Errorf("Error in request: %v", err)
	}

	// Downloding actual file should result in a HTTP 200 response

	if statusCode := resp.StatusCode; statusCode != http.StatusOK {
		t.Errorf("HTTP Response status: got %d, want %d", statusCode, http.StatusOK)
	}
}

func TestAccessFileThatDoesNotExist(t *testing.T) {

	email := "testuser"
	pat := "testpat"
	path := "pingme2.txt"

	// Request contents of a file that exists

	resp, err := apiRequest(email, pat, path)

	if err != nil {
		t.Errorf("Error in request: %v", err)
	}

	// API response should be a HTTP 404

	if statusCode := resp.StatusCode; statusCode != http.StatusNotFound {
		t.Errorf("HTTP Response status: got %d, want %d", statusCode, http.StatusNotFound)
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll: %v", err)
	}
}

func TestAccessWithoutBasicAuthFails(t *testing.T) {

	path := "pingme.txt"

	// Request contents of a file that exists

	resp, err := apiRequest("", "", path)

	if err != nil {
		t.Errorf("Error in request: %v", err)
	}

	// API response should be a HTTP 401 Forbidden

	if statusCode := resp.StatusCode; statusCode != http.StatusUnauthorized {
		t.Errorf("HTTP Response status: got %d, want %d", statusCode, http.StatusUnauthorized)
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll: %v", err)
	}
}

func TestAccessWithInvalidPATFails(t *testing.T) {

	email := "testuser"
	pat := "invalidpat"
	path := "pingme.txt"

	// Request contents of a file that exists

	resp, err := apiRequest(email, pat, path)

	if err != nil {
		t.Errorf("Error in request: %v", err)
	}

	// API response should be a HTTP 401 Forbidden

	if statusCode := resp.StatusCode; statusCode != http.StatusUnauthorized {
		t.Errorf("HTTP Response status: got %d, want %d", statusCode, http.StatusUnauthorized)
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll: %v", err)
	}
}
