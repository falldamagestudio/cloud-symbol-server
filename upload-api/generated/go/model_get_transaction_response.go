/*
 * Cloud Symbol Store Uplaod API
 *
 * This is the API that is used to upload symbols to Cloud Symbol Store
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GetTransactionResponse struct {

	Description string `json:"description,omitempty"`

	BuildId string `json:"buildId,omitempty"`

	Timestamp string `json:"timestamp,omitempty"`

	Files []GetFileResponse `json:"files,omitempty"`
}

// AssertGetTransactionResponseRequired checks if the required fields are not zero-ed
func AssertGetTransactionResponseRequired(obj GetTransactionResponse) error {
	for _, el := range obj.Files {
		if err := AssertGetFileResponseRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseGetTransactionResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GetTransactionResponse (e.g. [][]GetTransactionResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGetTransactionResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGetTransactionResponse, ok := obj.(GetTransactionResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGetTransactionResponseRequired(aGetTransactionResponse)
	})
}