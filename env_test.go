package main

import "testing"

func TestEvalEnv(t *testing.T) {
	vm := setupVM(t, "(def {a} 100)")
	assertEnvEqual(t, vm, "a", "100")
}

func TestEvalEnv2(t *testing.T) {
	vm := setupVM(t, "(def {a b} 50 200)")
	assertEnvEqual(t, vm, "a", "50")
	assertEnvEqual(t, vm, "b", "200")
}

func setupVM(t *testing.T, s string) VM {
	l := NewLexer(s)
	p := NewParser(l)
	n := p.Parse()

	vm := NewVM()
	res := vm.Eval(n)
	r := printAst(res)
	if r != "()" {
		t.Errorf("Expected [%v] but got [%v]", "()", r)
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
