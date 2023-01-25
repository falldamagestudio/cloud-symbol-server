/*
 * Cloud Symbol Server Admin API
 *
 * This is the API that is used to manage stores and uploads in Cloud Symbol Server
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UploadFileResponse struct {

	FileName string `json:"fileName"`

	BlobIdentifier string `json:"blobIdentifier"`

	// Short-lived signed URL where the client should upload the file to, or blank if the file already exists in the storage backend
	Url string `json:"url,omitempty"`

	Hash string `json:"hash,omitempty"`
}

// AssertUploadFileResponseRequired checks if the required fields are not zero-ed
func AssertUploadFileResponseRequired(obj UploadFileResponse) error {
	elements := map[string]interface{}{
		"fileName": obj.FileName,
		"blobIdentifier": obj.BlobIdentifier,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

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
