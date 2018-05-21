package main

import (
	"testing"
)

func TestParseEmpty(t *testing.T) {
	l := NewLexer("")
	p := NewParser(l)
	p.Parse()
}

func TestParseNum(t *testing.T) {
	l := NewLexer("0")
	p := NewParser(l)
	p.Parse()
}

func TestParseSym(t *testing.T) {
	l := NewLexer("a")
	p := NewParser(l)
	p.Parse()
}

func TestParseEmptySubExpr(t *testing.T) {
	l := NewLexer("()")
	p := NewParser(l)
	p.Parse()
}

func TestParseEmpty2SubExpr(t *testing.T) {
	l := NewLexer("((()))")
	p := NewParser(l)
	p.Parse()
}

// func TestParseIncompleteEmptySubExpr1(t *testing.T) {
// 	l := NewLexer("(")
// 	p := NewParser(l)
// 	p.Parse()
// }

// func TestParseIncompleteEmptySubExpr2(t *testing.T) {
// 	l := NewLexer(")")
// 	p := NewParser(l)
// 	p.Parse()
// }

func TestParseSingleElemSubExpr(t *testing.T) {
	l := NewLexer("(x)")
	p := NewParser(l)
	p.Parse()
}

func TestParse3ElemSubExpr(t *testing.T) {
	l := NewLexer("(+ 5 6)")
	p := NewParser(l)
	p.Parse()
}

func TestParseSubElem1SubExpr(t *testing.T) {
	l := NewLexer("(+ (* 3 4 3) 6)")
	p := NewParser(l)
	p.Parse()
}

func TestParseSubElem2SubExpr(t *testing.T) {
	l := NewLexer("(+ 4 (/ 9 3))")
	p := NewParser(l)
	p.Parse()
}

func TestParseSubElemSubExpr(t *testing.T) {
	l := NewLexer("(+ (* 1 2 3 4 5) (/ 9 (+ 2 1)))")
	p := NewParser(l)
	p.Parse()
}
