/*
 * Cloud Symbol Store Uplaod API
 *
 * This is the API that is used to upload symbols to Cloud Symbol Store
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UploadFileRequest struct {

	FileName string `json:"fileName,omitempty"`

	Hash string `json:"hash,omitempty"`
}

// AssertUploadFileRequestRequired checks if the required fields are not zero-ed
func AssertUploadFileRequestRequired(obj UploadFileRequest) error {
	return nil
}

// AssertRecurseUploadFileRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UploadFileRequest (e.g. [][]UploadFileRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUploadFileRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUploadFileRequest, ok := obj.(UploadFileRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUploadFileRequestRequired(aUploadFileRequest)
	})
}