package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command-line arguments
	hexStr   := flag.String ("hex",  "00", "Hexadecimal value")
	fileSize := flag.Int64  ("size", 1024, "File size (bytes)")
	filePath := flag.String ("file", "",   "File path")
    blockSize:= flag.Int64  ("bs",   4096, "Block size (bytes)")
	flag.Parse()

    // Validate command-line flags
    if *filePath == "" {
        fmt.Println("Please specify file path")
        return
    }

	// Convert hexadecimal value to byte
	hexByte, err := hex.DecodeString(*hexStr)
	if err != nil {
		panic(err)
	}

	// Open file
	file, err := os.Create(*filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Fill file with hexadecimal value

/* one byte write
	for i := int64(0); i < *fileSize; i++ {
		file.Write(hexByte)
	}
*/

// block write. This is much faster than writing one byte at a time.
	block := make([]byte, *blockSize)
	for i := range block {
		block[i] = hexByte[0]
	}

	var written int64
	for written < *fileSize {
		toWrite := *fileSize - written
		if toWrite > *blockSize {
			toWrite = *blockSize
		}
		n, err := file.Write(block[:toWrite])
		if err != nil {
			panic(err)
		}
		written += int64(n)
	}

	fmt.Printf("Wrote %d bytes to file %s\n", written, *filePath)
}

