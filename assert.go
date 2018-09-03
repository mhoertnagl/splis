package main

import "fmt"

type assertion struct {
	ok  bool
	err *errNode
}

func NewAssertion() *assertion {
	return &assertion{true, nil}
}

func (a *assertion) Failed() bool {
	return !a.ok
}

func (a *assertion) Error() *errNode {
	return a.err
}

func (a *assertion) AssertIntEqual(x int, y int, xName string, yName string) bool {
	if a.ok && x != y {
		msg := fmt.Sprintf("%s [%d] and %s [%d] must be equal.", xName, x, yName, y)
		a.ok = false
		a.err = NewErrNode(msg)
	}
	return a.Failed()
}

func (a *assertion) AssertDivisble(x float64) bool {
	if a.ok && x == 0 {
		msg := fmt.Sprintf("Cannot divide by 0.")
		a.ok = false
		a.err = NewErrNode(msg)
	}
	return a.Failed()
}

func (a *assertion) AssertLen(seq []Node, l int, funName string) bool {
	if a.ok && len(seq) != l {
		msg := fmt.Sprintf("%s requires exactly [%d] arguments.", funName, l)
		a.ok = false
		a.err = NewErrNode(msg)
	}
	return a.Failed()
}

func (a *assertion) AssertMin(seq []Node, min int, funName string) bool {
	if a.ok && len(seq) < min {
		msg := fmt.Sprintf("%s requires at least [%d] arguments.", funName, min)
		a.ok = false
		a.err = NewErrNode(msg)
	}
	return a.Failed()
}

func (a *assertion) AssertType(n Node, t NodeType, argName string) bool {
	if a.ok && n.Type() != t {
		msg := fmt.Sprintf("%s must be of type [%v] but is [%v].", argName, t.Name(), n.Type().Name())
		a.ok = false
		a.err = NewErrNode(msg)
	}
	return a.Failed()
}
