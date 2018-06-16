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
	assertEvalEqual(t, "(undefined)", "Unbound symbol [undefined].\n")
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

func TestEvalSum1f(t *testing.T) {
	assertEvalEqual(t, "(+ 1 {})", "Cannot add non-number [{}].\n")
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

func TestEvalTooManyArgsLambda(t *testing.T) {
	assertEvalEqual(t, "((lambda {a b} {+ a b}) 1 2 3)", "Too many arguments [(lambda {a b} {+ a b})].\n")
}

func TestEvalEqual0(t *testing.T) {
	assertEvalEqual(t, "(== 0 {})", "0")
}

func TestEvalEqual1(t *testing.T) {
	assertEvalEqual(t, "(== 0 0)", "1")
}

func TestEvalEqual1f(t *testing.T) {
	assertEvalEqual(t, "(== 0 1)", "0")
}

func TestEvalEqual2(t *testing.T) {
	assertEvalEqual(t, "(== {x} {x})", "1")
}

func TestEvalEqual2f(t *testing.T) {
	assertEvalEqual(t, "(== {x} {y})", "0")
}

func TestEvalEqual3(t *testing.T) {
	assertEvalEqual(t, "(== {(+ 1 2)} {(+ 1 2)})", "1")
}

func TestEvalEqual3f(t *testing.T) {
	assertEvalEqual(t, "(== {(+ 1 2)} {(+ 2 1)})", "0")
}

func TestEvalEqual3f2(t *testing.T) {
	assertEvalEqual(t, "(== {(+ 1 2)} {(+ 1)})", "0")
}

func TestEvalEqual4t1(t *testing.T) {
	assertEvalEqual(t, "(== (lambda {a b} {+ a b}) (lambda {a b} {+ a b}))", "1")
}

func TestEvalEqual4t2(t *testing.T) {
	assertEvalEqual(t, "(== (lambda {a b} {+ a b}) (lambda {c b} {+ a b}))", "0")
}

func TestEvalEqual4f1(t *testing.T) {
	assertEvalEqual(t, "(== (lambda {a} {+ a b}) (lambda {a b} {+ a b}))", "0")
}

func TestEvalEqual4f2(t *testing.T) {
	assertEvalEqual(t, "(== (lambda {a b} {+ a}) (lambda {a b} {+ a b}))", "0")
}

func TestEvalEqual5(t *testing.T) {
	assertEvalEqual(t, "(== + +)", "1")
}

func TestEvalEqual5f(t *testing.T) {
	assertEvalEqual(t, "(== + <)", "0")
}

func TestEvalNE(t *testing.T) {
	assertEvalEqual(t, "(!= 0 0)", "0")
}

func TestEvalNEf(t *testing.T) {
	assertEvalEqual(t, "(!= 0 1)", "1")
}

func TestEvalIf1(t *testing.T) {
	assertEvalEqual(t, "(if (< 1 2) {+ 1 2} {4})", "3")
}

func TestEvalIf2(t *testing.T) {
	assertEvalEqual(t, "(if (< 2 1) {+ 1 2} {4})", "4")
}

func TestEvalIff1(t *testing.T) {
	assertEvalEqual(t, "(if a {+ 1 2} {4})", "First argument of if must be of type [Number] but is [Error].\n")
}

func TestEvalIff2(t *testing.T) {
	assertEvalEqual(t, "(if (< 1 2) (+ 1 2) {4})", "Second argument of if must be of type [Q-Expression] but is [Number].\n")
}

func TestEvalIff3(t *testing.T) {
	assertEvalEqual(t, "(if (< 1 2) {+ 1 2} 4)", "Third argument of if must be of type [Q-Expression] but is [Number].\n")
}

func TestEvalAnd1(t *testing.T) {
	assertEvalEqual(t, "(&& (< 1 2))", "1")
}

func TestEvalAnd2(t *testing.T) {
	assertEvalEqual(t, "(&& (< 2 1))", "0")
}

func TestEvalAnd3(t *testing.T) {
	assertEvalEqual(t, "(&& (< 1 2) (< 2 3))", "1")
}

func TestEvalAnd4(t *testing.T) {
	assertEvalEqual(t, "(&& (< 2 1) (< 2 3))", "0")
}

func TestEvalAnd5(t *testing.T) {
	assertEvalEqual(t, "(&& (< 1 2) (< 3 2))", "0")
}

func TestEvalLoad(t *testing.T) {
	assertEvalEqual(t, "(load \"test/load.splis\")", "\"(def {x} 1)\"")
}

func TestEvalExecute(t *testing.T) {
	assertEvalEqual(t, "(execute \"(+ 1 4)\")", "5")
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
