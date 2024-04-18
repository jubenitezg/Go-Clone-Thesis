package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"metadata-extraction/extractor"
	"metadata-extraction/model"
	"os"
)

var (
	inputPath  *string
	outputDir  *string
	singleLine *bool
	from       *int
)

const (
	outputFile = "output.json"
)

func init() {
	inputPath = flag.String("input-path", "", "Path to the input file")
	outputDir = flag.String("output-directory", "", "Path to the output directory")
	singleLine = flag.Bool("single-line", false, "Whether to output the readme in a single line (default: false)")
	from = flag.Int("from", 0, "From which repository to start extracting")
}

func save(prevMetadata []*model.Metadata, metadata *model.Metadata) {
	prevMetadata = append(prevMetadata, metadata)
	buffer := new(bytes.Buffer)
	enc := json.NewEncoder(buffer)
	enc.SetEscapeHTML(false)
	if !*singleLine {
		enc.SetIndent("", "  ")
	}
	err := enc.Encode(prevMetadata)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	err = os.WriteFile(*outputDir+fmt.Sprintf("/%s", outputFile), buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func main() {
	flag.Parse()
	if *inputPath == "" || *outputDir == "" {
		flag.Usage()
		return
	}
	var prevMetadata []*model.Metadata
	prevFile, err := os.ReadFile(*outputDir + fmt.Sprintf("/%s", outputFile))
	if err == nil {
		err = json.Unmarshal(prevFile, &prevMetadata)
		if err != nil {
			fmt.Println("Error unmarshalling old file:", err)
		}
	} else {
		fmt.Println("Error reading old file:", err)
	}
	var repositories []*model.Repository
	if _, err := os.Stat(*inputPath); os.IsNotExist(err) {
		fmt.Printf("file %s does not exist\n", *inputPath)
	}
	file, err := os.ReadFile(*inputPath)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(file, &repositories)
	if err != nil {
		fmt.Println(err)
	}

	for i, repository := range repositories[*from:1] {
		fmt.Println("Repository #" + fmt.Sprint(i) + ": " + repository.Owner + "/" + repository.Name)
		metadata, err := extractor.GetRepositoryMetadata(repository)
		if err != nil {
			fmt.Println(err)
		}
		if metadata != nil {
			save(prevMetadata, metadata)
		} else {
			fmt.Println("Skipping repository" + fmt.Sprint(i))
		}
	}

}
