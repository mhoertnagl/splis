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
	e.SetNum("true", 1)
	e.SetNum("false", 0)
	e.SetFun("+", vm.evalAdd)
	e.SetFun("<", vm.evalLT)
	e.SetFun("def", vm.defVar)
	e.SetFun("eval", vm.evalNodes)
	e.SetFun("lambda", vm.makeLambda)
	return vm
}

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

	as := []Node{}
	for i := 1; i < n.Len(); i++ {
		a := vm.eval(e, n.Cell(i))
		as = append(as, a)
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
		// Pop the first parameter of the function.
		h := n.Pop()
		// fmt.Printf("Lambda Param: %s\n", printAst(h))
		// Evaluate the argument and bind the result to the parameter.
		v := vm.eval(n.env, a)
		n.env.Set(h.name, v)
	}

	// There are fewer arguments then function parameters. Return the function
	// and its environment with available arguments bound.
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
		case *numNode:
			sum += v.val
			break
		default:
			fmt.Printf("Cannot add non-number [%v].\n", v)
			break
		}
	}
	return &numNode{sum}
}

func (vm *vm) evalLT(e Env, as []Node) Node {
	t := NewAssertion()
	t.AssertLen(as, 2, "<")
	t.AssertType(as[0], NUM_NODE, "First argument of <")
	t.AssertType(as[1], NUM_NODE, "Second argument of <")

	if t.Failed() {
		return t.Error()
	}

	x, _ := as[0].(*numNode)
	y, _ := as[1].(*numNode)

	if x.val < y.val {
		return &numNode{1}
	}
	return &numNode{0}
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

func (vm *vm) makeLambda(e Env, as []Node) Node {
	t := NewAssertion()
	t.AssertLen(as, 2, "lambda")
	t.AssertType(as[0], QXP_NODE, "First argument of lambda")
	t.AssertType(as[1], QXP_NODE, "Second argument of lambda")

	if t.Failed() {
		return t.Error()
	}

	ps, _ := as[0].(*qExprNode)
	body, _ := as[1].(*qExprNode)

	ss := []*symNode{}
	for i := 0; i < ps.Len(); i++ {
		p := ps.Cell(i)

		if t.AssertType(p, SYM_NODE, "Lambda parameter") {
			return t.Error()
		}

		s, _ := ps.Cell(i).(*symNode)
		ss = append(ss, s)
	}

	return NewLambdaNode(NewSubEnvironment(e), ss, body)
}

func (vm *vm) defVar(e Env, as []Node) Node {
	t := NewAssertion()
	t.AssertMin(as, 2, "Def")
	t.AssertType(as[0], QXP_NODE, "First argument of def")

	if t.Failed() {
		return t.Error()
	}

	ps, _ := as[0].(*qExprNode)

	if t.AssertIntEqual(ps.Len(), len(as)-1, "Number of names", "definitions") {
		return t.Error()
	}

	for i := 0; i < ps.Len(); i++ {
		p := ps.Cell(i)

		if t.AssertType(p, SYM_NODE, "Def parameter") {
			return t.Error()
		}

		s, _ := ps.Cell(i).(*symNode)
		v := vm.eval(e, as[i+1])
		e.Set(s.name, v)
	}

	return NewSExprNode()
}
