package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"readme_extraction/extractor"
	"readme_extraction/model"
)

var (
	inputPath  *string
	outputDir  *string
	singleLine *bool
	from       *int
	to         *int
)

const (
	outputFile = "output.json"
)

func init() {
	inputPath = flag.String("input-path", "", "Path to the input file")
	outputDir = flag.String("output-directory", "", "Path to the output directory")
	singleLine = flag.Bool("single-line", false, "Whether to output the readme in a single line (default: false)")
	from = flag.Int("from", 0, "From which repository to start extracting")
	to = flag.Int("to", -1, "To which repository to extract exclusive")
}

func main() {
	flag.Parse()
	if *inputPath == "" || *outputDir == "" {
		flag.Usage()
		return
	}
	var oldRepositories []model.Repository
	oldFile, err := os.ReadFile(*outputDir + fmt.Sprintf("/%s", outputFile))
	if err != nil {
	}
	err = json.Unmarshal(oldFile, &oldRepositories)
	if err != nil {
		fmt.Println("Error unmarshalling old file:", err)
		return
	}
	repositories, err := extractor.
		NewReadmeExtractor(inputPath, outputDir, from, to).
		Extract()
	if err != nil {
		fmt.Println("Error extracting readmes:", err)
	}
	// save the repositories to a file even if there are errors
	oldRepositories = append(oldRepositories, repositories...)
	buffer := new(bytes.Buffer)
	enc := json.NewEncoder(buffer)
	enc.SetEscapeHTML(false)
	if !*singleLine {
		enc.SetIndent("", "  ")
	}
	err = enc.Encode(oldRepositories)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	err = os.WriteFile(*outputDir+fmt.Sprintf("/%s", outputFile), buffer.Bytes(), 0644)
}
