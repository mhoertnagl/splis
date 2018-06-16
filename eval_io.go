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

// func (vm *vm) evalLoad(e Env, as []Node) Node {
// 	t := NewAssertion()
// 	t.AssertMin(as, 1, "Load")

// 	if t.Failed() {
// 		return t.Error()
// 	}

// 	for _, a := range as {
// 		v := vm.eval(e, a)

// 		if t.AssertType(v, STR_NODE, "module") {
// 			return t.Error()
// 		}

// 	}
// 	return NewSExprNode()
// }

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
	n := p.Parse()
	return vm.eval(e, n)
}

// func (vm *vm) evalExecute(e Env, as []Node) Node {
// 	t := NewAssertion()
// 	t.AssertMin(as, 1, "Execute")

// 	if t.Failed() {
// 		return t.Error()
// 	}

// 	for _, a := range as {
// 		v := vm.eval(e, a)

// 		if t.AssertType(v, STR_NODE, "source text") {
// 			return t.Error()
// 		}
// 		s := v.(*strNode)
// 		l := NewLexer(s.str)
// 		p := NewParser(l)
// 		n := p.Parse()
// 		vm.eval(e, n)
// 	}
// 	return NewSExprNode()
// }
