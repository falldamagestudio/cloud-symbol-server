package cli

import (
	"bytes"
	"reflect"
	"testing"
)

var pageMap = [...]uint32{1, 0, 2}
var content = [...]byte{4, 5, 6, 7, 0, 1, 2, 3, 8, 9, 10, 11}

const pageSize = 4

func TestPagedStream_PageRemap(t *testing.T) {

	pagedStream := &PdbPagedStream{
		Reader:   bytes.NewReader(content[:]),
		PageMap:  pageMap[:],
		PageSize: pageSize,
	}

	// Read a section that covers three pages; part of the first, all of the second, and part of the third

	data := make([]byte, 9)
	if bytesRead, err := pagedStream.ReadAt(data, 1); bytesRead != len(data) || err != nil {
		if bytesRead != len(data) {
			t.Fatalf("Expected to read %v bytes, got %v", len(data), bytesRead)
		} else {
			t.Fatalf("pagedStream.ReadAt() returned %v", err)
		}
	}

	expectedData := [...]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(expectedData[:], data) {
		t.Fatalf("Expected to read %v, but got %v", expectedData, data)
	}
}
