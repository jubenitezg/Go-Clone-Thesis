package func_extractor

import (
	"github.com/google/uuid"
	"go-func-extractor/common"
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
}

func NewFuncExtractor(projectPath string) *FuncExtractor {
	return &FuncExtractor{
		ProjectPath: projectPath,
	}
}

func (f *FuncExtractor) ExtractFunctions() ([]common.CodeFragment, error) {
	files, err := findAllFiles(f.ProjectPath, ".go")
	if err != nil {
		return nil, err
	}
	functions, err := extractFunctions(files)
	if err != nil {
		return nil, err
	}
	codeFragments := make([]common.CodeFragment, 0)
	for _, function := range functions {
		code, err := getFunctionCode(function)
		if err != nil {
			return nil, err
		}
		codeFragment := common.CodeFragment{
			Code: code,
			Id:   uuid.New().String(),
		}
		codeFragments = append(codeFragments, codeFragment)
	}
	return codeFragments, nil
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

func getFunctionCode(function *ast.FuncDecl) (string, error) {
	fs := token.NewFileSet()
	var buf strings.Builder
	if err := printer.Fprint(&buf, fs, function); err != nil {
		return "", err
	}
	return buf.String(), nil
}
