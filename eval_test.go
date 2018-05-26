package main

import (
	"testing"
)

func TestEvalEmpty(t *testing.T) {
	assertEvalEqual(t, "()", "()")
}

func TestEvalUndefinedFunction(t *testing.T) {
	assertEvalEqual(t, "(undefined)", "(undefined)")
}

func TestEvalSum1(t *testing.T) {
	assertEvalEqual(t, "(+ 1)", "1")
}

func TestEvalSum2(t *testing.T) {
	assertEvalEqual(t, "(+ 1 1)", "2")
}

func TestEvalSum3(t *testing.T) {
	assertEvalEqual(t, "(+ 1 1 1)", "3")
}

func TestEvalSum4(t *testing.T) {
	assertEvalEqual(t, "(+ (+ 1 1) (+ 1 1))", "4")
}

func TestEvalQExpr(t *testing.T) {
	assertEvalEqual(t, "{(+ 1 1)}", "{(+ 1 1)}")
}

func assertEvalEqual(t *testing.T, s string, e string) {
	l := NewLexer(s)
	p := NewParser(l)
	n := p.Parse()
	vm := NewVM()
	res := vm.Eval(n)
	r := printAst(res)
	if r != e {
		t.Errorf("Expected [%v] but got [%v]", e, r)
	}
}
