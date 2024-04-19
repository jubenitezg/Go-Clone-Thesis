package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"metadata-extraction/extractor"
	"metadata-extraction/model"
	"os"
	"strings"
)

var (
	inputPath  *string
	outputFile *string
	singleLine *bool
	from       *int
	to         *int
)

var errors = color.New(color.FgRed).Add(color.Bold)
var info = color.New(color.FgYellow)

func init() {
	inputPath = flag.String("input-path", "", "Path to the input file")
	outputFile = flag.String("output-file", "output.json", "Path to the output directory")
	singleLine = flag.Bool("single-line", false, "Whether to output the readme in a single line (default: false)")
	from = flag.Int("from", 0, "From which repository to start extracting")
	to = flag.Int("to", -1, "To which repository to extract exclusive")
}

func save(prevMetadata []model.Metadata) {
	buffer := new(bytes.Buffer)
	enc := json.NewEncoder(buffer)
	enc.SetEscapeHTML(false)
	if !*singleLine {
		enc.SetIndent("", "  ")
	}
	err := enc.Encode(prevMetadata)
	if err != nil {
		errors.Println("Error marshalling JSON:", err)
		return
	}
	err = os.WriteFile(fmt.Sprintf("output/%s", *outputFile), buffer.Bytes(), 0644)
	if err != nil {
		errors.Println("Error writing to file:", err)
	}
}

func main() {
	flag.Parse()
	if *inputPath == "" || *outputFile == "" {
		flag.Usage()
		return
	}
	var prevMetadata []model.Metadata
	prevFile, err := os.ReadFile(fmt.Sprintf("output/%s", outputFile))
	if err == nil {
		err = json.Unmarshal(prevFile, &prevMetadata)
		if err != nil {
			errors.Println("Error unmarshalling old file:", err)
		}
	} else {
		errors.Println("Error reading old file:", err)
	}
	var repositories []model.Repository
	if _, err := os.Stat(*inputPath); os.IsNotExist(err) {
		errors.Printf("file %s does not exist\n", *inputPath)
	}
	file, err := os.ReadFile(*inputPath)
	if err != nil {
		errors.Println(err)
	}
	err = json.Unmarshal(file, &repositories)
	if err != nil {
		errors.Println(err)
	}

	if to == nil || *to == -1 {
		*to = len(repositories)
	}

	info.Println("Range: " + fmt.Sprint(*from) + " to " + fmt.Sprint(*to))
	for i, repository := range repositories[*from:*to] {
		info.Println("Repository #" + fmt.Sprint(i) + ": " + repository.Owner + "/" + repository.Name)
		metadata, err := extractor.GetRepositoryMetadata(&repository)
		if err != nil {
			errors.Println(err)
			if strings.Contains(err.Error(), "rate limit exceeded") {
				errors.Println("Rate limit exceeded, current repository processed: " + fmt.Sprint(i))
				os.Exit(1)
			}
		}
		if metadata != nil {
			info.Println("Saving repository" + fmt.Sprint(i))
			prevMetadata = append(prevMetadata, *metadata)
			save(prevMetadata)
		} else {
			info.Println("Skipping repository" + fmt.Sprint(i))
		}
	}
}
