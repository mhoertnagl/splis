package main

import "testing"

func TestEvalEnv(t *testing.T) {
	vm := setupVM(t, "(def {a} 100)", "()")
	assertEnvEqual(t, vm, "a", "100")
}

func TestEvalEnv2(t *testing.T) {
	vm := setupVM(t, "(def {a b} 50 200)", "()")
	assertEnvEqual(t, vm, "a", "50")
	assertEnvEqual(t, vm, "b", "200")
}

func TestEvalEnv3(t *testing.T) {
	setupVM(t, "(def {a} 50 200)", "Number of names [1] and definitions [2] must be the equal.\n")
}

func TestEvalEnv4(t *testing.T) {
	setupVM(t, "(def {a b} 50)", "Number of names [2] and definitions [1] must be the equal.\n")
}

func TestEvalEnv5(t *testing.T) {
	setupVM(t, "(def {5 b} 50 50)", "Def parameter must be of type [Symbol] but is [Number].\n")
}

func TestEvalEnv6(t *testing.T) {
	setupVM(t, "(def {a})", "Def requires at least [2] arguments.\n")
}

// func TestEvalEnv6(t *testing.T) {
// 	setupVM(t, "(def (a) 50)", "Number of defined names [2] and definitions [1] must be the equal.\n")
// }

// func TestEvalSubEnv(t *testing.T) {
// 	vm := setupVM(t, "(def {a b} 50 200)")
// 	assertEnvEqual(t, vm, "a", "50")
// 	assertEnvEqual(t, vm, "b", "200")
// }

func setupVM(t *testing.T, s string, q string) VM {
	l := NewLexer(s)
	p := NewParser(l)
	n := p.Parse()

	vm := NewVM()
	res := vm.Eval(n)
	r := printAst(res)
	if r != q {
		t.Errorf("Expected [%v] but got [%v]", q, r)
	}
	return vm
}

func assertEnvEqual(t *testing.T, vm VM, s string, e string) {
	l := NewLexer(s)
	p := NewParser(l)
	n := p.Parse()
	res := vm.Eval(n)
	r := printAst(res)
	if r != e {
		t.Errorf("Expected [%v] but got [%v]", e, r)
	}
}
