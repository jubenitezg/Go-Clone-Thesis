package main

import (
	"flag"
	"fmt"
	"go-extractor/extractor"
	"os"
	"strings"
)

var (
	hash       *bool
	cleanPrint *bool
	file       *string
	maxLength  *int
	maxWidth   *int
)

func init() {
	hash = flag.Bool("hash", false, "hash path (default: false)")
	cleanPrint = flag.Bool("clean-print", false, "formatted print path (default: false)")
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
	ex, err := extractor.NewAstPathExtractor(*file, *hash, *maxLength, *maxWidth)
	if err != nil {
		fmt.Println("Error creating extractor:", err)
		return
	}
	paths := ex.GenerateProgramAstPaths()
	for _, path := range paths {
		if *cleanPrint {
			for _, s := range strings.Split(path, " ") {
				fmt.Println(s)
			}
		} else {
			fmt.Println(path)
		}
	}
}
