package admin_api_test

import (
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
