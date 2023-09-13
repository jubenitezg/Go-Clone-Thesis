package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/ast/inspector"
)

func parseGoFile(filename string) (*ast.File, *token.FileSet, error) {
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filename, nil, 0)
	if err != nil {
		return nil, nil, err
	}
	return node, fs, nil
}

func extractFunctions(parsedAST *ast.File, fset *token.FileSet) []*ast.FuncDecl {
	var funcDecls []*ast.FuncDecl

	ast.Inspect(parsedAST, func(node ast.Node) bool {
		if funcDecl, ok := node.(*ast.FuncDecl); ok {
			funcDecls = append(funcDecls, funcDecl)
		}
		return true // Continue the traversal
	})

	for _, funcDecl := range funcDecls {
		fmt.Println("Function: ", funcDecl.Name)
		fmt.Println("Start: ", fset.Position(funcDecl.Pos()))
		fmt.Println("End: ", fset.Position(funcDecl.End()))
	}

	return funcDecls
}

func discoverNodeParentsWithStack(parsedAST []*ast.File, fset *token.FileSet) {
	insp := inspector.New(parsedAST)
	insp.WithStack(nil, func(n ast.Node, push bool, stack []ast.Node) bool {
		if bexpr, ok := n.(*ast.BinaryExpr); push && ok && bexpr.Op == token.ADD {
			for i := len(stack) - 2; i >= 0; i-- {
				if _, ok := stack[i].(*ast.BinaryExpr); ok {
					fmt.Printf("found BinaryExpr(+) as a child of another binary expr: %v\n",
						fset.Position(n.Pos()))
					break
				}
			}
		}
		return true
	})
}

func extractLeavesFromFunc(funcDecl *ast.FuncDecl, fset *token.FileSet) []ast.Node {
	var leafNodes []ast.Node

	ast.Inspect(funcDecl, func(node ast.Node) bool {
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
	return leafNodes
}

//func generatePath()

func main() {
	filename := "test.go"
	parsedAST, fset, err := parseGoFile(filename)
	if err != nil {
		fmt.Println("Error parsing Go file:", err)
		return
	}

	//fmt.Println(parsedAST)
	discoverNodeParentsWithStack([]*ast.File{parsedAST}, fset)
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
