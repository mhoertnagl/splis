package main

import "fmt"

type Parser interface {
	Parse()
}

type parser struct {
	lex Lexer
	cur Token
}

func NewParser(lex Lexer) Parser {
	return &parser{lex, nil}
}

func (p *parser) Parse() {
	p.next()
	p.expr()
	fmt.Println()
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

func (p *parser) expr() {
	switch p.cur.Kind() {
	case EOF:
		return
	case ERR:
		panic("Lexer Error")
	case NUM:
		fmt.Printf("%s ", p.cur.Value())
		// Return number node
		break
	case SYM:
		fmt.Printf("%s ", p.cur.Value())
		// Return symbol node
		break
	case PAR:
		p.subExpr()
		break
	}
}

func (p *parser) subExpr() {
	p.expect("(")
	p.next()
	for p.cur.Value() != ")" {
		p.expr()
		p.next()
	}
	p.expect(")")
}
