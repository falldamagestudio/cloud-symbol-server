/*
 * Cloud Symbol Store Uplaod API
 *
 * This is the API that is used to upload symbols to Cloud Symbol Store
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UploadTransactionRequest struct {

	Description string `json:"description,omitempty"`

	BuildId string `json:"buildId,omitempty"`

	Files []UploadFileRequest `json:"files,omitempty"`
}

// AssertUploadTransactionRequestRequired checks if the required fields are not zero-ed
func AssertUploadTransactionRequestRequired(obj UploadTransactionRequest) error {
	for _, el := range obj.Files {
		if err := AssertUploadFileRequestRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseUploadTransactionRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UploadTransactionRequest (e.g. [][]UploadTransactionRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUploadTransactionRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUploadTransactionRequest, ok := obj.(UploadTransactionRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUploadTransactionRequestRequired(aUploadTransactionRequest)
	})
}