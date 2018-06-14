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
