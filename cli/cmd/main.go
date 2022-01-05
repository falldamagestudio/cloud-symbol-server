package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/falldamagestudio/cloud-symbol-store/cli"
)

type TransactionRequest struct {
	Description string              `json:"description"`
	Files       []UploadFileRequest `json:"files"`
}

type UploadFileRequest struct {
	FileName string `json:"filename"`
	Hash     string `json:"hash"`
}

func createTransaction(description string, fileNames []string) (*TransactionRequest, error) {

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

	transaction := &TransactionRequest{
		Description: description,
		Files:       files,
	}

	return transaction, nil
}

func uploadTransaction(transactionRequest TransactionRequest) error {

	log.Printf("fake transaction: %v", transactionRequest)

	return nil
}

func upload(description string, fileNames []string) error {
	transaction, err := createTransaction(description, fileNames)
	if err != nil {
		return err
	}

	return uploadTransaction(*transaction)
}

func main() {

	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.Parse()
	if verbose {
		fmt.Println("verbose is on")
	}

	operation := flag.Arg(0)

	if operation == "upload" {
		_ = upload(flag.Arg(1), []string{flag.Arg(2)})
	}
}
