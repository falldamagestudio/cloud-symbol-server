/*
 * Cloud Symbol Server Admin API
 *
 * This is the API that is used to manage stores and uploads in Cloud Symbol Server
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GetTokenResponse struct {

	// Personal Access Token This token can be used for authentication when accessing non-token related APIs
	Token string `json:"token,omitempty"`

	// Textual description of token Users fill this in to remind themselves the purpose of a token and/or where it is used
	Description string `json:"description,omitempty"`

	// Creation timestamp, in RFC3339 format
	CreationTimestamp string `json:"creationTimestamp,omitempty"`
}

// AssertGetTokenResponseRequired checks if the required fields are not zero-ed
func AssertGetTokenResponseRequired(obj GetTokenResponse) error {
	return nil
}

// AssertRecurseGetTokenResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GetTokenResponse (e.g. [][]GetTokenResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGetTokenResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGetTokenResponse, ok := obj.(GetTokenResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGetTokenResponseRequired(aGetTokenResponse)
	})
}
