package main

func (vm *vm) evalIf(e Env, as []Node) Node {
	t := NewAssertion()
	t.AssertLen(as, 3, "If")
	t.AssertType(as[0], NUM_NODE, "First argument of if")
	t.AssertType(as[1], QXP_NODE, "Second argument of if")
	t.AssertType(as[2], QXP_NODE, "Third argument of if")

	if t.Failed() {
		return t.Error()
	}

	if as[0].(*numNode).val != 0 {
		return vm.evalNode(e, as[1])
	}
	return vm.evalNode(e, as[2])
}
