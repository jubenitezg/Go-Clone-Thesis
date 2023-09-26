package common

import (
	"fmt"
	"go/ast"
	"reflect"
)

// AstNode represents a node in the ast
// it contains the node itself, the path to the node and if it is a leaf node
type AstNode struct {
	Node ast.Node
	Path []AstNode
	Leaf bool
}

// Type returns the type of the node
func (a *AstNode) Type() string {
	t := reflect.TypeOf(a.Node)
	var tp string
	if t.Kind() == reflect.Ptr {
		tp = t.Elem().Name()
	} else {
		tp = t.Name()
	}
	if op, ok := getOperator(a.Node); ok {
		tp = fmt.Sprintf("%s:%s", tp, op)
	}

	return tp
}

// String returns a string representation of the node only if it is a leaf node
// otherwise it returns an empty string
func (a *AstNode) String() string {
	value := ""
	if a.Leaf {
		switch a.Node.(type) {
		case *ast.Ident:
			value = a.Node.(*ast.Ident).Name
		case *ast.BasicLit:
			value = a.Node.(*ast.BasicLit).Value
		}
	}
	return value
}

func getOperator(node any) (string, bool) {
	switch n := node.(type) {
	case *ast.BinaryExpr:
		return n.Op.String(), true
	case *ast.UnaryExpr:
		return n.Op.String(), true
	case *ast.AssignStmt:
		return n.Tok.String(), true
	default:
		return "", false
	}
}
