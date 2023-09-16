package main

import (
	"fmt"
	"go-extractor/extractor"
)

func main() {
	filename := "extractor/test_assets/test.go"
	ex, err := extractor.NewExtractor(filename, false)
	if err != nil {
		fmt.Println("Error creating extractor:", err)
		return
	}
	ex.GenerateProgramAstPaths()
}
