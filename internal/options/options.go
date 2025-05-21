package options

import "flag"

type AppOptions struct {
	NumFiles      int
	NumParagraphs int
	OutputDir     string
}

func NewOpts() AppOptions {
	var appOpts AppOptions

	flag.IntVar(&appOpts.NumFiles, "num-files", 1, "number of files to create")
	flag.IntVar(&appOpts.NumParagraphs, "num-paragraphs", 3, "number of paragraphs in each file")
	flag.StringVar(&appOpts.OutputDir, "out", ".", "directory into which files will be generated")

	flag.Parse()

	return appOpts
}
