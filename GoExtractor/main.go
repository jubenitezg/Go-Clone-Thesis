package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func parseGoFile(filename string) (*ast.File, *token.FileSet, error) {
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filename, nil, 0)
	if err != nil {
		return nil, nil, err
	}
	return node, fs, nil
}

func extractLeaves(parsedAST *ast.File, fset *token.FileSet) {
	var leafNodes []ast.Node

	ast.Inspect(parsedAST, func(node ast.Node) bool {
		switch node.(type) {
		case *ast.Ident, *ast.BasicLit, *ast.CompositeLit:
			leafNodes = append(leafNodes, node)
		}
		return true // Continue the traversal
	})

	for _, leafNode := range leafNodes {
		fmt.Printf("Leaf Node Type: %T\n", leafNode)
		fmt.Println("Leaf Node: ", leafNode.Pos(), leafNode.End())
		fmt.Println(fset.Position(leafNode.Pos()).Line)
	}
}

func main() {
	filename := "test.go"
	parsedAST, fset, err := parseGoFile(filename)
	if err != nil {
		fmt.Println("Error parsing Go file:", err)
		return
	}

	fmt.Println(parsedAST)
	extractLeaves(parsedAST, fset)
}
