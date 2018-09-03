package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO: do primitive?
// TODO: cond primitive?
// TODO: | cs list processing
// TODO: let primitive
// TODO: Assertion object has to be global. and provide a stak trace.
// TODO: Boolean type.
// TODO: dedicated node to return success.

func main() {

	rd := bufio.NewReader(os.Stdin)
	vm := NewVM()
	vm.Load("lib/prelude")

	args := os.Args[1:]

	// Load additional files.
	if len(args) > 0 {
		for _, arg := range args {
			vm.Load(arg)
		}
	}

	for {
		fmt.Print("splis> ")

		e, _ := rd.ReadString('\n')
		e = strings.TrimSpace(e)

		if e == ":exit" {
			break
		}

		l := NewLexer(e)
		p := NewParser(l)
		ns := p.Parse()

		for _, n := range ns {
			r := vm.Eval(n)
			s := printAst(r)
			fmt.Println(s)
		}
	}
}
