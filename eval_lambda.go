package main

func (vm *vm) makeLambda(e Env, as []Node) Node {
	t := NewAssertion()
	t.AssertLen(as, 2, "lambda")
	t.AssertType(as[0], QXP_NODE, "First argument of lambda")
	t.AssertType(as[1], QXP_NODE, "Second argument of lambda")

	if t.Failed() {
		return t.Error()
	}

	ps, _ := as[0].(*seqNode)
	body, _ := as[1].(*seqNode)

	ss := []*symNode{}
	for i := 0; i < len(ps.cells); i++ {
		p := ps.cells[i]

		if t.AssertType(p, SYM_NODE, "Lambda parameter") {
			return t.Error()
		}

		s, _ := p.(*symNode)
		ss = append(ss, s)
	}

	return NewLambdaNode(NewSubEnvironment(e), ss, body)
}

func (vm *vm) evalLambda(n *lambdaNode, as []Node) Node {
	plen := len(n.ps)
	alen := len(as)

	// We cannot apply more arguments than there are parameters for the
	// function.
	if plen < alen {
		return NewErrNode("Too many arguments [%v].\n", printAst(n))
	}

	for _, a := range as {
		// Pop the first parameter of the function.
		h := n.Pop()
		// Evaluate the argument and bind the result to the parameter.
		v := vm.eval(n.env, a)
		n.env.Set(h.name, v)
	}

	// There are fewer arguments then function parameters. Return the function
	// and its environment with available arguments bound.
	if plen > alen {
		return n
	}
	// All parameters are bound and the lambda function can be evaluated.
	n.body.typ = SXP_NODE
	return vm.eval(n.env, n.body)
}
