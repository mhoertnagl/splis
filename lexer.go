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
	// STR are strings.
	STR
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
	case '"':
		return l.nextString()
		// case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		// 	return l.nextNum()
	}
	if isWhitespace(c) {
		l.readWhile(isWhitespace)
		return l.Next()
	}
	if l.test("//") {
		l.readWhile(not(isNewline))
		return l.Next()
	}
	if l.test("/*") {
		return l.skipMultiLineComment("/*", "*/")
	}
	if isDec(c) {
		return l.nextNum()
	}
	if isSym(c) {
		return l.nextSymbol()
	}
	return l.err("Unrecognized character [%s].", string(c))
}

func (l *lexer) skipMultiLineComment(start string, end string) Token {
	l.expect(start)
	for l.test(end) == false {
		if l.read() == eof {
			return l.err("Stray multi line comment. Missing [%s].", end)
		}
	}
	l.expect(end)
	return l.Next()
}

// nextNum accepts numbers of the pattern \d+(\.\d*)? .
// The returned token however does not contain the converted numeric value but
// the scanned textual representation of the number.
func (l *lexer) nextNum() Token {
	// if l.test("0b") {
	// 	l.expect("0b")
	// 	l.readWhile(isBin)
	// } else if l.test("0x") {
	// 	l.expect("0x")
	// 	l.readWhile(isHex)
	// } else {
	// 	l.readWhile(isDec)
	// 	if l.test(".") {
	// 		l.read()
	// 		l.readWhile(isDec)
	// 	}
	// }
	l.readWhile(isDec)
	if l.test(".") {
		l.read()
		l.readWhile(isDec)
	}
	return l.token(NUM)
}

// nextString accepts strings that may contain escape sequences. Will return a
// string Token containing the string value without trailing quotes.
func (l *lexer) nextString() Token {
	l.skip()
	for {
		switch l.peek() {
		case eof:
			return l.err("Missing closing quote.")
		case '"':
			l.skip()
			return l.token(STR)
		case '\\':
			l.skip()
			e := l.peek()
			switch e {
			case 'n':
				l.append('\n')
				l.skip()
			case '\\':
				l.read()
			default:
				return l.err("Unrecognized escape sequence [\\%s].", string(e))
			}
			break
		default:
			l.read()
			break
		}
	}
}

// NOTE: Symbols may not contain any decimal digits.
func (l *lexer) nextSymbol() Token {
	l.readWhile(isSym)
	return l.token(SYM)
}

// peek at the next character on the input stream. Will not consume the
// character.
func (l *lexer) peek() rune {
	if l.pos < l.len {
		return l.input[l.pos]
	}
	return eof
}

// append a character to the end of the lexeme.
func (l *lexer) append(c rune) {
	l.val = append(l.val, c)
}

// skip the next character on the input stream. This will however update the
// lexer's positional state.
func (l *lexer) skip() rune {
	c := l.peek()
	if c == eof {
		return eof
	}
	if c == '\n' {
		l.line++
		l.col = 1
	} else {
		l.col++
	}
	l.pos++
	return c
}

// read peeks and consumes a single character from the input stream.
func (l *lexer) read() rune {
	c := l.skip()
	if c == eof {
		return eof
	}
	l.append(c)
	return c
}

// readWhile consumes characters from the input as long as the predicate for
// each character returns true.
func (l *lexer) readWhile(pred func(rune) bool) {
	for c := l.peek(); c != eof && pred(c); c = l.peek() {
		l.read()
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
// an error message and abort further consumption.
func (l *lexer) expect(s string) {
	for _, c := range s {
		r := l.peek()
		if r == rune(c) {
			l.read()
		} else {
			t := string(l.val[len(l.val)-len(s):])
			panic("FATAL: Lexer expected [" + s + "] but got [" + t + "].")
		}
	}
}

// token returns a new Token of a certain type coalescing the lexer's gathered
// scanning state.
func (l *lexer) token(kind Kind) Token {
	return NewToken(kind, string(l.val), NewPos(l.line, l.col))
}

// err returns an error Token with an error message.
func (l *lexer) err(msg string, args ...interface{}) Token {
	return NewToken(ERR, fmt.Sprintf(msg, args...), NewPos(l.line, l.col))
}

// not returns true iff the function f returns false and vice versa.
func not(f func(rune) bool) func(rune) bool {
	return func(c rune) bool { return f(c) == false }
}

// isWhitespace returns true iff the rune is one of [ \t\r\n].
func isWhitespace(c rune) bool {
	return strings.ContainsRune(" \t\r\n", c)
}

// isNewline returns true iff the rune is '\n'.
func isNewline(c rune) bool {
	return c == '\n'
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
		strings.ContainsRune(" \t\r\n0123456789(){}", c) == false
	//return (isWhitespace(c) || isDec(c) || c == '(' || c == ')') == false
}
