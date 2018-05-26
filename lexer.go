package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Kind specifies the type of token returned.
type Kind int

// eof is the End Of File character rune.
const eof = rune(0)

const (
	// PAR are parentehsis tokens.
	PAR Kind = iota
	// CBR are curly braces tokens.
	CBR
	// NUM are numeric tokens.
	NUM
	// SYM is any character that is not a whitespace, a number or a parenthesis.
	SYM
	// EOF is the End Of File token.
	EOF
	// ERR indicates an error token.
	ERR
)

// Pos indicates the line number and character starting position of the token
// within the input string.
type Pos interface {
	Line() int
	Col() int
}

type pos struct {
	line int
	col  int
}

// NewPos returns a new token position.
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

// Token is the basic block of any scanned input.
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

// NewToken return a new token.
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

// Lexer turns a stream of characters into a stream of tokens.
type Lexer interface {

	// Next consumes and returns the next lexer token.
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

// NewLexer creates a new Lexer.
func NewLexer(input string) Lexer {
	rs := []rune(input)
	return &lexer{0, 1, 1, nil, rs, len(rs)}
}

// Next consumes and returns the next lexer token.
func (l *lexer) Next() Token {
	l.val = []rune("")
	c := l.peek()
	switch c {
	case eof:
		return l.token(EOF)
	case '(', ')':
		l.read()
		return l.token(PAR)
	case '{', '}':
		l.read()
		return l.token(CBR)
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

// nextNum accepts numbers of the following patterns:
//  o \d+
//  o 0b[01]+
//  o 0x[0-9A-F]+
// The returned token however does not contain the converted numeric value but
// the scanned textual representation of the number.
func (l *lexer) nextNum() Token {
	if l.test("0b") {
		l.expect("0b")
		l.readWhile(isBin)
	} else if l.test("0x") {
		l.expect("0x")
		l.readWhile(isHex)
	} else {
		l.readWhile(isDec)
	}
	return l.token(NUM)
}

// NOTE: Symbols may not contain any decimal digits.
func (l *lexer) nextSymbol() Token {
	l.readWhile(isSym)
	return l.token(SYM)
}

// peek at the next character in the input stream. Will not consume the
// character.
func (l *lexer) peek() rune {
	if l.pos < l.len {
		return l.input[l.pos]
	}
	return eof
}

// read peeks and consumes a single character from the input.
func (l *lexer) read() rune {
	c := l.peek()
	if c == eof {
		return eof
	} else if c == '\n' {
		l.line++
		l.col = 1
	} else {
		l.col++
	}
	l.val = append(l.val, c)
	l.pos++
	return c
}

// readWhile consumes characters from the input as long as the predicate for
// each character returns true.
func (l *lexer) readWhile(pred func(rune) bool) {
	c := l.peek()
	for c != eof && pred(c) {
		l.read()
		c = l.peek()
	}
}

// test compares the next characters with the string s and returns true iff
// the input matches s. It does not consume any characters in this process.
func (l *lexer) test(s string) bool {
	for i, c := range s {
		p := l.pos + i
		if p >= l.len || rune(c) != l.input[p] {
			return false
		}
	}
	return true
}

// expect compares the next characters with the string s and will consume any
// matching character up until the first mismatch if any. A mismatch will prompt
// an error message. Subsequent potentially matching characters will not be
// consumed.
func (l *lexer) expect(s string) {
	for _, c := range s {
		r := l.peek()
		if r == rune(c) {
			l.read()
		} else {
			// Print Error.
			return
		}
	}
}

// token returns a new Token of a certain type coalescing the lexer's gathered
// scanning state.
func (l *lexer) token(kind Kind) Token {
	//return NewToken(kind, string(l.val), NewPos(l.line, l.col))
	t := NewToken(kind, string(l.val), NewPos(l.line, l.col))
	//fmt.Printf("%s\n", t)
	return t
}

// isWhitespace returns true iff the rune is one of [ \t\r\n].
func isWhitespace(c rune) bool {
	return strings.ContainsRune(" \t\r\n", c)
}

// isDec returns true iff the rune is a decimal digit.
func isDec(c rune) bool {
	return strings.ContainsRune("0123456789", c)
}

// isDec returns true iff the rune is either '0' or '1'.
func isBin(c rune) bool {
	return strings.ContainsRune("01", c)
}

// isHex returns true iff the rune is a hexadecimal digit. Note however, that
// the lower-case hexadecimal digits [a-f] are not supported.
func isHex(c rune) bool {
	return strings.ContainsRune("0123456789ABCDEF", c)
}

// isSym returns true iff the rune is printable as defined in unicode.IsPrint
// and is not a whitespace character a decimal digit, '(' or ')'.
func isSym(c rune) bool {
	return unicode.IsPrint(c) &&
		strings.ContainsRune(" \t\r\n0123456789()", c) == false
	//return (isWhitespace(c) || isDec(c) || c == '(' || c == ')') == false
}
