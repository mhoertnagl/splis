package main

import (
	"strconv"
	"strings"
)

type Node interface {
}

type NumNode interface {
	Value() int32
}

type numNode struct {
	val int32
}

func NewNumNode(s string) NumNode {
	return &numNode{parseNumber(s)}
}

func parseNumber(s string) int32 {
	var n int64
	if strings.HasPrefix(s, "0b") {
		n, _ = strconv.ParseInt(s[2:], 2, 32)
	} else if strings.HasPrefix(s, "0x") {
		n, _ = strconv.ParseInt(s[2:], 16, 32)
	} else {
		n, _ = strconv.ParseInt(s, 10, 32)
	}
	return int32(n)
}

func (n *numNode) Value() int32 {
	return n.val
}

type SymNode interface {
	Name() string
}

type symNode struct {
	name string
}

func NewSymNode(s string) SymNode {
	return &symNode{s}
}

func (n *symNode) Name() string {
	return n.name
}

type SeqNode interface {
	Push(c Node)
	Cell(i int) Node
	Len() int
}

type sExprNode struct {
	cells []Node
}

func NewSExprNode() SeqNode {
	return &sExprNode{[]Node{}}
}

func (n *sExprNode) Push(c Node) {
	n.cells = append(n.cells, c)
}

func (n *sExprNode) Cell(i int) Node {
	return n.cells[i]
}

func (n *sExprNode) Len() int {
	return len(n.cells)
}

type qExprNode struct {
	cells []Node
}

func NewQExprNode() SeqNode {
	return &qExprNode{[]Node{}}
}

func (n *qExprNode) Push(c Node) {
	n.cells = append(n.cells, c)
}

func (n *qExprNode) Cell(i int) Node {
	return n.cells[i]
}

func (n *qExprNode) Len() int {
	return len(n.cells)
}
