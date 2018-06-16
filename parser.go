package main

import "fmt"

type Parser interface {
	Parse() Node
}

type parser struct {
	lex Lexer
	cur Token
}

func NewParser(lex Lexer) Parser {
	return &parser{lex, nil}
}

func (p *parser) Parse() Node {
	p.next()
	return p.expr()
}

func (p *parser) next() {
	p.cur = p.lex.Next()
}

func (p *parser) nextIsNot(val string) bool {
	p.next()
	v := p.cur.Value()
	if p.cur.Kind() == EOF {
		fmt.Printf("Unexprected end of file. Expecting [%s].", val)
		panic("Unexprected end of file. Expecting [" + val + "].")
	}
	return v != val
}

func (p *parser) expect(val string) {
	v := p.cur.Value()
	if v != val {
		fmt.Printf("Expecting [%s] but got [%s].\n", val, v)
		panic("Expecting [" + val + "] but got [" + v + "].")
	}
}

func (p *parser) seq(l string, expr func() Node, r string, n SeqNode) Node {
	p.expect(l)
	for p.nextIsNot(r) {
		n.Push(expr())
	}
	p.expect(r)
	return n
}

func (p *parser) expr() Node {
	switch p.cur.Kind() {
	case EOF:
		return NewErrNode("Unexpected end of file.")
	case ERR:
		return NewErrNode(p.cur.Value())
	case NUM:
		return NewNumNode(p.cur.Value())
	case STR:
		return NewStrNode(p.cur.Value())
	case SYM:
		return NewSymNode(p.cur.Value())
	case PAR:
		return p.seq("(", p.expr, ")", NewSExprNode())
	case CBR:
		return p.seq("{", p.expr, "}", NewQExprNode())
	}
	return NewErrNode("Unrecognized token [%s].", p.cur)
}
