/*
 * Cloud Symbol Store Uplaod API
 *
 * This is the API that is used to upload symbols to Cloud Symbol Store
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UploadTransactionResponse struct {

	Id string `json:"id,omitempty"`

	Files []UploadFileResponse `json:"files,omitempty"`
}

// AssertUploadTransactionResponseRequired checks if the required fields are not zero-ed
func AssertUploadTransactionResponseRequired(obj UploadTransactionResponse) error {
	for _, el := range obj.Files {
		if err := AssertUploadFileResponseRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseUploadTransactionResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UploadTransactionResponse (e.g. [][]UploadTransactionResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUploadTransactionResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUploadTransactionResponse, ok := obj.(UploadTransactionResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUploadTransactionResponseRequired(aUploadTransactionResponse)
	})
}
