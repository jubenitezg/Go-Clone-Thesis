package extractor

import (
	"fmt"
	"go-extractor/common"
	"go-extractor/constant"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"slices"
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

func generatePathCompare(funcDecl *ast.FuncDecl, source *ast.Node, target *ast.Node) string {
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
		pathBuilder.WriteString(fmt.Sprintf("%s%s%s%s", constant.Start, getType(current), constant.End, constant.Up))
	}

	common := sourceTreeStack[len(sourceTreeStack)-commonPrefix]
	pathBuilder.WriteString(fmt.Sprintf("%s%s%s", constant.Start, getType(common), constant.End))

	for i := len(targetTreeStack) - commonPrefix - 1; i >= 0; i-- {
		current := targetTreeStack[i]
		pathBuilder.WriteString(fmt.Sprintf("%s%s%s%s", constant.Down, constant.Start, getType(current), constant.End))
	}
	return pathBuilder.String()
}

func (e *Extractor) GeneratePathForFunctionsCompare() []string {
	var paths []string
	for _, funcDecl := range e.Functions {
		leaves := extractLeavesFromFunc(funcDecl)
		for i := 0; i < len(leaves)-1; i++ {
			for j := i + 1; j < len(leaves); j++ {
				path := generatePathCompare(funcDecl, &leaves[i].Node, &leaves[j].Node)
				li := ""
				switch leaves[i].Node.(type) {
				case *ast.Ident:
					li = leaves[i].Node.(*ast.Ident).Name
				case *ast.BasicLit:
					li = leaves[i].Node.(*ast.BasicLit).Value
				}
				lj := ""
				switch leaves[j].Node.(type) {
				case *ast.Ident:
					lj = leaves[j].Node.(*ast.Ident).Name
				case *ast.BasicLit:
					lj = leaves[j].Node.(*ast.BasicLit).Value
				}
				//fmt.Println(fmt.Sprintf("%s,%s,%s", li, path, lj))
				paths = append(paths, fmt.Sprintf("%s,%s,%s", li, path, lj))
			}
		}
	}

	return paths
}

func (e *Extractor) GeneratePathForFunctions() []string {
	var paths []string
	for _, funcDecl := range e.Functions {
		leaves := extractLeavesFromFunc(funcDecl)
		for i := 0; i < len(leaves)-1; i++ {
			for j := i + 1; j < len(leaves); j++ {
				path := generatePath(leaves[i], leaves[j])
				li := ""
				switch leaves[i].Node.(type) {
				case *ast.Ident:
					li = leaves[i].Node.(*ast.Ident).Name
				case *ast.BasicLit:
					li = leaves[i].Node.(*ast.BasicLit).Value
				}
				lj := ""
				switch leaves[j].Node.(type) {
				case *ast.Ident:
					lj = leaves[j].Node.(*ast.Ident).Name
				case *ast.BasicLit:
					lj = leaves[j].Node.(*ast.BasicLit).Value
				}
				//fmt.Println(fmt.Sprintf("%s,%s,%s", li, path, lj))
				paths = append(paths, fmt.Sprintf("%s,%s,%s", li, path, lj))
			}
		}
	}
	return paths
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
		case *ast.Ident, *ast.BasicLit:
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
		pathSb.WriteString(fmt.Sprintf("%s%s%s%s", constant.Start, getType(source.Path[j]), constant.End, constant.Up))
	}
	pathSb.WriteString(fmt.Sprintf("%s%s%s", constant.Start, getType(firstAncestor), constant.End))
	for j := ancestorIdx; j < len(target.Path); j++ {
		pathSb.WriteString(fmt.Sprintf("%s%s%s%s", constant.Down, constant.Start, getType(target.Path[j]), constant.End))
	}

	return pathSb.String()
}

func getType(v any) string {
	tp := ""
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		tp = t.Elem().Name()
	} else {
		tp = t.Name()
	}
	op := ""
	switch v.(type) {
	case *ast.BinaryExpr:
		op = v.(*ast.BinaryExpr).Op.String()
	case *ast.UnaryExpr:
		op = v.(*ast.UnaryExpr).Op.String()
	case *ast.AssignStmt:
		op = v.(*ast.AssignStmt).Tok.String()
	}
	if len(op) > 0 {
		tp += ":" + op
	}
	return tp
}
