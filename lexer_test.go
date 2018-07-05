package main

import (
	"testing"
)

func TestLexEmpty(t *testing.T) {
	l := NewLexer("")
	assertLex(t, l.Next(), NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexNum0(t *testing.T) {
	l := NewLexer("0")
	assertLex(t, l.Next(), NewToken(NUM, "0", NewPos(1, 1)))
	assertLex(t, l.Next(), NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexNum0123456789(t *testing.T) {
	l := NewLexer("0123456789")
	assertLex(t, l.Next(), NewToken(NUM, "0123456789", NewPos(1, 1)))
	assertLex(t, l.Next(), NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexNumBin(t *testing.T) {
	l := NewLexer("0b1001")
	assertLex(t, l.Next(), NewToken(NUM, "0b1001", NewPos(1, 1)))
	assertLex(t, l.Next(), NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexNumHex(t *testing.T) {
	l := NewLexer("0x0123456789ABCDEF")
	assertLex(t, l.Next(), NewToken(NUM, "0x0123456789ABCDEF", NewPos(1, 1)))
	assertLex(t, l.Next(), NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexEmptyString(t *testing.T) {
	l := NewLexer("\"\"")
	assertLex(t, l.Next(), NewToken(STR, "", NewPos(1, 1)))
	assertLex(t, l.Next(), NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexSimpleString(t *testing.T) {
	l := NewLexer("\"abc\"")
	assertLex(t, l.Next(), NewToken(STR, "abc", NewPos(1, 1)))
	assertLex(t, l.Next(), NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexEscapedNewlineString(t *testing.T) {
	l := NewLexer("\"a\\nbc\"")
	assertLex(t, l.Next(), NewToken(STR, "a\nbc", NewPos(1, 1)))
	assertLex(t, l.Next(), NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexEscapedBackslashString(t *testing.T) {
	l := NewLexer("\"a\\\\bc\"")
	assertLex(t, l.Next(), NewToken(STR, "a\\bc", NewPos(1, 1)))
	assertLex(t, l.Next(), NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexEscapedUnrecognizedString(t *testing.T) {
	l := NewLexer("\"a\\xbc\"")
	assertLex(t, l.Next(), NewToken(ERR, "Unrecognized escape sequence [\\x].", NewPos(1, 1)))
}

func TestLexOpenString(t *testing.T) {
	l := NewLexer("\"abc")
	assertLex(t, l.Next(), NewToken(ERR, "Missing closing quote.", NewPos(1, 1)))
}

func TestLexPars(t *testing.T) {
	l := NewLexer("()")
	assertLex(t, l.Next(), NewToken(PAR, "(", NewPos(1, 1)))
	assertLex(t, l.Next(), NewToken(PAR, ")", NewPos(1, 2)))
	assertLex(t, l.Next(), NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexSymbol(t *testing.T) {
	l := NewLexer("x")
	assertLex(t, l.Next(), NewToken(SYM, "x", NewPos(1, 1)))
	assertLex(t, l.Next(), NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexSExrp(t *testing.T) {
	l := NewLexer("(< a b)")
	assertLex(t, l.Next(), NewToken(PAR, "(", NewPos(1, 1)))
	assertLex(t, l.Next(), NewToken(SYM, "<", NewPos(1, 2)))
	assertLex(t, l.Next(), NewToken(SYM, "a", NewPos(1, 4)))
	assertLex(t, l.Next(), NewToken(SYM, "b", NewPos(1, 6)))
	assertLex(t, l.Next(), NewToken(PAR, ")", NewPos(1, 7)))
	assertLex(t, l.Next(), NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexComplexSExrp(t *testing.T) {
	s := `
		(if (> a 0)
			(+ a 1)
			0)
	`
	assertLexSeq(t, s,
		NewToken(PAR, "(", NewPos(1, 1)),
		NewToken(SYM, "if", NewPos(1, 2)),
		NewToken(PAR, "(", NewPos(1, 1)),
		NewToken(SYM, ">", NewPos(1, 4)),
		NewToken(SYM, "a", NewPos(1, 6)),
		NewToken(NUM, "0", NewPos(1, 7)),
		NewToken(PAR, ")", NewPos(1, 1)),
		NewToken(PAR, "(", NewPos(1, 1)),
		NewToken(SYM, "+", NewPos(1, 4)),
		NewToken(SYM, "a", NewPos(1, 6)),
		NewToken(NUM, "1", NewPos(1, 7)),
		NewToken(PAR, ")", NewPos(1, 1)),
		NewToken(NUM, "0", NewPos(1, 7)),
		NewToken(PAR, ")", NewPos(1, 1)),
		NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexCurlyBraces(t *testing.T) {
	assertLexSeq(t, "{}",
		NewToken(CBR, "{", NewPos(1, 1)),
		NewToken(CBR, "}", NewPos(1, 2)),
		NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexDef(t *testing.T) {
	assertLexSeq(t, "(def {a} 100)",
		NewToken(PAR, "(", NewPos(1, 1)),
		NewToken(SYM, "def", NewPos(1, 2)),
		NewToken(CBR, "{", NewPos(1, 1)),
		NewToken(SYM, "a", NewPos(1, 2)),
		NewToken(CBR, "}", NewPos(1, 2)),
		NewToken(NUM, "100", NewPos(1, 7)),
		NewToken(PAR, ")", NewPos(1, 2)),
		NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexComment1(t *testing.T) {
	assertLexSeq(t, "// Test",
		NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexComment2(t *testing.T) {
	assertLexSeq(t, "1 // Test",
		NewToken(NUM, "1", NewPos(1, 1)),
		NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexComment3(t *testing.T) {
	assertLexSeq(t, "// Test\n1",
		NewToken(NUM, "1", NewPos(1, 1)),
		NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexMLComment1(t *testing.T) {
	assertLexSeq(t, "/* Test\nTest */",
		NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexMLComment2(t *testing.T) {
	assertLexSeq(t, "1 /* Test\nTest */",
		NewToken(NUM, "1", NewPos(1, 1)),
		NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexMLComment3(t *testing.T) {
	assertLexSeq(t, "/* Test\nTest */ 1",
		NewToken(NUM, "1", NewPos(1, 1)),
		NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexMLComment4(t *testing.T) {
	assertLexSeq(t, "/* Test\nTest */ 1 /* Test */ 2",
		NewToken(NUM, "1", NewPos(1, 1)),
		NewToken(NUM, "2", NewPos(1, 1)),
		NewToken(EOF, "", NewPos(1, 1)))
}

func TestLexMLComment5(t *testing.T) {
	assertLexSeq(t, "/* Test\nTest ",
		NewToken(ERR, "Stray multi line comment. Missing [*/].", NewPos(1, 1)))
}

func TestLexUnrecognized(t *testing.T) {
	assertLexSeq(t, "\f",
		NewToken(ERR, "Unrecognized character [].", NewPos(1, 1)))
}

func assertLex(t *testing.T, a Token, e Token) {
	if a.Kind() != e.Kind() {
		t.Errorf("Unexpected token kind [%d]. Expecting kind [%d]", a.Kind(), e.Kind())
	}
	if a.Value() != e.Value() {
		t.Errorf("Unexpected token value [%s]. Expecting value [%s]", a.Value(), e.Value())
	}
}

func assertLexSeq(t *testing.T, s string, es ...Token) {
	l := NewLexer(s)
	for _, e := range es {
		assertLex(t, l.Next(), e)
	}
}
