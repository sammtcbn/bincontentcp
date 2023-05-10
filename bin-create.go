package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Define command-line arguments
	hexStr   := flag.String ("hex",     "",   "Hexadecimal value")
	fileSize := flag.Int64  ("size",    0,    "File size (bytes)")
	filePath := flag.String ("file",    "",   "File path")
    blockSize:= flag.Int64  ("bs",      4096, "Block size (bytes)")
	randHex  := flag.Bool   ("randhex", false, "Fill the file with random hex patterns.")
	flag.Parse()

    // Validate command-line flags
    if *filePath == "" {
        fmt.Println("Error: please specify file path")
        os.Exit(1)
    }

	if *fileSize == 0 {
		fmt.Println("Error: file size must be specified.")
		os.Exit(1)
	}

	// Check if both hex and randhex options are used
	if (*randHex) {
		if (len(*hexStr) != 0) {
		    fmt.Println("Error: hex and randhex options cannot be used together.")
			os.Exit(1)
	    }
	} else {
		if (len(*hexStr) == 0) {
			fmt.Println("Please specify hex")
			os.Exit(1)
		}
	}

	if (*randHex) {
		// Set random seed
		rand.Seed(time.Now().UnixNano())
	}

	// Open file
	file, err := os.Create(*filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

/* one byte write
	for i := int64(0); i < *fileSize; i++ {
		file.Write(hexByte)
	}
*/

// block write. This is much faster than writing one byte at a time.

	var written int64
	for written < *fileSize {
		toWrite := *fileSize - written
		if toWrite > *blockSize {
			toWrite = *blockSize
		}

		// Create a byte slice to hold the block data
		blockData := make([]byte, *blockSize)

		// Fill the byte slice with the specified hex pattern or random values
		if (*randHex) {
			for i := range blockData {
				blockData[i] = byte(rand.Intn(256))
			}
		} else {
			// Convert hexadecimal value to byte
			hexBytes, err := hex.DecodeString(*hexStr)
			if err != nil {
				panic(err)
			}
			for i := range blockData {
				blockData[i] = hexBytes[0]
			}
		}

		n, err := file.Write(blockData[:toWrite])
		if err != nil {
			panic(err)
		}
		written += int64(n)
	}

	fmt.Printf("Wrote %d bytes to file %s\n", written, *filePath)
}

