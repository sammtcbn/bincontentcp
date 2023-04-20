package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	// Define command-line flags
	filename := flag.String("file", "", "file name")
	start := flag.Int64("start", 0, "start position")
	end := flag.Int64("end", -1, "end position")
	char := flag.String("char", "00", "hex character to replace")
	flag.Parse()

	// Validate command-line flags
	if *filename == "" {
		fmt.Println("Please specify file name")
		return
	}
	if *start < 0 {
		fmt.Println("Start position must be greater than or equal to 0")
		return
	}
	if *end < *start && *end != -1 {
		fmt.Println("End position must be greater than or equal to start position")
		return
	}
	if len(*char) != 2 {
		fmt.Println("Hex character must be 2 characters")
		return
	}

	// Read the file
	content, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Println("Unable to read file:", err)
		return
	}

	// Calculate end position
	if *end == -1 {
		*end = int64(len(content))
	}

	// Convert hex character to byte
	charByte, err := hex.DecodeString(*char)
	if err != nil {
		fmt.Println("Invalid hex character:", err)
		return
	}

	// Replace all binary in the specified range with the specified hex character
	for i := *start; i < *end; i++ {
		content[i] = charByte[0]
	}

	// Write the file
	err = ioutil.WriteFile(*filename, content, os.ModePerm)
	if err != nil {
		fmt.Println("Unable to write file:", err)
		return
	}

	// Show the result
	fmt.Printf("All binary between %d and %d in file %s have been replaced with %s\n", *start, *end, *filename, strconv.Quote(*char))
}

