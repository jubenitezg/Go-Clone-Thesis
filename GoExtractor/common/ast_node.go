package common

import "go/ast"

type AstNode struct {
	Node ast.Node
	Path []ast.Node
	Leaf bool
}
