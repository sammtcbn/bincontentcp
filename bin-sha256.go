package main

import (
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

	// Set the position for reading the specified range
	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return nil, err
	}

	// Calculate the length of the range
	if length == 0 {
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
	filename := "example.txt" // Replace with the filename of the file to calculate checksum
	offset := int64(0)        // Starting offset
	endOffset := int64(0)     // Ending offset, takes precedence if specified
	length := int64(1024)     // Length of the range, ignored if end offset is specified

	checksum, err := calculateChecksum(filename, offset, endOffset, length)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("SHA256 Checksum: %x\n", checksum)
}

