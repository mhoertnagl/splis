package main

import (
	"testing"
)

func assertPrintEqual(t *testing.T, s string) {
	l := NewLexer(s)
	p := NewParser(l)
	n := p.Parse()
	r := printAst(n)
	if r != s {
		t.Errorf("Expected [%v] but got [%v]", s, r)
	}
}

// func TestPrintEmpty(t *testing.T) {
// 	assertPrintEqual(t, "")
// }

func TestPrintNum0(t *testing.T) {
	assertPrintEqual(t, "0")
}

func TestPrintEmptyString(t *testing.T) {
	assertPrintEqual(t, "\"\"")
}

func TestPrintString(t *testing.T) {
	assertPrintEqual(t, "\"splis\"")
}

func TestPrintEmptySExpr(t *testing.T) {
	assertPrintEqual(t, "()")
}

func TestPrintSExpr1(t *testing.T) {
	assertPrintEqual(t, "(+ 1 1)")
}

func TestPrintSExpr2(t *testing.T) {
	assertPrintEqual(t, "(+ (* 1 2) 1)")
}

func TestPrintSExpr3(t *testing.T) {
	assertPrintEqual(t, "(+ (* 1 2) (/ 9 3))")
}

func TestPrintQExpr1(t *testing.T) {
	assertPrintEqual(t, "{}")
}

func TestPrintQExpr2(t *testing.T) {
	assertPrintEqual(t, "{+ 2 2}")
}

func TestPrintQExpr3(t *testing.T) {
	assertPrintEqual(t, "{+ (+ 1 1) 2}")
}

// func TestPrintFun(t *testing.T) {
// 	assertPrintEqual(t, "(+)")
// }
