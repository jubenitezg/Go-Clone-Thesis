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
	parsedAst        *ast.File
	fSet             *token.FileSet
	Functions        []*ast.FuncDecl
	FunctionFeatures map[string][]string
	hash             bool
}

func NewExtractor(file string, hash bool) (*Extractor, error) {
	fs := token.NewFileSet()
	parsedAst, err := parser.ParseFile(fs, file, nil, 0)
	if err != nil {
		return nil, err
	}
	functions := extractFunctions(parsedAst)
	ex := &Extractor{
		parsedAst:        parsedAst,
		fSet:             fs,
		Functions:        functions,
		FunctionFeatures: map[string][]string{},
		hash:             hash,
	}
	return ex, nil
}

func (e *Extractor) GenerateProgramAstPaths() []string {
	var programPaths []string
	_, err := e.generatePathForFunctions()
	if err != nil {
		return nil
	}
	for k, v := range e.FunctionFeatures {
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("%s ", k))
		for _, path := range v {
			sb.WriteString(fmt.Sprintf("%s ", path))
		}
		fmt.Println(sb.String())
		programPaths = append(programPaths, sb.String())
	}

	return programPaths
}

func (e *Extractor) generatePathForFunctions() ([]string, error) {
	var totalPaths []string
	for _, funcDecl := range e.Functions {
		leaves := extractLeavesFromFunc(funcDecl)
		funcName := funcDecl.Name.Name
		for i := 0; i < len(leaves)-1; i++ {
			for j := i + 1; j < len(leaves); j++ {
				nodeRelation, err := generatePathRelation(&leaves[i], &leaves[j])
				if err != nil {
					fmt.Println("Error generating path relation:", err)
					return nil, err
				}
				if e.hash {
					e.FunctionFeatures[funcName] = append(e.FunctionFeatures[funcName], nodeRelation.StringWithHash())
					totalPaths = append(totalPaths, nodeRelation.StringWithHash())
				} else {
					e.FunctionFeatures[funcName] = append(e.FunctionFeatures[funcName], nodeRelation.String())
					totalPaths = append(totalPaths, nodeRelation.String())
				}
			}
		}
	}

	return totalPaths, nil
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
			path := make([]common.AstNode, len(stack))
			for i, n := range stack {
				path[i] = common.AstNode{
					Node: n,
					Leaf: false,
				}
			}
			leaf := common.AstNode{
				Node: node,
				Leaf: true,
			}
			path = append(path, leaf)
			leaf.Path = path
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

func mergePaths(source *common.AstNode, target *common.AstNode) ([]common.AstNode, common.AstNode, []common.AstNode) {
	s, n, m := 0, len(source.Path), len(target.Path)
	for s < min(n, m) && source.Path[s].Node == target.Path[s].Node {
		s++
	}
	prefix := make([]common.AstNode, 0, len(source.Path)-s)
	for i := len(source.Path) - 1; i >= s; i-- {
		prefix = append(prefix, source.Path[i])
	}
	var lca common.AstNode
	if s > 0 {
		lca = source.Path[s-1]
	}
	suffix := target.Path[s:]

	return prefix, lca, suffix
}

func generatePathRelation(source *common.AstNode, target *common.AstNode) (*common.NodeRelation, error) {
	var pathSb strings.Builder
	prefix, lca, suffix := mergePaths(source, target)
	for _, node := range prefix {
		pathSb.WriteString(fmt.Sprintf("%s%s%s%s", constant.Start, node.Type(), constant.End, constant.Up))
	}
	pathSb.WriteString(fmt.Sprintf("%s%s%s", constant.Start, lca.Type(), constant.End))
	for _, node := range suffix {
		pathSb.WriteString(fmt.Sprintf("%s%s%s%s", constant.Down, constant.Start, node.Type(), constant.End))
	}

	return common.NewNodeRelation(source, target, pathSb.String())
}
