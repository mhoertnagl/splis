package main

type VM interface {
	Eval(n Node) Node
}

type vm struct {
	env Env
}

func NewVM() VM {
	e := NewEnvironment()
	vm := &vm{e}
	e.SetNum("true", 1) // Ersetzen durch prelude.splis > (def {true} 1) (def {false} 0)
	e.SetNum("false", 0)
	e.SetFun("+", vm.evalAdd)
	e.SetFun("<", vm.evalLT)
	e.SetFun("==", vm.evalEQ)
	e.SetFun("!=", vm.evalNE)
	e.SetFun("def", vm.defVar)
	e.SetFun("eval", vm.evalNodes)
	e.SetFun("lambda", vm.makeLambda)
	e.SetFun("if", vm.evalIf)
	e.SetFun("&&", vm.evalAnd)
	return vm
}

// TODO: Bis auf diese Funktion muss nichts an VM hängen.

func (vm *vm) Eval(n Node) Node {
	return vm.eval(vm.env, n)
}

func (vm *vm) eval(e Env, n Node) Node {
	switch v := n.(type) {
	case *symNode:
		return e.Get(v.name)
	case *sExprNode:
		return vm.evalSeq(e, v)
	}
	// Return Numbers and Q-Expressions as-is.
	return n
}

func (vm *vm) evalSeq(e Env, n SeqNode) Node {
	if n.Len() == 0 {
		return n
	}

	// Evaluate all function arguments.
	as := []Node{}
	for i := 1; i < n.Len(); i++ {
		a := vm.eval(e, n.Cell(i))
		as = append(as, a)
	}

	// Evaluate the sequence head. If it evaluates to a function or lambda, apply.
	fn := vm.eval(e, n.Cell(0))
	switch f := fn.(type) {
	case *funNode:
		return f.fun(e, as)
	case *lambdaNode:
		return vm.evalLambda(f, as)
	default:
		return f
	}
}

func (vm *vm) evalNodes(e Env, ns []Node) Node {
	var r Node = NewSExprNode()
	for _, n := range ns {
		r = vm.evalNode(e, n)
	}
	return r
}

func (vm *vm) evalNode(e Env, n Node) Node {
	switch v := n.(type) {
	case *symNode:
		return e.Get(v.name)
	case *sExprNode:
		return vm.evalSeq(e, v)
	case *qExprNode:
		return vm.evalSeq(e, v)
	default:
		return v
	}
}
