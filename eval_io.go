package main

import (
	"io/ioutil"
)

func (vm *vm) evalLoad(e Env, as []Node) Node {
	t := NewAssertion()
	t.AssertLen(as, 1, "Load")

	if t.Failed() {
		return t.Error()
	}

	v := vm.eval(e, as[0])

	if t.AssertType(v, STR_NODE, "module") {
		return t.Error()
	}

	p := v.(*strNode)
	buf, err := ioutil.ReadFile(p.str)

	if err != nil {
		return NewErrNode("%s", err)
	}
	return NewStrNode(string(buf))
}

func (vm *vm) evalExecute(e Env, as []Node) Node {
	t := NewAssertion()
	t.AssertLen(as, 1, "Execute")

	if t.Failed() {
		return t.Error()
	}

	v := vm.eval(e, as[0])

	if t.AssertType(v, STR_NODE, "source text") {
		return t.Error()
	}

	s := v.(*strNode)
	l := NewLexer(s.str)
	p := NewParser(l)
	ns := p.Parse()
	return vm.eval(e, ns[0])
}
