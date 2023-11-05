package common

import (
	"fmt"
	"hash/fnv"
)

// NodeRelation represents a relation between two ast nodes
// it contains the source node, the target node and the path between them, hashed and not hashed
type NodeRelation struct {
	Source     *AstNode
	Target     *AstNode
	Path       string
	HashedPath uint64
}

// NewNodeRelation creates a new NodeRelation
func NewNodeRelation(source, target *AstNode, path string) (*NodeRelation, error) {
	h := fnv.New64()
	_, err := h.Write([]byte(path))
	if err != nil {
		return nil, err
	}
	return &NodeRelation{
		Source:     source,
		Target:     target,
		Path:       path,
		HashedPath: h.Sum64(),
	}, nil
}

// String returns a string representation of the NodeRelation, not hashed
func (n *NodeRelation) String() string {
	return fmt.Sprintf("%s,%s,%s", n.Source.String(), n.Path, n.Target.String())
}

// StringWithHash returns a string representation of the NodeRelation, hashed
func (n *NodeRelation) StringWithHash() string {
	return fmt.Sprintf("%s,%d,%s", n.Source.String(), n.HashedPath, n.Target.String())
}
