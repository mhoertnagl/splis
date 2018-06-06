package main

import (
	"testing"
)

func TestEvalEmpty(t *testing.T) {
	assertEvalEqual(t, "()", "()")
}

func TestEvalNum(t *testing.T) {
	assertEvalEqual(t, "99", "99")
}

func TestEvalSym(t *testing.T) {
	assertEvalEqual(t, "false", "0")
}

func TestEvalUndefinedFunction(t *testing.T) {
	assertEvalEqual(t, "(undefined)", "")
}

func TestEvalSingleNumInSExpr(t *testing.T) {
	assertEvalEqual(t, "(1)", "1")
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

func TestInvariantQExpr(t *testing.T) {
	assertEvalEqual(t, "{(+ 1 1)}", "{(+ 1 1)}")
}

func TestLT0(t *testing.T) {
	assertEvalEqual(t, "(< 1 2)", "1")
}

func TestLT1(t *testing.T) {
	assertEvalEqual(t, "(< 1 1)", "0")
}

func TestLT2(t *testing.T) {
	assertEvalEqual(t, "(< 1 0)", "0")
}

func TestLT3(t *testing.T) {
	assertEvalEqual(t, "(< 1 0 0)", "< requires exactly [2] arguments.\n")
}

func TestLT4(t *testing.T) {
	assertEvalEqual(t, "(< {} 0)", "First argument of < must be of type [Number] but is [Q-Expression].\n")
}

func TestLT5(t *testing.T) {
	assertEvalEqual(t, "(< 1 {})", "Second argument of < must be of type [Number] but is [Q-Expression].\n")
}

func TestEvalNum2(t *testing.T) {
	assertEvalEqual(t, "(eval 1)", "1")
}

func TestEvalSym2(t *testing.T) {
	assertEvalEqual(t, "(eval true)", "1")
}

func TestEvalSExpr(t *testing.T) {
	assertEvalEqual(t, "(eval (+ 3 3 3))", "9")
}

func TestEvalIdQExpr(t *testing.T) {
	assertEvalEqual(t, "(eval {1})", "1")
}

func TestEvalQExpr(t *testing.T) {
	assertEvalEqual(t, "(eval {+ 1 1})", "2")
}

func TestEvalIDLambda(t *testing.T) {
	assertEvalEqual(t, "((lambda {a} {a}) 666)", "666")
}

func TestEvalSimpleLambda(t *testing.T) {
	assertEvalEqual(t, "((lambda {a b} {+ a b}) 1 2)", "3")
}

func TestEvalPartialLambda(t *testing.T) {
	assertEvalEqual(t, "(((lambda {a b} {+ a b}) 1) 2)", "3")
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
