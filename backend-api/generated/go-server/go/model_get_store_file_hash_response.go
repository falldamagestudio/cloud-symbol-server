/*
 * Cloud Symbol Server Admin API
 *
 * This is the API that is used to manage stores and uploads in Cloud Symbol Server
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GetStoreFileHashResponse struct {

	Hash string `json:"hash,omitempty"`

	Status StoreFileHashStatus `json:"status,omitempty"`
}

// AssertGetStoreFileHashResponseRequired checks if the required fields are not zero-ed
func AssertGetStoreFileHashResponseRequired(obj GetStoreFileHashResponse) error {
	return nil
}

// AssertRecurseGetStoreFileHashResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GetStoreFileHashResponse (e.g. [][]GetStoreFileHashResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGetStoreFileHashResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGetStoreFileHashResponse, ok := obj.(GetStoreFileHashResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGetStoreFileHashResponseRequired(aGetStoreFileHashResponse)
	})
}
