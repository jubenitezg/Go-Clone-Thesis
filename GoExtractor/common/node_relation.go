package common

import (
	"fmt"
	"hash/fnv"
)

type NodeRelation struct {
	Source     *AstNode
	Target     *AstNode
	Path       string
	HashedPath uint64
}

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

func (n *NodeRelation) String() string {
	return fmt.Sprintf("%s,%s,%s", n.Source.String(), n.Path, n.Target.String())
}

func (n *NodeRelation) StringWithHash() string {
	return fmt.Sprintf("%s,%d,%s", n.Source.String(), n.HashedPath, n.Target.String())
}
