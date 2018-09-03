package main

import (
	"fmt"
	"strconv"
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
	Copy() Node
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

func (n *errNode) Copy() Node {
	return n
}

type numNode struct {
	val float64
}

func NewNumNodeFromString(s string) *numNode {
	return NewNumNode(parseNumber(s))
}

func NewNumNode(n float64) *numNode {
	return &numNode{n}
}

func (n *numNode) Type() NodeType {
	return NUM_NODE
}

// Numbers are immutable?
func (n *numNode) Copy() Node {
	return NewNumNode(n.val)
}

func parseNumber(s string) float64 {
	// var n float64
	// if strings.HasPrefix(s, "0b") {
	// 	n, _ = strconv.ParseInt(s[2:], 2, 32)
	// } else if strings.HasPrefix(s, "0x") {
	// 	n, _ = strconv.ParseInt(s[2:], 16, 32)
	// } else {
	// 	n, _ = strconv.ParseInt(s, 10, 32)
	// }
	n, _ := strconv.ParseFloat(s, 64)
	return float64(n)
}

type strNode struct {
	str string
}

// Strings are immutable?
func NewStrNode(s string) *strNode {
	return &strNode{s}
}

func (n *strNode) Type() NodeType {
	return STR_NODE
}

func (n *strNode) Copy() Node {
	return NewStrNode(n.str)
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

// Symbols are immutable?
func (n *symNode) Copy() Node {
	return NewSymNode(n.name)
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

func (n *seqNode) Copy() Node {
	m := &seqNode{n.typ, []Node{}}
	for _, c := range n.cells {
		m.Push(c.Copy())
	}
	return m
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

func (n *funNode) Copy() Node {
	// FunNode is immutable.
	return n
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

func (n *lambdaNode) Copy() Node {
	m := NewLambdaNode(n.env.Copy(), []*symNode{}, n.body.Copy().(*seqNode))
	for _, p := range n.ps {
		m.ps = append(m.ps, p.Copy().(*symNode))
	}
	return m
}

func (n *lambdaNode) Pop() *symNode {
	h := n.ps[0]
	n.ps = n.ps[1:]
	return h
}
