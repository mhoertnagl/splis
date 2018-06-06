package main

import (
	"strconv"
	"strings"
)

type NodeType int

const (
	ERR_NODE NodeType = iota
	NUM_NODE
	SYM_NODE
	SXP_NODE
	QXP_NODE
	FUN_NODE
	LBD_NODE
)

func (t NodeType) Name() string {
	switch t {
	case ERR_NODE:
		return "Error"
	case NUM_NODE:
		return "Number"
	case SYM_NODE:
		return "Symbol"
	case SXP_NODE:
		return "S-Expression"
	case QXP_NODE:
		return "Q-Expression"
	case FUN_NODE:
		return "Function"
	case LBD_NODE:
		return "Lambda"
	default:
		return "unknown"
	}
}

type Node interface {
	Type() NodeType
}

type errNode struct {
	msg string
}

func NewErrNode(msg string) *errNode {
	return &errNode{msg}
}

func (n *errNode) Type() NodeType {
	return ERR_NODE
}

type numNode struct {
	val int32
}

func NewNumNode(s string) *numNode {
	return &numNode{parseNumber(s)}
}

func (n *numNode) Type() NodeType {
	return NUM_NODE
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

type symNode struct {
	name string
}

func NewSymNode(s string) *symNode {
	return &symNode{s}
}

func (n *symNode) Type() NodeType {
	return SYM_NODE
}

type SeqNode interface {
	Type() NodeType
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

func (n *sExprNode) Type() NodeType {
	return SXP_NODE
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

func (n *qExprNode) Type() NodeType {
	return QXP_NODE
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

type Fun func(Env, []Node) Node

type funNode struct {
	fun Fun
}

func NewFunNode(f Fun) *funNode {
	return &funNode{f}
}

func (n *funNode) Type() NodeType {
	return FUN_NODE
}

type lambdaNode struct {
	env  Env
	ps   []*symNode
	body SeqNode
}

func NewLambdaNode(e Env, ps []*symNode, body SeqNode) *lambdaNode {
	return &lambdaNode{e, ps, body}
}

func (n *lambdaNode) Type() NodeType {
	return LBD_NODE
}

func (n *lambdaNode) Pop() *symNode {
	h := n.ps[0]
	n.ps = n.ps[1:]
	return h
}
