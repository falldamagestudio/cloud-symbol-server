/*
 * Cloud Symbol Server Admin API
 *
 * This is the API that is used to manage stores and uploads in Cloud Symbol Server
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PaginationResponse struct {

	Total int32 `json:"total,omitempty"`
}

// AssertPaginationResponseRequired checks if the required fields are not zero-ed
func AssertPaginationResponseRequired(obj PaginationResponse) error {
	return nil
}

// AssertRecursePaginationResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PaginationResponse (e.g. [][]PaginationResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePaginationResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPaginationResponse, ok := obj.(PaginationResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPaginationResponseRequired(aPaginationResponse)
	})
}