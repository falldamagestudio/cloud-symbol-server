package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/falldamagestudio/cloud-symbol-store/cli"
)

type UploadTransactionRequest struct {
	Description string              `json:"description"`
	BuildId     string              `json:"buildid"`
	Files       []UploadFileRequest `json:"files"`
}

type UploadFileRequest struct {
	FileName string `json:"filename"`
	Hash     string `json:"hash"`
}

func createUploadTransaction(description string, buildId string, fileNames []string) (*UploadTransactionRequest, error) {

	files := []UploadFileRequest{}

	for _, fileName := range fileNames {

		hash, err := cli.GetPdbHash(fileName)

		if err != nil {
			log.Printf("Error while parsing PDB: %v\n", err)
			return nil, err
		}

		uploadFileRequest := UploadFileRequest{
			FileName: fileName,
			Hash:     *hash,
		}

		files = append(files, uploadFileRequest)
	}

	uploadTransaction := &UploadTransactionRequest{
		Description: description,
		BuildId:     buildId,
		Files:       files,
	}

	return uploadTransaction, nil
}

func uploadTransaction(uploadTransactionRequest UploadTransactionRequest) error {

	log.Printf("fake transaction: %v", uploadTransactionRequest)

	return nil
}

func matchFiles(patterns []string) ([]string, error) {

	files := []string{}

	for _, pattern := range patterns {

		matches, err := filepath.Glob(pattern)
		if err != nil {
			return nil, err
		}

		files = append(files, matches...)
	}

	return files, nil
}

func upload(description string, buildId string, patterns []string) error {

	fileNames, err := matchFiles(patterns)
	if err != nil {
		return err
	}

	if len(fileNames) == 0 {
		log.Printf("No files matched by patterns; skipping upload")
		return nil
	}

	transaction, err := createUploadTransaction(description, buildId, fileNames)
	if err != nil {
		return err
	}

	return uploadTransaction(*transaction)
}

func mainInt() int {

	var verbose bool
	var description string
	var buildId string
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.StringVar(&description, "description", "", "Textual description")
	flag.StringVar(&buildId, "buildId", "", "Build ID")

	flag.Parse()
	if verbose {
		fmt.Println("verbose is on")
	}

	operation := flag.Arg(0)

	if operation == "upload" {
		patterns := []string{}
		for i := 1; i < flag.NArg(); i++ {
			patterns = append(patterns, flag.Arg(i))
		}

		if len(patterns) == 0 {
			log.Printf("You must provide at least one pattern for upload")
			return 1
		} else {
			err := upload(description, buildId, patterns)
			if err != nil {
				return 1
			} else {
				return 0
			}
		}
	}

	return 0
}

func main() {
	result := mainInt()
	os.Exit(result)
}
