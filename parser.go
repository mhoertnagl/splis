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
	//fmt.Println()
}

func (p *parser) next() {
	p.cur = p.lex.Next()
}

func (p *parser) expect(val string) {
	//p.next()
	if p.cur.Value() != val {
		fmt.Printf("Expecting [" + val + "] but got [" + p.cur.Value() + "].\n")
	} else {
		fmt.Printf("%s ", val)
	}
}

func (p *parser) expr() Node {
	switch p.cur.Kind() {
	case EOF:
		return nil
	case ERR:
		panic("Lexer Error")
	case NUM:
		fmt.Printf("%s ", p.cur.Value())
		// Return number node
		//break
		return NewNumNode(p.cur.Value())
	case SYM:
		fmt.Printf("%s ", p.cur.Value())
		// Return symbol node
		//break
		return NewSymNode(p.cur.Value())
	case PAR:
		return p.subExpr()
	}
	return nil
}

func (p *parser) subExpr() Node {
	n := NewSExprNode()
	p.expect("(")
	p.next()
	for p.cur.Value() != ")" {
		n.Push(p.expr())
		p.next()
	}
	p.expect(")")
	return n
}
