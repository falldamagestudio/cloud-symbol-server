/*
 * Cloud Symbol Server Admin API
 *
 * This is the API that is used to manage stores and uploads in Cloud Symbol Server
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GetStoresResponse struct {
	Items []string
}

// AssertGetStoresResponseRequired checks if the required fields are not zero-ed
func AssertGetStoresResponseRequired(obj GetStoresResponse) error {
	return nil
}

// AssertRecurseGetStoresResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GetStoresResponse (e.g. [][]GetStoresResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGetStoresResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGetStoresResponse, ok := obj.(GetStoresResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGetStoresResponseRequired(aGetStoresResponse)
	})
}