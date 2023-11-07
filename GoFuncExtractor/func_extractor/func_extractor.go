package func_extractor

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type FuncExtractor struct {
	ProjectPath string
	Config      *FuncExtractorConfig
}

type FuncExtractorConfig struct {
	SingleLine bool
}

func NewFuncExtractor(projectPath string, c *FuncExtractorConfig) *FuncExtractor {
	return &FuncExtractor{
		ProjectPath: projectPath,
		Config:      c,
	}
}

func (f *FuncExtractor) ExtractFunctions() {
	files, err := findAllFiles(f.ProjectPath, ".go")
	if err != nil {
		panic(err)
	}
	functions, err := extractFunctions(files)
	if err != nil {
		panic(err)
	}
	for _, function := range functions {
		code, err := getFunctionCode(function, f.Config.SingleLine)
		if err != nil {
			panic(err)
		}
		println(code)
	}
}

func findAllFiles(dir string, ext string) ([]string, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, err
	}
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ext) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
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

func getFunctionCode(function *ast.FuncDecl, singleLine bool) (string, error) {
	fs := token.NewFileSet()
	var buf strings.Builder
	if err := printer.Fprint(&buf, fs, function); err != nil {
		return "", err
	}
	code := buf.String()
	if singleLine {
		code = strings.ReplaceAll(code, "\n", "\\n")
		return code, nil
	}
	return code, nil
}
