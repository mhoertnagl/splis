package main

import (
	"testing"
)

func assertNum(t *testing.T, r Node, n int32) {
	rr, ok := r.(NumNode)
	if !ok {
		t.Errorf("Expected type [NumNode] but got type [%v]", rr)
	}
	if rr.Value() != n {
		t.Errorf("Expected number [%v] but got [%v]", n, rr.Value())
	}
}

func assertSym(t *testing.T, r Node, name string) {
	rr, ok := r.(SymNode)
	if !ok {
		t.Errorf("Expected type [SymNode] but got type [%v]", rr)
	}
	if rr.Name() != name {
		t.Errorf("Expected symbol name [%v] but got [%v]", name, rr.Name())
	}
}

func assertSExpr(t *testing.T, r Node, len int) SExprNode {
	rr, ok := r.(SExprNode)
	if !ok {
		t.Errorf("Expected type [SExprNode] but got type [%v]", rr)
	}
	if rr.Len() != len {
		t.Errorf("Expected S-Expr length of [%v] but got [%v]", len, rr.Len())
	}
	return rr
}

func TestParseEmpty(t *testing.T) {
	l := NewLexer("")
	p := NewParser(l)
	r := p.Parse()
	if r != nil {
		t.Errorf("Expected [%v] but got [%v]", r, nil)
	}
}

func TestParseNum(t *testing.T) {
	l := NewLexer("0")
	p := NewParser(l)
	r := p.Parse()
	assertNum(t, r, 0)
}

func TestParseBinNum(t *testing.T) {
	l := NewLexer("0b101")
	p := NewParser(l)
	r := p.Parse()
	assertNum(t, r, 5)
}

func TestParseHexNum(t *testing.T) {
	l := NewLexer("0xABC")
	p := NewParser(l)
	r := p.Parse()
	assertNum(t, r, 2748)
}

func TestParseSym(t *testing.T) {
	l := NewLexer("a")
	p := NewParser(l)
	r := p.Parse()
	assertSym(t, r, "a")
}

func TestParseEmptySubExpr(t *testing.T) {
	l := NewLexer("()")
	p := NewParser(l)
	r := p.Parse()
	assertSExpr(t, r, 0)
}

func TestParseEmpty2SubExpr(t *testing.T) {
	l := NewLexer("((()))")
	p := NewParser(l)
	r := p.Parse()
	s0 := assertSExpr(t, r, 1)
	s1 := assertSExpr(t, s0.Cell(0), 1)
	assertSExpr(t, s1.Cell(0), 0)
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
	r := p.Parse()
	s := assertSExpr(t, r, 1)
	assertSym(t, s.Cell(0), "x")
}

func TestParse3ElemSubExpr(t *testing.T) {
	l := NewLexer("(+ 5 6)")
	p := NewParser(l)
	r := p.Parse()
	s := assertSExpr(t, r, 3)
	assertSym(t, s.Cell(0), "+")
	assertNum(t, s.Cell(1), 5)
	assertNum(t, s.Cell(2), 6)
}

func TestParseSubElem1SubExpr(t *testing.T) {
	l := NewLexer("(+ (* 3 4 3) 6)")
	p := NewParser(l)
	r := p.Parse()
	s0 := assertSExpr(t, r, 3)
	assertSym(t, s0.Cell(0), "+")
	s1 := assertSExpr(t, s0.Cell(1), 4)
	assertSym(t, s1.Cell(0), "*")
	assertNum(t, s1.Cell(1), 3)
	assertNum(t, s1.Cell(2), 4)
	assertNum(t, s1.Cell(3), 3)
	assertNum(t, s0.Cell(2), 6)
}

func TestParseSubElem2SubExpr(t *testing.T) {
	l := NewLexer("(+ 4 (/ 9 3))")
	p := NewParser(l)
	r := p.Parse()
	s0 := assertSExpr(t, r, 3)
	assertSym(t, s0.Cell(0), "+")
	assertNum(t, s0.Cell(1), 4)
	s1 := assertSExpr(t, s0.Cell(2), 3)
	assertSym(t, s1.Cell(0), "/")
	assertNum(t, s1.Cell(1), 9)
	assertNum(t, s1.Cell(2), 3)
}

func TestParseSubElemSubExpr(t *testing.T) {
	l := NewLexer("(+ (* 1 2 3 4 5) (/ 9 (+ 2 1)))")
	p := NewParser(l)
	p.Parse()
}
