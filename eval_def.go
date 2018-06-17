package main

func (vm *vm) defVar(e Env, as []Node) Node {
	t := NewAssertion()
	t.AssertMin(as, 2, "Def")
	t.AssertType(as[0], QXP_NODE, "First argument of def")

	if t.Failed() {
		return t.Error()
	}

	ps, _ := as[0].(*seqNode)

	if t.AssertIntEqual(len(ps.cells), len(as)-1, "Number of names", "definitions") {
		return t.Error()
	}

	for i := 0; i < len(ps.cells); i++ {
		p := ps.cells[i]

		if t.AssertType(p, SYM_NODE, "Def parameter") {
			return t.Error()
		}

		s, _ := p.(*symNode)
		v := vm.eval(e, as[i+1])
		//e.Set(s.name, v)
		vm.env.Set(s.name, v)
	}

	return NewSExprNode()
}
