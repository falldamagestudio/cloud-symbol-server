package cli

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const msf7Signature = "Microsoft C/C++ MSF 7.00\r\n\x1ADS\x00\x00\x00"
const msf7SignatureLength = 32

func isMSF7Format(file os.File) (bool, error) {

	signatureBuffer := make([]byte, msf7SignatureLength)
	if count, err := file.ReadAt(signatureBuffer, 0); count != msf7SignatureLength || err != nil {

		if err != nil {
			return false, err
		} else {
			return false, nil
		}
	}

	signature := string(signatureBuffer[:])
	if msf7Signature != signature {
		return false, nil
	}

	return true, nil
}

// type MSF7Header struct {
// 	// Page size in bytes.
// 	PageSize int32
// 	// Page number of free page map.
// 	FreePageMapPageNum uint16
// 	// Number of pages.
// 	NPages uint16
// 	// Stream information about the stream table.
// 	StreamTblInfo StreamInfo
// 	// Maps from stream page number to page number.
// 	PageNumMap []uint16 // length: math.Ceil(msfHdr.StreamTblInfo.Size / msfHdr.PageSize)
// 	// align until page boundry.
// }

type MSF7Superblock struct {
	FileMagic         [msf7SignatureLength]byte
	BlockSize         uint32
	FreeBlockMapBlock uint32
	NumBlocks         uint32
	NumDirectoryBytes uint32
	Unknown           uint32
	BlockMapAddr      uint32
}

func readMSF7Superblock(r io.Reader) (*MSF7Superblock, error) {

	msf7Superblock := &MSF7Superblock{}

	if err := binary.Read(r, binary.LittleEndian, &msf7Superblock.FileMagic); err != nil {
		return nil, fmt.Errorf("Error while reading file header: %w", err)
	}

	if err := binary.Read(r, binary.LittleEndian, &msf7Superblock.BlockSize); err != nil {
		return nil, fmt.Errorf("Error while reading file header: %w", err)
	}

	if err := binary.Read(r, binary.LittleEndian, &msf7Superblock.FreeBlockMapBlock); err != nil {
		return nil, fmt.Errorf("Error while reading file header: %w", err)
	}

	if err := binary.Read(r, binary.LittleEndian, &msf7Superblock.NumBlocks); err != nil {
		return nil, fmt.Errorf("Error while reading file header: %w", err)
	}

	if err := binary.Read(r, binary.LittleEndian, &msf7Superblock.NumDirectoryBytes); err != nil {
		return nil, fmt.Errorf("Error while reading file header: %w", err)
	}

	if err := binary.Read(r, binary.LittleEndian, &msf7Superblock.Unknown); err != nil {
		return nil, fmt.Errorf("Error while reading file header: %w", err)
	}

	if err := binary.Read(r, binary.LittleEndian, &msf7Superblock.BlockMapAddr); err != nil {
		return nil, fmt.Errorf("Error while reading file header: %w", err)
	}

	return msf7Superblock, nil
}

type MSF struct {
}

func readBlock(file os.File, blockIndex int, blockSize int) ([]byte, error) {

	block := make([]byte, blockSize)
	if bytesRead, err := file.ReadAt(block, int64(blockIndex)*int64(blockSize)); bytesRead != blockSize || err != nil {
		if err != nil {
			return nil, err
		} else {
			return nil, errors.New(fmt.Sprintf("Expected to read %v bytes, but only read %v", blockSize, bytesRead))
		}
	}

	return block, nil
}

func readPageMap(file os.File, blockMapAddr int, blockSize int, numDirectoryBytes int) ([]uint32, error) {
	block, err := readBlock(file, blockMapAddr, blockSize)
	if err != nil {
		return nil, err
	}

	numBlockMapEntries := (numDirectoryBytes + blockSize - 1) / blockSize

	blockMap := make([]uint32, numBlockMapEntries)

	if err = binary.Read(bytes.NewReader(block), binary.LittleEndian, blockMap); err != nil {
		return nil, err
	}

	return blockMap, nil
}

func readStreamDirectoryBlocks(file os.File, blockMap []uint32, blockSize int, numDirectoryBytes int) ([]uint32, error) {

	numStreamDirectoryBlocks := (numDirectoryBytes + blockSize - 1) / blockSize
	streamDirectoryBlocks := make([]uint32, numStreamDirectoryBlocks)

	streamDirectoryEntriesPerBlock := blockSize / 4

	for blockMapBlockId, blockMapBlockValue := range blockMap {
		block, err := readBlock(file, int(blockMapBlockValue), blockSize)
		if err != nil {
			return nil, err
		}

		if err = binary.Read(bytes.NewReader(block), binary.LittleEndian, streamDirectoryBlocks[blockMapBlockId*streamDirectoryEntriesPerBlock:numStreamDirectoryBlocks]); err != nil {
			return nil, err
		}

	}

	return streamDirectoryBlocks, nil
}

type StreamDirectoryHeader struct {
	NumStreams           uint32
	StreamSizes          []uint32
	StreamPageMapOffsets []uint32
	StreamPageMapCounts  []uint32
}

func readStreamPageMap(pagedStreamDirectory *PdbPagedStream, streamDirectoryHeader *StreamDirectoryHeader, streamId int) ([]uint32, error) {
	if streamId < 0 || streamId >= int(streamDirectoryHeader.NumStreams) {
		return nil, errors.New(fmt.Sprintf("StreamId out of bounds: id %v, but there are only %v streams", streamId, streamDirectoryHeader.NumStreams))
	}

	pageMap := make([]uint32, streamDirectoryHeader.StreamPageMapCounts[streamId])
	pagedStreamDirectory.Seek(int64(streamDirectoryHeader.StreamPageMapOffsets[streamId]), 0)
	if err := binary.Read(pagedStreamDirectory, binary.LittleEndian, &pageMap); err != nil {
		return nil, fmt.Errorf("Error while reading stream directory header: %w", err)
	}

	return pageMap, nil
}

func readStreamDirectoryHeader(pagedStreamDirectory *PdbPagedStream) (*StreamDirectoryHeader, error) {

	streamDirectoryHeader := &StreamDirectoryHeader{}

	if err := binary.Read(pagedStreamDirectory, binary.LittleEndian, &streamDirectoryHeader.NumStreams); err != nil {
		return nil, fmt.Errorf("Error while reading stream directory header: %w", err)
	}

	streamDirectoryHeader.StreamSizes = make([]uint32, streamDirectoryHeader.NumStreams)
	streamDirectoryHeader.StreamPageMapOffsets = make([]uint32, streamDirectoryHeader.NumStreams)
	streamDirectoryHeader.StreamPageMapCounts = make([]uint32, streamDirectoryHeader.NumStreams)

	if err := binary.Read(pagedStreamDirectory, binary.LittleEndian, &streamDirectoryHeader.StreamSizes); err != nil {
		return nil, fmt.Errorf("Error while reading stream directory header: %w", err)
	}

	offset := (1 + streamDirectoryHeader.NumStreams) * 4

	for i, streamSize := range streamDirectoryHeader.StreamSizes {

		streamDirectoryHeader.StreamPageMapOffsets[i] = offset
		numPages := (streamSize + pagedStreamDirectory.PageSize - 1) / pagedStreamDirectory.PageSize
		streamDirectoryHeader.StreamPageMapCounts[i] = numPages
		offset += numPages * 4
	}

	return streamDirectoryHeader, nil
}

type Guid struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

type PdbStreamHeader struct {
	Version   uint32
	Signature uint32
	Age       uint32
	UniqueId  Guid
}

const PdbStreamIndex = 1

func readPdbStreamHeader(pagedStreamDirectory *PdbPagedStream, streamDirectoryHeader *StreamDirectoryHeader) (*PdbStreamHeader, error) {

	pageMap, err := readStreamPageMap(pagedStreamDirectory, streamDirectoryHeader, PdbStreamIndex)
	if err != nil {
		return nil, err
	}

	pagedPdbStream := PdbPagedStream{
		Reader:   pagedStreamDirectory.Reader,
		PageMap:  pageMap,
		PageSize: pagedStreamDirectory.PageSize,
	}

	pdbStreamHeader := &PdbStreamHeader{}
	if err := binary.Read(&pagedPdbStream, binary.LittleEndian, &pdbStreamHeader.Version); err != nil {
		return nil, fmt.Errorf("Error while reading PDB stream header: %w", err)
	}
	if err := binary.Read(&pagedPdbStream, binary.LittleEndian, &pdbStreamHeader.Signature); err != nil {
		return nil, fmt.Errorf("Error while reading PDB stream header: %w", err)
	}
	if err := binary.Read(&pagedPdbStream, binary.LittleEndian, &pdbStreamHeader.Age); err != nil {
		return nil, fmt.Errorf("Error while reading PDB stream header: %w", err)
	}
	if err := binary.Read(&pagedPdbStream, binary.LittleEndian, &pdbStreamHeader.UniqueId.Data1); err != nil {
		return nil, fmt.Errorf("Error while reading PDB stream header: %w", err)
	}
	if err := binary.Read(&pagedPdbStream, binary.LittleEndian, &pdbStreamHeader.UniqueId.Data2); err != nil {
		return nil, fmt.Errorf("Error while reading PDB stream header: %w", err)
	}
	if err := binary.Read(&pagedPdbStream, binary.LittleEndian, &pdbStreamHeader.UniqueId.Data3); err != nil {
		return nil, fmt.Errorf("Error while reading PDB stream header: %w", err)
	}
	if err := binary.Read(&pagedPdbStream, binary.BigEndian, &pdbStreamHeader.UniqueId.Data4); err != nil {
		return nil, fmt.Errorf("Error while reading PDB stream header: %w", err)
	}

	return pdbStreamHeader, nil

}

func getMSF7Hash(file os.File) (*string, error) {

	msf7Superblock, err := readMSF7Superblock(&file)
	if err != nil {
		return nil, err
	}

	streamDirectoryPageMap, err := readPageMap(file, int(msf7Superblock.BlockMapAddr), int(msf7Superblock.BlockSize), int(msf7Superblock.NumDirectoryBytes))
	if err != nil {
		return nil, err
	}

	pagedStreamDirectory := PdbPagedStream{
		Reader:   &file,
		PageMap:  streamDirectoryPageMap,
		PageSize: msf7Superblock.BlockSize,
	}

	streamDirectoryHeader, err := readStreamDirectoryHeader(&pagedStreamDirectory)

	pdbStreamHeader, err := readPdbStreamHeader(&pagedStreamDirectory, streamDirectoryHeader)
	if err != nil {
		return nil, err
	}

	hash := fmt.Sprintf("%08X%04X%04X%s%d", pdbStreamHeader.UniqueId.Data1, pdbStreamHeader.UniqueId.Data2, pdbStreamHeader.UniqueId.Data3, strings.ToUpper(hex.EncodeToString(pdbStreamHeader.UniqueId.Data4[:])), pdbStreamHeader.Age)

	return &hash, nil
}

func GetPdbHash(pdbPath string) (*string, error) {

	file, err := os.Open(pdbPath)

	if err != nil {
		return nil, err
	}

	isMSF7, err := isMSF7Format(*file)

	if err != nil {
		return nil, fmt.Errorf("Unable to parse file: %w", err)
	}

	if isMSF7 {
		return getMSF7Hash(*file)
	} else {
		return nil, errors.New("Unknown file format")
	}

}
