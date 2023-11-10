package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	. "go-func-extractor/func_extractor"
	"os"
)

var (
	projectPath *string
)

func init() {
	projectPath = flag.String("project-path", "", "path to project to parse (required)")
}

func main() {
	flag.Parse()
	if *projectPath == "" {
		flag.Usage()
		return
	}
	funcExt := NewFuncExtractor(*projectPath)
	functions, err := funcExt.ExtractFunctions()
	if err != nil {
		fmt.Println("Error extracting functions:", err)
		os.Exit(1)
	}
	for _, function := range functions {
		var buffer bytes.Buffer
		encoder := json.NewEncoder(&buffer)
		encoder.SetEscapeHTML(false)
		if err = encoder.Encode(function); err != nil {
			fmt.Println("Error encoding function:", err)
		} else {
			fmt.Print(buffer.String())
		}
	}
}
