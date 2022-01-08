/*
 * Cloud Symbol Store Uplaod API
 *
 * This is the API that is used to upload symbols to Cloud Symbol Store
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UploadFileResponse struct {

	FileName string `json:"fileName,omitempty"`

	Hash string `json:"hash,omitempty"`

	Url string `json:"url,omitempty"`
}

// AssertUploadFileResponseRequired checks if the required fields are not zero-ed
func AssertUploadFileResponseRequired(obj UploadFileResponse) error {
	return nil
}

// AssertRecurseUploadFileResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UploadFileResponse (e.g. [][]UploadFileResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUploadFileResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUploadFileResponse, ok := obj.(UploadFileResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUploadFileResponseRequired(aUploadFileResponse)
	})
}