package main

import (
	"fmt"
	"os"

	"github.com/TheSeaGiraffe/textgen/internal/options"
	"github.com/TheSeaGiraffe/textgen/internal/text"
)

func main() {
	// Parse command line flags
	appOpts := options.NewOpts()

	err := text.GenDummyFiles(appOpts.NumFiles, appOpts.NumParagraphs, appOpts.OutputDir)
	if err != nil {
		fmt.Printf("There was an error generating dummy files: %s", err)
		os.Exit(1)
	}
}
