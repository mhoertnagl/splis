package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO: Assertion object has to be global. and provide a stak trace.
// TODO: Floating point numbers statt integer?
// TODO: Boolean type.
// TODO: do primitive?
// TODO: cond primitive?
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
