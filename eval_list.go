package main

import "fmt"

func (vm *vm) evalList(e Env, as []Node) Node {
	n := NewQExprNode()
	for _, a := range as {
		n.Push(a)
	}
	return n
}

func (vm *vm) evalHead(e Env, as []Node) Node {
	t := NewAssertion()

	if t.AssertLen(as, 1, "Head") {
		return t.Error()
	}

	if t.AssertType(as[0], QXP_NODE, "Argument of head") {
		return t.Error()
	}

	n := as[0].(*seqNode)

	if t.AssertMin(n.cells, 1, "List provided to head") {
		return t.Error()
	}

	q := NewQExprNode()
	q.Push(n.cells[0])
	return q
}

func (vm *vm) evalTail(e Env, as []Node) Node {
	t := NewAssertion()
	if t.AssertLen(as, 1, "Tail") {
		return t.Error()
	}

	if t.AssertType(as[0], QXP_NODE, "Argument of tail") {
		return t.Error()
	}

	n := as[0].(*seqNode)

	if t.AssertMin(n.cells, 1, "List provided to tail") {
		return t.Error()
	}

	q := NewQExprNode()
	for i := 1; i < len(n.cells); i++ {
		q.Push(n.cells[i])
	}
	return q
}

func (vm *vm) evalJoin(e Env, as []Node) Node {
	t := NewAssertion()

	q := NewQExprNode()
	for i, a := range as {
		if t.AssertType(a, QXP_NODE, fmt.Sprintf("Argument [%d] of tail", i+1)) {
			return t.Error()
		}

		n := a.(*seqNode)
		for _, e := range n.cells {
			q.Push(e)
		}
	}
	return q
}
