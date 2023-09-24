package main

import (
	"flag"
	"fmt"
	"go-extractor/extractor"
	"os"
)

var (
	hash      *bool
	file      *string
	maxLength *int
	maxWidth  *int
)

func init() {
	hash = flag.Bool("hash", false, "hash path (default: false)")
	file = flag.String("file", "", "file to parse (required)")
	maxLength = flag.Int("max-length", 200, "maximum length of path (default: 200)")
	maxWidth = flag.Int("max-width", 200, "maximum width of path (default: 200)")
	flag.Usage = func() {
		_, err := fmt.Fprintf(os.Stderr, "Usage of gextr\n")
		if err != nil {
			panic(err)
		}
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	ex, err := extractor.NewExtractor(*file, *hash, *maxLength, *maxWidth)
	if err != nil {
		fmt.Println("Error creating extractor:", err)
		return
	}
	ex.GenerateProgramAstPaths()
}
