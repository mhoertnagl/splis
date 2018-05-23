package main

import (
	"bytes"
	"fmt"
)

func printAst(n Node) string {
	switch v := n.(type) {
	case NumNode:
		return fmt.Sprintf("%d", v.Value())
	case SymNode:
		return v.Name()
	case SExprNode:
		return printSExpr(v)
	}
	return ""
}

func printSExpr(v SExprNode) string {
	var b bytes.Buffer
	b.WriteString("(")
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(printAst(v.Cell(i)))
	}
	b.WriteString(")")
	return b.String()
}
