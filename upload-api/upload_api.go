package upload_api

type UploadFileRequest struct {
	FileName string `json:"filename"`
	Hash     string `json:"hash"`
}

type UploadTransactionRequest struct {
	Description string              `json:"description"`
	BuildId     string              `json:"buildid"`
	Files       []UploadFileRequest `json:"files"`
}

type UploadFileResponse struct {
	FileName string `json:"filename"`
	Hash     string `json:"hash"`
	Url      string `json:"url"`
}

type UploadTransactionResponse struct {
	Files []UploadFileResponse `json:"files"`
}
