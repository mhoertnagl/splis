package main

import (
	"bytes"
	"fmt"
)

func printAst(n Node) string {
	switch v := n.(type) {
	case *errNode:
		return v.msg
	case *numNode:
		return fmt.Sprintf("%d", v.val)
	case *symNode:
		return v.name
	case *sExprNode:
		return printSeq(v, "(", printAst, ")")
	case *qExprNode:
		return printSeq(v, "{", printAst, "}")
	// case *funNode:
	// 	return fmt.Sprintf("<fun>")
	case *lambdaNode:
		return fmt.Sprintf("(lambda %s) %s", printAst(v.body), printEnv(v.env))
	}
	return ""
}

func printSeq(v SeqNode, l string, printElements func(Node) string, r string) string {
	var b bytes.Buffer
	b.WriteString(l)
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(printElements(v.Cell(i)))
	}
	b.WriteString(r)
	return b.String()
}

func printEnv(e Env) string {
	var b bytes.Buffer
	b.WriteString("[")
	for k, v := range e.All() {
		b.WriteString(", ")
		b.WriteString(k)
		b.WriteString("=")
		b.WriteString(printAst(v))
	}
	b.WriteString("]")
	return b.String()
}
