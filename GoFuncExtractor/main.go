package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

var (
	projectPath *string
)

func init() {
	projectPath = flag.String("project-path", "", "path to project to parse (required)")
}

func findAllFiles(dir string) []string {
	var files []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".go") {
			files = append(files, path)
		}
		//fmt.Println(info.Name())
		return nil
	})
	return files
}

func extractFunctions(filePaths []string) ([]*ast.FuncDecl, error) {
	fs := token.NewFileSet()
	functions := make([]*ast.FuncDecl, 0)
	for _, filePath := range filePaths {
		file, err := parser.ParseFile(fs, filePath, nil, 0)
		if err != nil {
			return nil, err
		}
		for _, decl := range file.Decls {
			if fn, ok := decl.(*ast.FuncDecl); ok {
				functions = append(functions, fn)
			}
		}
	}
	return functions, nil
}

func getFunctionCode(function *ast.FuncDecl) (string, error) {
	fs := token.NewFileSet()
	var buf strings.Builder
	if err := printer.Fprint(&buf, fs, function); err != nil {
		return "", err
	}
	codeWithLiteral := strings.ReplaceAll(buf.String(), "\n", "\\n")
	return codeWithLiteral, nil
}

func main() {
	flag.Parse()
	if *projectPath == "" {
		flag.Usage()
		return
	}
	files := findAllFiles(*projectPath)
	fmt.Println(files)
	functions, err := extractFunctions(files)
	if err != nil {
		fmt.Println("Error extracting functions:", err)
		return
	}
	for _, fn := range functions {
		fmt.Println(getFunctionCode(fn))
	}
}
