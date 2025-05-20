package main

import (
	"fmt"
	"os"

	"github.com/TheSeaGiraffe/textgen/internal/text"
)

func main() {
	// Write 5 files with 3 paragraphs each to "out" directory
	numFiles := 5
	numParagraphs := 3
	outputDir := "out"
	err := text.GenDummyFiles(numFiles, numParagraphs, outputDir)
	if err != nil {
		fmt.Printf("There was an error generating dummy files: %s", err)
		os.Exit(1)
	}
}
