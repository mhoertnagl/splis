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
	e.SetFun("eval", vm.evalEval)
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
	case *seqNode:
		if v.typ == SXP_NODE {
			return vm.evalSeq(e, v)
		}
	}
	// Return Numbers, Strings and Q-Expressions as-is.
	return n
}

func (vm *vm) evalSeq(e Env, n *seqNode) Node {
	if len(n.cells) == 0 {
		return n
	}

	// Evaluate all function arguments.
	as := []Node{}
	for i := 1; i < len(n.cells); i++ {
		a := vm.eval(e, n.cells[i])
		as = append(as, a)
	}

	// Evaluate the sequence head. If it evaluates to a function or lambda,
	// apply the arguments.
	fn := vm.eval(e, n.cells[0])
	switch f := fn.(type) {
	case *funNode:
		return f.fun(e, as)
	case *lambdaNode:
		return vm.evalLambda(f, as)
	default:
		return f
	}
}

func (vm *vm) evalEval(e Env, ns []Node) Node {
	var r Node = NewSExprNode()
	for _, n := range ns {
		if n.Type() == QXP_NODE {
			q, _ := n.(*seqNode)
			q.typ = SXP_NODE
		}
		r = vm.eval(e, n)
	}
	return r
}
