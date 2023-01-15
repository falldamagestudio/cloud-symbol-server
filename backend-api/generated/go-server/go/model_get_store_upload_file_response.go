/*
 * Cloud Symbol Server Admin API
 *
 * This is the API that is used to manage stores and uploads in Cloud Symbol Server
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GetStoreUploadFileResponse struct {

	FileName string `json:"fileName"`

	Hash string `json:"hash"`

	Status StoreUploadFileStatus `json:"status"`
}

// AssertGetStoreUploadFileResponseRequired checks if the required fields are not zero-ed
func AssertGetStoreUploadFileResponseRequired(obj GetStoreUploadFileResponse) error {
	elements := map[string]interface{}{
		"fileName": obj.FileName,
		"hash": obj.Hash,
		"status": obj.Status,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseGetStoreUploadFileResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GetStoreUploadFileResponse (e.g. [][]GetStoreUploadFileResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGetStoreUploadFileResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGetStoreUploadFileResponse, ok := obj.(GetStoreUploadFileResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGetStoreUploadFileResponseRequired(aGetStoreUploadFileResponse)
	})
}