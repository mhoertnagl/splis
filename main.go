package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO: Floating point numbers statt integer?
// TODO: Boolean type.
// TODO: List primitives.
// TODO: Erzeugt die Funktion in einem separaten environment und damit
//       ist die Funktion nicht woanders zugänglich. LÖSUNG: def global,
//       var lokal.
// TODO: do primitive?
// TODO: cond primitive?
// TODO: dedicated node to return success.

func main() {

	reader := bufio.NewReader(os.Stdin)
	vm := NewVM()

	for {
		fmt.Print("splis> ")

		e, _ := reader.ReadString('\n')
		e = strings.TrimSpace(e)

		if e == "exit" {
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

	// if len(os.Args) != 2 {
	// 	panic("Please specify a file.")
	// }

	// f := os.Args[1]
	// buf, err := ioutil.ReadFile(f)

	// if err != nil {
	// 	panic(err)
	// }

	// e := string(buf)
	// l := NewLexer(e)
	// p := NewParser(l)
	// ns := p.Parse()
	// vm := NewVM()
	// for _, n := range ns {
	// 	r := vm.Eval(n)
	// 	s := printAst(r)
	// 	fmt.Println(s)
	// }
}
