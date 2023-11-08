package main

import (
	"flag"
	"fmt"
	. "go-func-extractor/func_extractor"
	"os"
)

var (
	projectPath *string
	singleLine  *bool
)

func init() {
	projectPath = flag.String("project-path", "", "path to project to parse (required)")
	singleLine = flag.Bool("single-line", false, "single line function (default: false)")
}

func main() {
	flag.Parse()
	if *projectPath == "" {
		flag.Usage()
		return
	}
	funcExt := NewFuncExtractor(*projectPath, &FuncExtractorConfig{SingleLine: *singleLine})
	err := funcExt.ExtractFunctions()
	if err != nil {
		fmt.Println("Error extracting functions:", err)
		os.Exit(1)
	}
}
