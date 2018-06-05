package main

import "fmt"

type VM interface {
	Eval(n Node) Node
}

type vm struct {
	env Env
}

func NewVM() VM {
	e := NewEnvironment()
	vm := &vm{e}
	e.Set("true", NewNumNode("1"))
	e.Set("false", NewNumNode("0"))
	e.Set("+", NewFunNode(vm.evalAdd))
	e.Set("def", NewFunNode(vm.defVar))
	e.Set("eval", NewFunNode(vm.evalNodes))
	e.Set("lambda", NewFunNode(vm.makeLambda))
	return vm
}

func (vm *vm) Eval(n Node) Node {
	return vm.eval(vm.env, n)
}

func (vm *vm) eval(e Env, n Node) Node {
	switch v := n.(type) {
	case SymNode:
		return e.Get(v.Name()) // Call Eval?
	case *sExprNode:
		return vm.evalSeq(e, v) //vm.evalSExpr(e, v)
	}
	// Return Numbers and Q-Expressions as-is.
	return n
}

func (vm *vm) evalSeq(e Env, n SeqNode) Node {
	if n.Len() == 0 {
		return n
	}

	as := []Node{}
	for i := 1; i < n.Len(); i++ {
		as = append(as, vm.eval(e, n.Cell(i)))
	}

	fn := vm.eval(e, n.Cell(0))
	switch f := fn.(type) {
	case *funNode:
		return f.fun(e, as)
	case *lambdaNode:
		return vm.evalLambda(f, as)
	default:
		//fmt.Printf("First element of an S-Expression must be a Function-Symbol but is [%v].\n", f)
		return f
	}
}

func (vm *vm) evalLambda(n *lambdaNode, as []Node) Node {
	// fmt.Printf("Lambda: %s\n", printAst(n))
	plen := len(n.ps)
	alen := len(as)
	if plen < alen {
		fmt.Printf("Too many arguments [%v].\n", n)
		return NewSExprNode()
	}

	// fmt.Printf("# Params [%d]\n", len(as))
	for _, a := range as {
		h := n.Pop()
		// fmt.Printf("Lambda Param: %s\n", printAst(h))
		n.env.Set(h.Name(), vm.eval(n.env, a))
	}

	if plen > alen {
		// fmt.Printf("Partial Lambda: %s\n", printAst(n))
		return n
	}
	// fmt.Printf("Eval: %s\n", printAst(n))
	return vm.evalNode(n.env, n.body)
}

func (vm *vm) evalAdd(e Env, as []Node) Node {
	var sum int32
	for _, a := range as {
		switch v := a.(type) {
		case NumNode:
			sum += v.Value()
			break
		default:
			fmt.Printf("Cannot add non-number [%v].\n", v)
			break
		}
	}
	return &numNode{sum}
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
	case SymNode:
		return e.Get(v.Name()) // Call Eval?
	case *sExprNode:
		return vm.evalSeq(e, v)
	case *qExprNode:
		return vm.evalSeq(e, v)
	default:
		return v
	}
}

func (vm *vm) makeLambda(e Env, as []Node) Node {
	len := len(as)
	if len != 2 {
		fmt.Printf("Lambda requires exactly 2 arguments.\n")
		return NewSExprNode()
	}

	ps, pok := as[0].(*qExprNode)
	if !pok {
		fmt.Printf("Lambda requires the first argument to be a Q-Expression.\n")
		return NewSExprNode()
	}

	ss := []SymNode{}
	for i := 0; i < ps.Len(); i++ {
		p, sok := ps.Cell(i).(SymNode)
		if !sok {
			fmt.Printf("Parameter [%d] is not a Symbol.\n", i+1)
			return NewSExprNode()
		}
		ss = append(ss, p)
	}

	body, bok := as[1].(*qExprNode)
	if !bok {
		fmt.Printf("Lambda requires the second argument to be a Q-Expression.\n")
		return NewSExprNode()
	}

	//fmt.Printf("Make Lambda\n")
	return NewLambdaNode(NewSubEnvironment(e), ss, body)
}

func (vm *vm) defVar(e Env, as []Node) Node {
	alen := len(as) - 1
	if alen < 1 {
		fmt.Printf("Define requires at least 2 arguments.\n")
		return NewSExprNode()
	}

	ps, ok := as[0].(*qExprNode)
	if !ok {
		fmt.Printf("Define requires the first argument to be a Q-Expression.\n")
		return NewSExprNode()
	}

	plen := ps.Len()
	if plen != alen {
		fmt.Printf("Number of defined names [%d] and definitions [%d] must be the same.\n", plen, alen)
		return NewSExprNode()
	}

	for i := 0; i < ps.Len(); i++ {
		p, ok := ps.Cell(i).(SymNode)
		if !ok {
			fmt.Printf("Parameter [%d] is not a Symbol.\n", i+1)
			return NewSExprNode()
		}
		e.Set(p.Name(), vm.eval(e, as[i+1]))
	}

	return NewSExprNode()
}
