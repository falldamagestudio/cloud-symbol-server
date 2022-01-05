package main

import (
	"log"

	"github.com/falldamagestudio/cloud-symbol-store/cli"
)

func upload() error {

	hash, err := cli.GetPdbHash("example.pdb")

	if err != nil {
		log.Printf("Error while parsing PDB: %v\n", err)
		return err
	}

	log.Printf("Hash: %v\n", hash)

	return nil
}

func main() {

	// var verbose bool
	// flag.BoolVar(&verbose, "verbose", false, "verbose output")
	// flag.Parse()
	// if verbose {
	// 	fmt.Println("verbose is on")
	// }

	// operation := flag.Arg(0)

	// if operation == "upload" {
	// 	_ = upload(flag.Arg(1))
	// }

	upload()
}
