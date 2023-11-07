package main

import (
	"flag"
	. "go-func-extractor/func_extractor"
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
	funcExt.ExtractFunctions()
}
