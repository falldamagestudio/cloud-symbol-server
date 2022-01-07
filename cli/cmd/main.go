package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	cli "github.com/falldamagestudio/cloud-symbol-store/cli"
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

func upload(uploadAPIProtocol string, uploadAPIHost string, email string, pat string, description string, buildId string, patterns []string) error {

	fileNames, err := matchFiles(patterns)
	if err != nil {
		return err
	}

	if len(fileNames) == 0 {
		log.Printf("No files matched by patterns; skipping upload")
		return nil
	}

	filesWithHashes, err := cli.GetFilesWithHashes(fileNames)
	if err != nil {
		return err
	}

	uploadTransactionRequest, err := cli.CreateUploadTransaction(description, buildId, filesWithHashes)
	if err != nil {
		return err
	}

	uploadTransactionResponse, err := cli.InitiateUploadTransaction(uploadAPIProtocol, uploadAPIHost, email, pat, *uploadTransactionRequest)
	if err != nil {
		return err
	}

	err = cli.UploadMissingFiles(*uploadTransactionResponse, filesWithHashes)
	if err != nil {
		return err
	}

	return nil
}

func mainInt() int {

	var verbose bool
	var description string
	var buildId string
	var email string
	var pat string
	var uploadAPIProtocol string
	var uploadAPIHost string
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.StringVar(&description, "description", "", "Textual description")
	flag.StringVar(&buildId, "buildId", "", "Build ID")
	flag.StringVar(&email, "email", "", "Authentication email")
	flag.StringVar(&pat, "pat", "", "Authentication Personal Access Token")
	flag.StringVar(&uploadAPIProtocol, "uploadAPIProtocol", "", "Upload API protocol")
	flag.StringVar(&uploadAPIHost, "uploadAPIHost", "", "Upload API host")

	flag.Parse()
	if verbose {
		fmt.Println("verbose is on")
	}

	if email == "" || pat == "" {
		log.Printf("You must supply email and pat")
		return 1
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
			err := upload(uploadAPIProtocol, uploadAPIHost, email, pat, description, buildId, patterns)
			if err != nil {
				log.Printf("Error: %v", err)
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
