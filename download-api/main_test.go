package hello

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"

	"github.com/hashicorp/go-cleanhttp"
	retry "github.com/hashicorp/go-retryablehttp"
)

func apiRequest(email string, pat string, path string) (*http.Response, error) {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serviceUrl := ""
	if email != "" || pat != "" {
		serviceUrl = fmt.Sprintf("http://%s:%s@localhost:%s", email, pat, port)
	} else {
		serviceUrl = fmt.Sprintf("http://localhost:%s", port)
	}

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

	if _, err := resp.Location(); err == http.ErrNoLocation {
		t.Error("HTTP Response should include a Location header")
	}

	desiredRedirectedURL, _ := url.Parse("http://localhost:9199/b/default-bucket/o/pingme.txt?alt=media")

	location, err := resp.Location()
	if err != http.ErrNoLocation && !reflect.DeepEqual(location, desiredRedirectedURL) {
		t.Errorf("HTTP Response Location: got %v, want %v", location, desiredRedirectedURL)
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
