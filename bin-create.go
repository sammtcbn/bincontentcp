package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command-line arguments
	hexStr := flag.String("hex", "00", "Hexadecimal value")
	fileSize := flag.Int64("size", 1024, "File size (bytes)")
	filePath := flag.String("file", "", "File path")
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
	for i := int64(0); i < *fileSize; i++ {
		file.Write(hexByte)
	}
}

