package main

func (vm *vm) evalAnd(e Env, as []Node) Node {
	t := NewAssertion()

	for i := 0; i < len(as); i++ {
		if t.AssertType(as[i], NUM_NODE, "Argument of &&") {
			return t.Error()
		}
		if as[i].(*numNode).val == 0 {
			return &numNode{0}
		}
	}
	return &numNode{1}
}

func (vm *vm) evalOr(e Env, as []Node) Node {
	t := NewAssertion()

	for i := 0; i < len(as); i++ {
		if t.AssertType(as[i], NUM_NODE, "Argument of ||") {
			return t.Error()
		}
		if as[i].(*numNode).val != 0 {
			return &numNode{1}
		}
	}
	return &numNode{0}
}

func (vm *vm) evalNot(e Env, as []Node) Node {
	t := NewAssertion()
	t.AssertLen(as, 1, "!")

	if t.AssertType(as[0], NUM_NODE, "Argument of !") {
		return t.Error()
	}

	if as[0].(*numNode).val == 0 {
		return &numNode{1}
	}
	return &numNode{0}
}
