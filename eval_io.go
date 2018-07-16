package main

import (
	"fmt"
	"io/ioutil"
)

func (vm *vm) evalLoad(e Env, as []Node) Node {
	t := NewAssertion()

	if t.AssertLen(as, 1, "Load") {
		return t.Error()
	}

	if t.AssertType(as[0], STR_NODE, "module") {
		return t.Error()
	}

	p := as[0].(*strNode)
	buf, err := ioutil.ReadFile(p.str)

	if err != nil {
		return NewErrNode("%s", err)
	}
	return NewStrNode(string(buf))
}

func (vm *vm) evalExecute(e Env, as []Node) Node {
	t := NewAssertion()

	if t.AssertLen(as, 1, "Execute") {
		return t.Error()
	}

	if t.AssertType(as[0], STR_NODE, "source text") {
		return t.Error()
	}

	s := as[0].(*strNode)
	l := NewLexer(s.str)
	p := NewParser(l)
	ns := p.Parse()
	return vm.eval(e, ns[0])
}

func (vm *vm) evalPrint(e Env, as []Node) Node {
	t := NewAssertion()
	t.AssertMin(as, 1, "Print")

	if t.Failed() {
		return t.Error()
	}

	for i := 0; i < len(as); i++ {
		v := vm.eval(e, as[i])

		if t.AssertType(v, STR_NODE, "print argument") {
			return t.Error()
		}

		s := v.(*strNode)
		fmt.Print(s.str)
	}
	return NewSExprNode()
}

func (vm *vm) evalError(e Env, as []Node) Node {
	t := NewAssertion()
	t.AssertLen(as, 1, "Error")

	if t.Failed() {
		return t.Error()
	}

	v := vm.eval(e, as[0])

	if t.AssertType(v, STR_NODE, "print argument") {
		return t.Error()
	}

	s := v.(*strNode)
	return NewErrNode(s.str)
}
