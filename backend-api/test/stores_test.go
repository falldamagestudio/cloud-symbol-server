package backend_api

import (
	"net/http"
	"testing"
)

func TestCreateAndDestroyStore(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "test-store"

	err := deleteStore(apiClient, authContext, storeId, true)
	if err != nil {
		t.Fatalf("DeleteStore failed: %v", err)
	}

	stores1, err := getStores(apiClient, authContext)
	if err != nil {
		t.Fatalf("GetStores failed: %v", err)
	}
	if stringInSlice(storeId, stores1) {
		t.Fatalf("Store %v should not be among stores, but is: %v", storeId, stores1)
	}

	err = createStore(apiClient, authContext, storeId, false)
	if err != nil {
		t.Fatalf("CreateStore failed: %v", err)
	}

	stores2, err := getStores(apiClient, authContext)
	if err != nil {
		t.Fatalf("GetStores failed: %v", err)
	}
	if !stringInSlice(storeId, stores2) {
		t.Fatalf("Store %v should be among stores, but is not: %v", storeId, stores2)
	}

	err = deleteStore(apiClient, authContext, storeId, false)
	if err != nil {
		t.Fatalf("DeleteStore failed: %v", err)
	}

	stores3, err := getStores(apiClient, authContext)
	if err != nil {
		t.Fatalf("GetStores failed: %v", err)
	}
	if stringInSlice(storeId, stores3) {
		t.Fatalf("Store %v should not be among stores, but is: %v", storeId, stores3)
	}
}

func TestCreateStoreSucceedsIfStoreDoesNotExist(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "test-store"

	err := ensureTestStoreDoesNotExist(apiClient, authContext, storeId)
	if err != nil {
		t.Fatalf("ensureTestStoreDoesNotExist failed: %v", err)
	}

	r, err := apiClient.DefaultApi.CreateStore(authContext, storeId).Execute()
	desiredStatusCode := http.StatusOK
	if err != nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("CreateStore when store doesn't already exist is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestCreateStoreFailsIfStoreAlreadyExists(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "test-store"

	err := ensureTestStoreExists(apiClient, authContext, storeId)
	if err != nil {
		t.Fatalf("ensureTestStoreExists failed: %v", err)
	}

	r, err := apiClient.DefaultApi.CreateStore(authContext, storeId).Execute()
	desiredStatusCode := http.StatusConflict
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("CreateStore when store already exists is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestDeleteStoreSucceedsIfStoreAlreadyExists(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "test-store"

	err := ensureTestStoreExists(apiClient, authContext, storeId)
	if err != nil {
		t.Fatalf("ensureTestStoreExists failed: %v", err)
	}

	r, err := apiClient.DefaultApi.DeleteStore(authContext, storeId).Execute()
	desiredStatusCode := http.StatusOK
	if err != nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("DeleteStore when store already exists is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}

func TestDeleteStoreFailsIfStoreDoesNotAlreadyExist(t *testing.T) {

	email, pat := getTestEmailAndPat()

	authContext, apiClient := getAPIClient(email, pat)

	storeId := "test-store"

	err := ensureTestStoreDoesNotExist(apiClient, authContext, storeId)
	if err != nil {
		t.Fatalf("ensureTestStoreDoesNotExist failed: %v", err)
	}

	r, err := apiClient.DefaultApi.DeleteStore(authContext, storeId).Execute()
	desiredStatusCode := http.StatusNotFound
	if err == nil || desiredStatusCode != r.StatusCode {
		t.Fatalf("DeleseStore when store does not already exist is expected to give HTTP status code %v, but gave %v as response (err = %v)", desiredStatusCode, r.StatusCode, err)
	}
}
