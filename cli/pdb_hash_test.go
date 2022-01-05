package cli

import (
	"testing"
)

func TestGetPDBHash_NoSuchFile(t *testing.T) {

	const path = "example.pdb.doesnotexist"

	_, err := GetPdbHash(path)

	if err == nil {
		t.Errorf("GetHash() with an invalid path was expected to fail, but succeeded")
	}
}

func TestGetPDBHash(t *testing.T) {

	const path = "example.pdb"

	hash, err := GetPdbHash(path)

	if err != nil {
		t.Fatalf("GetPDBHash(%v) failed with err = %v", path, err)
	}

	const desiredHash = "7F416863ABF34C3E894BAD1739BAA5571"

	if desiredHash != *hash {
		t.Errorf("GetPDBHash(%v) was expected to return has %v, but returned %v", path, desiredHash, hash)
	}
}
