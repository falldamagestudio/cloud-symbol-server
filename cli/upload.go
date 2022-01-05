package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	retryablehttp "github.com/hashicorp/go-retryablehttp"

	upload_api "github.com/falldamagestudio/cloud-symbol-store/upload-api"
)

type FileWithHash struct {
	FileWithPath    string
	FileWithoutPath string
	Hash            string
}

func GetFilesWithHashes(fileNames []string) ([]FileWithHash, error) {

	filesWithHashes := []FileWithHash{}

	for _, fileName := range fileNames {

		hash, err := GetPdbHash(fileName)

		if err != nil {
			log.Printf("Error while parsing PDB: %v\n", err)
			return nil, err
		}

		filesWithHashes = append(filesWithHashes, FileWithHash{
			FileWithPath:    fileName,
			FileWithoutPath: path.Base(fileName),
			Hash:            *hash,
		})
	}

	return filesWithHashes, nil
}

func CreateUploadTransaction(description string, buildId string, filesWithHashes []FileWithHash) (*upload_api.UploadTransactionRequest, error) {

	files := []upload_api.UploadFileRequest{}

	for _, file := range filesWithHashes {

		uploadFileRequest := upload_api.UploadFileRequest{
			FileName: file.FileWithoutPath,
			Hash:     file.Hash,
		}

		files = append(files, uploadFileRequest)
	}

	uploadTransaction := &upload_api.UploadTransactionRequest{
		Description: description,
		BuildId:     buildId,
		Files:       files,
	}

	return uploadTransaction, nil
}

func InitiateUploadTransaction(uploadAPIProtocol string, uploadAPIHost string, email string, pat string, uploadTransactionRequest upload_api.UploadTransactionRequest) (*upload_api.UploadTransactionResponse, error) {

	body, err := json.Marshal(uploadTransactionRequest)
	if err != nil {
		return nil, err
	}

	serviceUrl := fmt.Sprintf("%s://%s:%s@%s", uploadAPIProtocol, email, pat, uploadAPIHost)
	path := "UploadAPI"

	response, err := retryablehttp.Post(serviceUrl+"/"+path, "application/json", body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Backend API call failed, status code = %v", response.StatusCode))
	}

	uploadTransactionResponse := &upload_api.UploadTransactionResponse{}

	if err = json.NewDecoder(response.Body).Decode(uploadTransactionResponse); err != nil {
		return nil, err
	}

	return uploadTransactionResponse, nil
}

func UploadMissingFiles(uploadTransactionResponse upload_api.UploadTransactionResponse, filesWithHashes []FileWithHash) error {

	for _, uploadFileResponse := range uploadTransactionResponse.Files {

		fileWithHash := (*FileWithHash)(nil)

		for i := 0; i < len(filesWithHashes); i++ {
			currentFileWithHash := filesWithHashes[i]
			if currentFileWithHash.FileWithoutPath == uploadFileResponse.FileName && currentFileWithHash.Hash == uploadFileResponse.Hash {
				fileWithHash = &currentFileWithHash
				break
			}
		}

		file, err := os.Open(fileWithHash.FileWithPath)
		if err != nil {
			return err
		}

		log.Printf("Uploading file %s...", fileWithHash.FileWithPath)

		retryClient := retryablehttp.NewClient()
		req, err := retryablehttp.NewRequest(http.MethodPut, uploadFileResponse.Url, file)
		if err != nil {
			log.Printf("Request creation failed with error %v", err)
			return err
		}

		response, err := retryClient.Do(req)
		defer response.Body.Close()

		if err != nil {
			log.Printf("Upload failed with error %v", err)
			return err
		}

		if response.StatusCode != http.StatusOK {
			responseBody, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Printf("Reading error response body failed: %v", err)
				return errors.New(fmt.Sprintf("Reading error response body failed: %v", err))
			}
			log.Printf("Upload failed with status code %v; body = %v", response.StatusCode, string(responseBody))
			return errors.New(fmt.Sprintf("Backend API call failed, status code = %v", response.StatusCode))
		}
	}

	return nil
}
