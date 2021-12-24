package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"

	"github.com/hashicorp/go-cleanhttp"
	retry "github.com/hashicorp/go-retryablehttp"
)

func TestService(t *testing.T) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serviceUrl := os.Getenv("SERVICE_URL")
	if serviceUrl == "" {
		serviceUrl = "http://localhost:" + port
	}

	httpClient := &http.Client{
		Transport: cleanhttp.DefaultPooledTransport(),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	retryClient := retry.NewClient()
	retryClient.HTTPClient = httpClient

	path := "/folder/file"

	req, err := retry.NewRequest(http.MethodGet, serviceUrl+path, nil)
	if err != nil {
		t.Fatalf("retry.NewRequest: %v", err)
	}

	token := os.Getenv("TOKEN")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := retryClient.Do(req)
	if err != nil {
		t.Fatalf("retryClient.Do: %v", err)
	}

	if statusCode := resp.StatusCode; statusCode != http.StatusTemporaryRedirect {
		t.Errorf("HTTP Response status: got %d, want %d", statusCode, http.StatusTemporaryRedirect)
	}

	if _, err := resp.Location(); err == http.ErrNoLocation {
		t.Error("HTTP Response should include a Location header")
	}

	desiredRedirectedURL, _ := url.Parse("http://localhost:9000/folder/file")

	if location, err := resp.Location(); err != http.ErrNoLocation && !reflect.DeepEqual(location, desiredRedirectedURL) {
		t.Errorf("HTTP Response Location: got %v, want %v", location, desiredRedirectedURL)
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll: %v", err)
	}
}
