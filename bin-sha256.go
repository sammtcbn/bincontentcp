package main

import (
	"flag"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func calculateChecksum(filename string, offset int64, endOffset int64, length int64) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Get the file size if endOffset is not specified
	if endOffset == -1 {
		fileInfo, err := file.Stat()
		if err != nil {
			return nil, err
		}
		endOffset = fileInfo.Size()
	}

	// Set the position for reading the specified range
	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return nil, err
	}

	// Calculate the length of the range
	if length == -1 {
		length = endOffset - offset
	}

	// Create a SHA256 hash object
	hash := sha256.New()

	// Read the specified range of data and update the hash value
	buffer := make([]byte, 4096) // Use a 4KB buffer
	var totalRead int64
	for totalRead < length {
		readSize := length - totalRead
		if readSize > int64(len(buffer)) {
			readSize = int64(len(buffer))
		}

		n, err := file.Read(buffer[:readSize])
		if err != nil {
			return nil, err
		}

		totalRead += int64(n)
		hash.Write(buffer[:n])
	}

	// Get the calculated SHA256 checksum
	checksum := hash.Sum(nil)
	return checksum, nil
}

func main() {
	// Define command-line arguments
	filePath    := flag.String ("file",    "",   "File path")
	startOffset := flag.Int64  ("start",   0,    "Starting Offset")
	endOffset   := flag.Int64  ("end",     -1,   "Ending Offset")
	length      := flag.Int64  ("length",  -1,   "Length of the range, ignored if end offset is specified")
	outbinPath  := flag.String ("outbin",  "",   "Output to BinFile")
	flag.Parse()

    // Validate command-line flags
    if *filePath == "" {
        fmt.Println("Error: please specify file path")
        os.Exit(1)
    }

	checksum, err := calculateChecksum(*filePath, *startOffset, *endOffset, *length)
	if err != nil {
		log.Fatal(err)
	}

	if (*outbinPath == "") {
		fmt.Printf("%x\n", checksum)
	} else {
		// Open file
		file, err := os.Create(*outbinPath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file.Write(checksum)
	}
}

