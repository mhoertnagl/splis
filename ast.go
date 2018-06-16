package main

import (
	"fmt"
	"strconv"
	"strings"
)

type NodeType int

const (
	ERR_NODE NodeType = iota
	NUM_NODE
	STR_NODE
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
	case STR_NODE:
		return "String"
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

func NewErrNode(msg string, args ...interface{}) *errNode {
	return &errNode{fmt.Sprintf(msg, args...)}
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

type strNode struct {
	str string
}

func NewStrNode(s string) *strNode {
	return &strNode{s}
}

func (n *strNode) Type() NodeType {
	return STR_NODE
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

type seqNode struct {
	typ   NodeType
	cells []Node
}

func NewSExprNode() *seqNode {
	return &seqNode{SXP_NODE, []Node{}}
}

func NewQExprNode() *seqNode {
	return &seqNode{QXP_NODE, []Node{}}
}

func (n *seqNode) Type() NodeType {
	return n.typ
}

func (n *seqNode) Push(c Node) {
	n.cells = append(n.cells, c)
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
	body *seqNode
}

func NewLambdaNode(e Env, ps []*symNode, body *seqNode) *lambdaNode {
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
