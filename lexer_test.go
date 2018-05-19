package main

import (
	"testing"
)

func assert(t *testing.T, a Token, e Token) {
	if a.Kind() != e.Kind() {
		t.Errorf("Unexpected token kind [%d]. Expecting kind [%d]", a.Kind(), e.Kind())
	}
	if a.Value() != e.Value() {
		t.Errorf("Unexpected token value [%s]. Expecting value [%s]", a.Value(), e.Value())
	}
}

func TestLexEmpty(t *testing.T) {
	l := NewLexer("")
	assert(t, l.Next(), NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexNum0(t *testing.T) {
	l := NewLexer("0")
	assert(t, l.Next(), NewToken(NUM, "0", NewPos(1, 1)))
}

func TestLexNum0123456789(t *testing.T) {
	l := NewLexer("0123456789")
	assert(t, l.Next(), NewToken(NUM, "0123456789", NewPos(1, 1)))
}

func TestLexNumBin(t *testing.T) {
	l := NewLexer("0b1001")
	assert(t, l.Next(), NewToken(NUM, "0b1001", NewPos(1, 1)))
}

func TestLexNumHex(t *testing.T) {
	l := NewLexer("0x0123456789ABCDEF")
	assert(t, l.Next(), NewToken(NUM, "0x0123456789ABCDEF", NewPos(1, 1)))
}

func TestLexPars(t *testing.T) {
	l := NewLexer("()")
	assert(t, l.Next(), NewToken(PAR, "(", NewPos(1, 1)))
	assert(t, l.Next(), NewToken(PAR, ")", NewPos(1, 2)))
}

func TestLexSymbol(t *testing.T) {
	l := NewLexer("x")
	assert(t, l.Next(), NewToken(SYM, "x", NewPos(1, 1)))
}

func TestLexSExrp(t *testing.T) {
	l := NewLexer("(< a b)")
	assert(t, l.Next(), NewToken(PAR, "(", NewPos(1, 1)))
	assert(t, l.Next(), NewToken(SYM, "<", NewPos(1, 2)))
	assert(t, l.Next(), NewToken(SYM, "a", NewPos(1, 4)))
	assert(t, l.Next(), NewToken(SYM, "b", NewPos(1, 6)))
	assert(t, l.Next(), NewToken(PAR, ")", NewPos(1, 7)))
}

func TestLexComplexSExrp(t *testing.T) {
	l := NewLexer(`
		(if (> a 0)
			(+ a 1)
			0)
	`)
	assert(t, l.Next(), NewToken(PAR, "(", NewPos(1, 1)))
	assert(t, l.Next(), NewToken(SYM, "if", NewPos(1, 2)))
	assert(t, l.Next(), NewToken(PAR, "(", NewPos(1, 1)))
	assert(t, l.Next(), NewToken(SYM, ">", NewPos(1, 4)))
	assert(t, l.Next(), NewToken(SYM, "a", NewPos(1, 6)))
	assert(t, l.Next(), NewToken(NUM, "0", NewPos(1, 7)))
	assert(t, l.Next(), NewToken(PAR, ")", NewPos(1, 1)))
	assert(t, l.Next(), NewToken(PAR, "(", NewPos(1, 1)))
	assert(t, l.Next(), NewToken(SYM, "+", NewPos(1, 4)))
	assert(t, l.Next(), NewToken(SYM, "a", NewPos(1, 6)))
	assert(t, l.Next(), NewToken(NUM, "1", NewPos(1, 7)))
	assert(t, l.Next(), NewToken(PAR, ")", NewPos(1, 1)))
	assert(t, l.Next(), NewToken(NUM, "0", NewPos(1, 7)))
	assert(t, l.Next(), NewToken(PAR, ")", NewPos(1, 1)))
}
