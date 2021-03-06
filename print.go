package main

import (
	"bytes"
	"fmt"
)

func printAst(n Node) string {
	switch v := n.(type) {
	case *errNode:
		return fmt.Sprintf("Error: %s\n", v.msg)
	case *numNode:
		return fmt.Sprintf("%g", v.val)
	case *strNode:
		return fmt.Sprintf("\"%s\"", v.str)
	case *symNode:
		return v.name
	case *seqNode:
		if v.typ == SXP_NODE {
			return printSeq(v, "(", printAst, ")")
		}
		return printSeq(v, "{", printAst, "}")
	// case *funNode:
	// 	return fmt.Sprintf("<fun>")
	case *lambdaNode:
		return fmt.Sprintf("(lambda %s %s)", printArgs(v.ps), printAst(v.body))
	}
	return ""
}

func printSeq(v *seqNode, l string, printElements func(Node) string, r string) string {
	var b bytes.Buffer
	b.WriteString(l)
	for i := 0; i < len(v.cells); i++ {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(printElements(v.cells[i]))
	}
	b.WriteString(r)
	return b.String()
}

func printArgs(as []*symNode) string {
	var b bytes.Buffer
	b.WriteString("{")
	for i := 0; i < len(as); i++ {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(as[i].name)
	}
	b.WriteString("}")
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
