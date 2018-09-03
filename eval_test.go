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
	assertEvalEqual(t, "(undefined)", "Error: Unbound symbol [undefined].\n")
}

func TestEvalSingleNumInSExpr(t *testing.T) {
	assertEvalEqual(t, "(1)", "1")
}

func TestEvalSum0(t *testing.T) {
	assertEvalEqual(t, "(+)", "0")
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
	assertEvalEqual(t, "(+ 1 {})", "Error: Argument of + must be of type [Number] but is [Q-Expression].\n")
}

func TestEvalDiff0(t *testing.T) {
	assertEvalEqual(t, "(-)", "0")
}

func TestEvalDiff1(t *testing.T) {
	assertEvalEqual(t, "(- 5)", "-5")
}

func TestEvalDiff2(t *testing.T) {
	assertEvalEqual(t, "(- 5 2)", "3")
}

func TestEvalDiff3(t *testing.T) {
	assertEvalEqual(t, "(- 5 2 3)", "0")
}

func TestEvalMul0(t *testing.T) {
	assertEvalEqual(t, "(*)", "1")
}

func TestEvalMul1(t *testing.T) {
	assertEvalEqual(t, "(* 5)", "5")
}

func TestEvalMul2(t *testing.T) {
	assertEvalEqual(t, "(* 5 3)", "15")
}

func TestEvalDiv0(t *testing.T) {
	assertEvalEqual(t, "(/)", "1")
}

func TestEvalDiv1(t *testing.T) {
	assertEvalEqual(t, "(/ 5)", "0.2")
}

func TestEvalDiv2(t *testing.T) {
	assertEvalEqual(t, "(/ 10 2)", "5")
}

func TestEvalDiv3(t *testing.T) {
	assertEvalEqual(t, "(/ 15 3 5)", "1")
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
	assertEvalEqual(t, "(< 1 0 0)", "Error: < requires exactly [2] arguments.\n")
}

func TestLT4(t *testing.T) {
	assertEvalEqual(t, "(< {} 0)", "Error: First argument of < must be of type [Number] but is [Q-Expression].\n")
}

func TestLT5(t *testing.T) {
	assertEvalEqual(t, "(< 1 {})", "Error: Second argument of < must be of type [Number] but is [Q-Expression].\n")
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

func TestEvalList0(t *testing.T) {
	assertEvalEqual(t, "(list)", "{}")
}

func TestEvalList1(t *testing.T) {
	assertEvalEqual(t, "(list 1)", "{1}")
}

func TestEvalList2(t *testing.T) {
	assertEvalEqual(t, "(list 1 {a})", "{1 {a}}")
}

func TestEvalHead0(t *testing.T) {
	assertEvalEqual(t, "(head {})", "Error: List provided to head requires at least [1] arguments.\n")
}

func TestEvalHead1(t *testing.T) {
	assertEvalEqual(t, "(head {} {})", "Error: Head requires exactly [1] arguments.\n")
}

func TestEvalHead2(t *testing.T) {
	assertEvalEqual(t, "(head 1)", "Error: Argument of head must be of type [Q-Expression] but is [Number].\n")
}

func TestEvalHead3(t *testing.T) {
	assertEvalEqual(t, "(head {1})", "{1}")
}

func TestEvalHead4(t *testing.T) {
	assertEvalEqual(t, "(head {1 2})", "{1}")
}

func TestEvalTail0(t *testing.T) {
	assertEvalEqual(t, "(tail {})", "Error: List provided to tail requires at least [1] arguments.\n")
}

func TestEvalTail1(t *testing.T) {
	assertEvalEqual(t, "(tail {} {})", "Error: Tail requires exactly [1] arguments.\n")
}

func TestEvalTail2(t *testing.T) {
	assertEvalEqual(t, "(tail 1)", "Error: Argument of tail must be of type [Q-Expression] but is [Number].\n")
}

func TestEvalTail3(t *testing.T) {
	assertEvalEqual(t, "(tail {1})", "{}")
}

func TestEvalTail4(t *testing.T) {
	assertEvalEqual(t, "(tail {1 2})", "{2}")
}

func TestEvalTail5(t *testing.T) {
	assertEvalEqual(t, "(tail {1 2 3})", "{2 3}")
}

func TestEvalJoin0(t *testing.T) {
	assertEvalEqual(t, "(join)", "{}")
}

func TestEvalJoin1(t *testing.T) {
	assertEvalEqual(t, "(join {})", "{}")
}

func TestEvalJoin2(t *testing.T) {
	assertEvalEqual(t, "(join {} {})", "{}")
}

func TestEvalJoin3(t *testing.T) {
	assertEvalEqual(t, "(join {a} {})", "{a}")
}

func TestEvalJoin4(t *testing.T) {
	assertEvalEqual(t, "(join {} {b})", "{b}")
}

func TestEvalJoin5(t *testing.T) {
	assertEvalEqual(t, "(join {a} {b})", "{a b}")
}

func TestEvalJoin6(t *testing.T) {
	assertEvalEqual(t, "(join {a b} {c})", "{a b c}")
}

func TestEvalJoin7(t *testing.T) {
	assertEvalEqual(t, "(join {a} {b c})", "{a b c}")
}

func TestEvalJoin8(t *testing.T) {
	assertEvalEqual(t, "(join {a b} {c d})", "{a b c d}")
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
	assertEvalEqual(t, "((lambda {a b} {+ a b}) 1 2 3)", "Error: Too many arguments [(lambda {a b} {+ a b})].\n")
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

func TestEvalEqual6(t *testing.T) {
	assertEvalEqual(t, "(== \"abc\" \"abc\")", "1")
}

func TestEvalEqual6f(t *testing.T) {
	assertEvalEqual(t, "(== \"abc\" \"xy\")", "0")
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

func TestEvalIf3(t *testing.T) {
	assertEvalEqual(t, "(if (< 2 1) {+ 1 2} {+ 4 1})", "5")
}

func TestEvalIff1(t *testing.T) {
	assertEvalEqual(t, "(if a {+ 1 2} {4})", "Error: First argument of if must be of type [Number] but is [Error].\n")
}

func TestEvalIff2(t *testing.T) {
	assertEvalEqual(t, "(if (< 1 2) (+ 1 2) {4})", "Error: Second argument of if must be of type [Q-Expression] but is [Number].\n")
}

func TestEvalIff3(t *testing.T) {
	assertEvalEqual(t, "(if (< 1 2) {+ 1 2} 4)", "Error: Third argument of if must be of type [Q-Expression] but is [Number].\n")
}

func TestEvalAnd0(t *testing.T) {
	assertEvalEqual(t, "(&&)", "1")
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

func TestEvalOr0(t *testing.T) {
	assertEvalEqual(t, "(||)", "0")
}

func TestEvalOr1(t *testing.T) {
	assertEvalEqual(t, "(|| (< 1 2))", "1")
}

func TestEvalOr2(t *testing.T) {
	assertEvalEqual(t, "(|| (< 2 1))", "0")
}

func TestEvalOr3(t *testing.T) {
	assertEvalEqual(t, "(|| (< 1 2) (< 2 3))", "1")
}

func TestEvalOr4(t *testing.T) {
	assertEvalEqual(t, "(|| (< 1 2) (< 3 2))", "1")
}

func TestEvalOr5(t *testing.T) {
	assertEvalEqual(t, "(|| (< 2 1) (< 3 2))", "0")
}

func TestEvalNot1(t *testing.T) {
	assertEvalEqual(t, "(! (< 1 2))", "0")
}

func TestEvalNot2(t *testing.T) {
	assertEvalEqual(t, "(! (< 2 1))", "1")
}

func TestEvalLoad(t *testing.T) {
	assertEvalEqual(t, "(load \"test/load\")", "\"(def {x} 1)\"")
}

func TestEvalExecute(t *testing.T) {
	assertEvalEqual(t, "(execute \"(+ 1 4)\")", "5")
}

func TestEvalExecute1(t *testing.T) {
	assertEvalEqual(t, "(execute \"/**\n* Hallo */\")", "()")
}

func TestEvalExecute2(t *testing.T) {
	assertEvalEqual(t, "(execute \"/**\n* Hallo */ 1\")", "1")
}

func TestEvalExecute3(t *testing.T) {
	assertEvalEqual(t, "(execute \"/**\n* Hallo */ 1 3\")", "3")
}

// func TestEvalExecute99(t *testing.T) {
// 	assertEvalEqual(t, "(execute (load \"lib/prelude.splis\"))", "5")
// }

func TestEvalPrint(t *testing.T) {
	assertEvalEqual(t, "(print \"a\" \"b\" \"c\\n\")", "()")
}

func TestEvalError(t *testing.T) {
	assertEvalEqual(t, "(error \"Oops\")", "Error: Oops\n")
}

func TestEvalMixed1(t *testing.T) {
	assertEvalEqual(t, "(head (tail (list 1 2 3 4 5)))", "{2}")
}

func TestEvalMixed2(t *testing.T) {
	assertEvalEqual(t, "(head (tail (tail (list 1 2 3 4 5))))", "{3}")
}

func TestEvalMixed3(t *testing.T) {
	// s := `
	// (def {len} (lambda {l} {
	//   if (== l {})
	//     {0}
	//     {+ 1 (len (tail l))}}))
	//
	// (len {1})
	// `
	assertEvalEqual(t, `(execute (load "test/debug"))`, "1")
}

func assertEvalEqual(t *testing.T, s string, e string) {
	l := NewLexer(s)
	p := NewParser(l)
	n := p.Parse()
	vm := NewVM()
	res := vm.Eval(n[0])
	r := printAst(res)
	if r != e {
		t.Errorf("Expected [%v] but got [%v]", e, r)
	}
}
