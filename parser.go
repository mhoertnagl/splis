package main

type Parser interface {
	Parse() []Node
}

type parser struct {
	lex Lexer
	cur Token
}

func NewParser(lex Lexer) Parser {
	return &parser{lex, nil}
}

func (p *parser) Parse() []Node {
	ns := []Node{}
	for p.next(); p.cur.Kind() != EOF; p.next() {
		n := p.expr()
		printAst(n)
		ns = append(ns, n)
	}
	return ns
}

func (p *parser) next() {
	p.cur = p.lex.Next()
}

func (p *parser) nextIsNot(val string) bool {
	p.next()
	return p.cur.Kind() != EOF && p.cur.Value() != val
}

func (p *parser) expect(val string) bool {
	return p.cur.Value() == val
}

func (p *parser) seq(l string, expr func() Node, r string, n *seqNode) Node {
	if p.expect(r) {
		return NewErrNode("Stray closing [%s]. Missing corresponding open [%s].", r, l)
	}
	p.expect(l)
	for p.nextIsNot(r) {
		n.Push(expr())
	}
	if p.expect(r) {
		return n
	}
	return NewErrNode("Stray open [%s]. Missing corresponding closing [%s].", l, r)
}

func (p *parser) expr() Node {
	switch p.cur.Kind() {
	case EOF:
		return NewErrNode("Unexpected end of file.")
	case ERR:
		return NewErrNode(p.cur.Value())
	case NUM:
		return NewNumNodeFromString(p.cur.Value())
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
