package cli

import (
	"errors"
	"fmt"
	"io"
)

type pagedStreamReader interface {
	io.Reader
	io.ReaderAt
}

type PdbPagedStream struct {
	Reader   pagedStreamReader
	PageMap  []uint32
	PageSize uint32
	Offset   int64
}

func min(a uint64, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func roundDown(value uint64, granularity uint64) uint64 {
	return value / granularity * granularity
}

func (pagedStream *PdbPagedStream) readChunk(output []byte, chunkStart uint64, chunkEnd uint64) (n int, err error) {
	chunkPageIndexVirtual := chunkStart / uint64(pagedStream.PageSize)
	chunkPageIndexPhysical := pagedStream.PageMap[chunkPageIndexVirtual]
	chunkPageOffset := chunkStart % uint64(pagedStream.PageSize)
	return pagedStream.Reader.ReadAt(output, int64(chunkPageIndexPhysical)*int64(pagedStream.PageSize)+int64(chunkPageOffset))
}

func (pagedStream *PdbPagedStream) ReadAt(output []byte, offset int64) (n int, err error) {

	uoffset := uint64(offset)
	upageSize := uint64(pagedStream.PageSize)

	end := min(uoffset+uint64(len(output)), uint64(len(pagedStream.PageMap))*upageSize)

	chunkStart := uoffset

	bytesRead := 0

	for chunkStart != end {
		chunkEnd := min(roundDown(chunkStart+upageSize, upageSize), end)
		count, err := pagedStream.readChunk(output[chunkStart-uoffset:chunkEnd-uoffset], chunkStart, chunkEnd)
		if err != nil {
			return bytesRead, err
		}
		bytesRead += count
		chunkStart = chunkEnd
	}

	return int(end - uint64(offset)), nil
}

func (pagedStream *PdbPagedStream) Read(output []byte) (n int, err error) {
	n, err = pagedStream.ReadAt(output, pagedStream.Offset)
	if err != nil {
		return
	}

	pagedStream.Offset += int64(len(output))
	return
}

func (pagedStream *PdbPagedStream) Seek(offset int64, whence int) (ret int64, err error) {
	if whence == 0 {
		pagedStream.Offset = offset
	} else if whence == 1 {
		pagedStream.Offset += offset
	} else if whence == 2 {
		pagedStream.Offset = int64(len(pagedStream.PageMap))*int64(pagedStream.PageSize) - offset
	} else {
		return 0, errors.New(fmt.Sprintf("Unknown whence %v", whence))
	}

	return offset, nil
}
