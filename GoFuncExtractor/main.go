package main

import (
	"encoding/json"
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
	singleLine = flag.Bool("single-line", false, "output JSON on a single line")
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
	var jsonData []byte
	if *singleLine {
		jsonData, err = json.Marshal(functions)
	} else {
		jsonData, err = json.MarshalIndent(functions, "", "    ")
	}
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}
