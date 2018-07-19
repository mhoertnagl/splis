package main

type VM interface {
	Eval(n Node) Node
	Load(f string)
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
	e.SetFun("-", vm.evalSub)
	e.SetFun("*", vm.evalMul)
	e.SetFun("/", vm.evalDiv)
	e.SetFun("<", vm.evalLT)
	e.SetFun("==", vm.evalEQ)
	e.SetFun("!=", vm.evalNE)
	e.SetFun("def", vm.evalDef)
	e.SetFun("var", vm.evalVar)
	e.SetFun("eval", vm.evalEval)
	e.SetFun("list", vm.evalList)
	e.SetFun("join", vm.evalJoin)
	e.SetFun("head", vm.evalHead)
	e.SetFun("tail", vm.evalTail)
	e.SetFun("lambda", vm.makeLambda)
	e.SetFun("if", vm.evalIf)
	e.SetFun("&&", vm.evalAnd)
	e.SetFun("||", vm.evalOr)
	e.SetFun("!", vm.evalNot)
	e.SetFun("load", vm.evalLoad)
	e.SetFun("execute", vm.evalExecute)
	e.SetFun("print", vm.evalPrint)
	e.SetFun("error", vm.evalError)
	return vm
}

func (vm *vm) Eval(n Node) Node {
	return vm.eval(vm.env, n)
}

func (vm *vm) Load(f string) {
	fs := []Node{NewStrNode(f)}
	cs := []Node{vm.evalLoad(vm.env, fs)}
	vm.evalExecute(vm.env, cs)
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
