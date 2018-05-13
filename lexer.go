package main

import (
	"fmt"
	"strings"
)

type Kind int

const eof = rune(0)

const (
	PAR Kind = iota
	NUM
	SYM
	EOF
	ERR
)

type Pos interface {
	Line() int
	Col() int
}

type pos struct {
	line int
	col  int
}

func NewPos(line int, col int) Pos {
	return &pos{line, col}
}

func (p *pos) Line() int {
	return p.line
}

func (p *pos) Col() int {
	return p.col
}

func (p *pos) String() string {
	return fmt.Sprintf("[%d:%d]", p.line, p.col)
}

type Token interface {
	Kind() Kind
	Value() string
	Pos() Pos
}

type token struct {
	kind Kind
	val  string
	pos  Pos
}

func NewToken(kind Kind, val string, pos Pos) Token {
	return &token{kind, val, pos}
}

func (t *token) Kind() Kind {
	return t.kind
}

func (t *token) Value() string {
	return t.val
}

func (t *token) Pos() Pos {
	return t.pos
}

func (t *token) String() string {
	return fmt.Sprintf("%d [%s] @ %s", t.kind, t.val, t.pos)
}

// Lexer scans the input text for tokens.
type Lexer interface {

	// Peek returns the next lexer token but does not consume it.
	Peek() Token

	// Next returns and consumes the next lexer token.
	Next() Token
}

type lexer struct {
	pos   int
	line  int
	col   int
	val   []rune
	input []rune
	len   int
}

func NewLexer(input string) Lexer {
	rs := []rune(input)
	return &lexer{0, 1, 1, nil, rs, len(rs)}
}

// func NewLexer(input []rune) Lexer {
// 	return &lexer{0, 1, 1, input}
// }

func (l *lexer) Peek() Token {
	//t := l.Next()
	return nil
}

func (l *lexer) Next() Token {
	l.val = []rune("")
	c := l.peek()
	switch c {
	case eof:
		return l.token(EOF)
	case '(', ')':
		l.read()
		return l.token(PAR)
	}
	if isWhitespace(c) {
		l.readWhile(isWhitespace)
		return l.Next()
	}
	if isDec(c) {
		return l.nextNum()
	}
	if isSym(c) {
		return l.nextSymbol()
	}
	return nil
}

func (l *lexer) peek() rune {
	if l.pos >= l.len {
		return eof
	}
	return l.input[l.pos]
}

// func (l *lexer) peeks() []rune {

// }

func (l *lexer) peeks(s string) bool {
	rs := []rune(s)
	for i, r := range rs {
		p := l.pos + i
		if p >= l.len || r != l.input[p] {
			return false
		}
	}
	return true
}

func (l *lexer) read() rune {
	c := l.peek()
	if c != eof {
		if c == '\n' {
			l.line++
			l.col = 1
		} else {
			l.col++
		}
		l.val = append(l.val, c)
		l.pos++
	}
	return c
}

func (l *lexer) readWhile(pred func(rune) bool) {
	c := l.peek()
	for c != eof && pred(c) {
		l.read()
		c = l.peek()
	}
}

func (l *lexer) expect(s string) {
	rs := []rune(s)
	for _, r := range rs {
		l.val = append(l.val, r)
	}
	l.pos += len(s)
}

func (l *lexer) token(kind Kind) Token {
	return NewToken(kind, string(l.val), NewPos(l.line, l.col))
}

func isWhitespace(c rune) bool {
	return strings.ContainsRune(" \t\r", c)
}

func isDec(c rune) bool {
	return strings.ContainsRune("0123456789", c)
}

func isBin(c rune) bool {
	return strings.ContainsRune("01", c)
}

func isHex(c rune) bool {
	return strings.ContainsRune("0123456789ABCDEF", c)
}

func isSym(c rune) bool {
	return (isWhitespace(c) || isDec(c) || c == '(' || c == ')') == false
}

func (l *lexer) nextNum() Token {
	if l.peeks("0b") {
		l.expect("0b")
		l.readWhile(isBin)
	} else if l.peeks("0x") {
		l.expect("0x")
		l.readWhile(isHex)
	} else {
		l.readWhile(isDec)
	}
	return l.token(NUM)
}

func (l *lexer) nextSymbol() Token {
	l.readWhile(isSym)
	return l.token(SYM)
}
