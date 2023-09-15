package extractor

import (
	"fmt"
	"go-extractor/common"
	"go-extractor/constant"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type Extractor struct {
	parsedAst *ast.File
	fSet      *token.FileSet
	Functions []*ast.FuncDecl
}

func New(file string) (*Extractor, error) {
	fs := token.NewFileSet()
	parsedAst, err := parser.ParseFile(fs, file, nil, 0)
	if err != nil {
		return nil, err
	}
	functions := extractFunctions(parsedAst)
	ex := &Extractor{
		parsedAst: parsedAst,
		fSet:      fs,
		Functions: functions,
	}
	return ex, nil
}

func (e *Extractor) GeneratePathForFunctions() {
	for _, funcDecl := range e.Functions {
		leaves := extractLeavesFromFunc(funcDecl)
		for i := 0; i < len(leaves)-1; i++ {
			for j := i + 1; j < len(leaves); j++ {
				path := generatePath(leaves[i], leaves[j])
				fmt.Println(path)
			}
		}
	}
}

func extractFunctions(parsedAst *ast.File) []*ast.FuncDecl {
	var funcDecls []*ast.FuncDecl
	ast.Inspect(parsedAst, func(node ast.Node) bool {
		if funcDecl, ok := node.(*ast.FuncDecl); ok {
			funcDecls = append(funcDecls, funcDecl)
		}
		return true
	})

	return funcDecls
}

func extractLeavesFromFunc(funcDecl *ast.FuncDecl) []common.AstNode {
	var leafNodes []common.AstNode
	var stack []ast.Node
	ast.Inspect(funcDecl, func(node ast.Node) bool {
		switch node.(type) {
		case *ast.Ident, *ast.BasicLit, *ast.CompositeLit:
			path := make([]ast.Node, len(stack))
			copy(path, stack)
			path = append(path, node)
			leaf := common.AstNode{
				Node: node,
				Path: path,
				Leaf: true,
			}
			leafNodes = append(leafNodes, leaf)
		}
		if node == nil {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, node)
		}
		return true
	})

	return leafNodes
}

func generatePath(source common.AstNode, target common.AstNode) string {
	var pathSb strings.Builder
	ancestorIdx := 0
	maxAncestorIdx := min(len(source.Path), len(target.Path))
	for ancestorIdx < maxAncestorIdx && source.Path[ancestorIdx] == target.Path[ancestorIdx] {
		ancestorIdx++
	}
	firstAncestor := source.Path[ancestorIdx-1]
	for j := len(source.Path) - 1; j >= ancestorIdx; j-- {
		pathSb.WriteString(fmt.Sprintf("%T%s", source.Path[j], constant.Up))
	}
	pathSb.WriteString(fmt.Sprintf("%T", firstAncestor))
	for j := ancestorIdx; j < len(target.Path); j++ {
		pathSb.WriteString(fmt.Sprintf("%s%T", constant.Down, target.Path[j]))
	}

	return pathSb.String()
}
