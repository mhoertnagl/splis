package main

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
