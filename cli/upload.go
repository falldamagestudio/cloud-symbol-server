package cli

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	retryablehttp "github.com/hashicorp/go-retryablehttp"

	openapi "github.com/falldamagestudio/cloud-symbol-store/cli/generated"
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

func CreateUploadTransaction(description string, buildId string, filesWithHashes []FileWithHash) (*openapi.UploadTransactionRequest, error) {

	files := []openapi.UploadFileRequest{}

	for _, file := range filesWithHashes {

		uploadFileRequest := openapi.UploadFileRequest{
			FileName: &file.FileWithoutPath,
			Hash:     &file.Hash,
		}

		files = append(files, uploadFileRequest)
	}

	uploadTransaction := &openapi.UploadTransactionRequest{
		Description: &description,
		BuildId:     &buildId,
		Files:       &files,
	}

	return uploadTransaction, nil
}

func InitiateUploadTransaction(uploadAPIProtocol string, uploadAPIHost string, email string, pat string, uploadTransactionRequest openapi.UploadTransactionRequest) (*openapi.UploadTransactionResponse, error) {

	serviceUrl := fmt.Sprintf("%s://%s:%s@%s", uploadAPIProtocol, email, pat, uploadAPIHost)

	configuration := openapi.NewConfiguration()
	configuration.Servers = []openapi.ServerConfiguration{
		{
			URL:         serviceUrl,
			Description: "Custom URL",
			Variables:   nil,
		},
	}
	api_client := openapi.NewAPIClient(configuration)
	uploadTransactionResponse, httpResponse, err := api_client.DefaultApi.CreateTransaction(context.Background()).UploadTransactionRequest(uploadTransactionRequest).Execute()

	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Backend API call failed, status code = %v", httpResponse.StatusCode))
	}

	return &uploadTransactionResponse, nil
}

func UploadMissingFiles(uploadTransactionResponse openapi.UploadTransactionResponse, filesWithHashes []FileWithHash) error {

	for _, uploadFileResponse := range *uploadTransactionResponse.Files {

		fileWithHash := (*FileWithHash)(nil)

		for i := 0; i < len(filesWithHashes); i++ {
			currentFileWithHash := filesWithHashes[i]
			if currentFileWithHash.FileWithoutPath == *uploadFileResponse.FileName && currentFileWithHash.Hash == *uploadFileResponse.Hash {
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
		req, err := retryablehttp.NewRequest(http.MethodPut, *uploadFileResponse.Url, file)
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
