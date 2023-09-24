package main

import (
	"flag"
	"fmt"
	"go-extractor/extractor"
	"os"
)

var (
	hash *bool
	file *string
)

func init() {
	hash = flag.Bool("hash", false, "hash path (default: false)")
	file = flag.String("file", "", "file to parse (required)")
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
	ex, err := extractor.NewExtractor(*file, *hash)
	if err != nil {
		fmt.Println("Error creating extractor:", err)
		return
	}
	ex.GenerateProgramAstPaths()
}
