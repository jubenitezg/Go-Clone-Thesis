package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"slices"
	"strings"
)

const Up = "↑"
const Down = "↓"

func parseGoFile(filename string) (*ast.File, *token.FileSet, error) {
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filename, nil, 0)
	if err != nil {
		return nil, nil, err
	}
	return node, fs, nil
}

func extractFunctions(parsedAST *ast.File) []*ast.FuncDecl {
	var funcDecls []*ast.FuncDecl
	ast.Inspect(parsedAST, func(node ast.Node) bool {
		if funcDecl, ok := node.(*ast.FuncDecl); ok {
			funcDecls = append(funcDecls, funcDecl)
		}
		return true
	})

	return funcDecls
}

// getTreeStack will return the stack of nodes from the root to the node provided
// it needs the function declaration to be able to get the stack
func getTreeStack(funcDecl *ast.FuncDecl, node *ast.Node) []ast.Node {
	var tmpStack []ast.Node
	var stack []ast.Node
	ast.Inspect(funcDecl, func(n ast.Node) bool {
		if n == *node {
			stack = make([]ast.Node, len(tmpStack))
			copy(stack, tmpStack)
			return false
		}
		if n == nil {
			tmpStack = tmpStack[:len(tmpStack)-1]
		} else {
			tmpStack = append(tmpStack, n)
		}
		return true
	})
	stack = append(stack, *node)
	slices.Reverse(stack)

	return stack
}

func extractLeavesFromFunc(funcDecl *ast.FuncDecl) []ast.Node {
	var leafNodes []ast.Node
	ast.Inspect(funcDecl, func(node ast.Node) bool {
		switch node.(type) {
		case *ast.Ident, *ast.BasicLit, *ast.CompositeLit:
			leafNodes = append(leafNodes, node)
		}
		return true
	})

	return leafNodes
}

func generatePath(funcDecl *ast.FuncDecl, source *ast.Node, target *ast.Node) string {
	sourceTreeStack := getTreeStack(funcDecl, source)
	targetTreeStack := getTreeStack(funcDecl, target)
	var pathBuilder strings.Builder
	commonPrefix := 0
	currentSourceAncestor := len(sourceTreeStack) - 1
	currentTargetAncestor := len(targetTreeStack) - 1
	for currentSourceAncestor >= 0 && currentTargetAncestor >= 0 &&
		sourceTreeStack[currentSourceAncestor] == targetTreeStack[currentTargetAncestor] {
		commonPrefix++
		currentSourceAncestor--
		currentTargetAncestor--
	}

	for i := 0; i < len(sourceTreeStack)-commonPrefix; i++ {
		current := sourceTreeStack[i]
		pathBuilder.WriteString(fmt.Sprintf("%T%s", current, Up))
	}

	common := sourceTreeStack[len(sourceTreeStack)-commonPrefix]
	pathBuilder.WriteString(fmt.Sprintf("%T", common))

	for i := len(targetTreeStack) - commonPrefix - 1; i >= 0; i-- {
		current := targetTreeStack[i]
		pathBuilder.WriteString(fmt.Sprintf("%s%T", Down, current))
	}
	return pathBuilder.String()
}

func main() {
	filename := "test.go"
	parsedAST, _, err := parseGoFile(filename)
	if err != nil {
		fmt.Println("Error parsing Go file:", err)
		return
	}

	//fmt.Println(parsedAST)
	functions := extractFunctions(parsedAST)
	leaves := extractLeavesFromFunc(functions[0])
	fmt.Println(generatePath(functions[0], &leaves[7], &leaves[10]))
	//for _, function := range extractFunctions(parsedAST, fset) {
	//	//fmt.Println("Function: ", function.Name.Name)
	//	leaves := extractLeavesFromFunc(function, fset)
	//	for i := 0; i < len(leaves); i++ {
	//		for j := i + 1; j < len(leaves); j++ {
	//			//fmt.Println(leaves[i].Pos(), leaves[j].Pos())
	//			//generatePath(leaves[i], leaves[j])
	//		}
	//	}
	//}
}
