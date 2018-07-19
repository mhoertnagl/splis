package main

func (vm *vm) evalAdd(e Env, as []Node) Node {
	t := NewAssertion()
	var sum int32
	for i := 0; i < len(as); i++ {
		if t.AssertType(as[i], NUM_NODE, "Argument of +") {
			return t.Error()
		}
		v := as[i].(*numNode)
		sum += v.val
	}
	return NewNumNode(sum)
}

func (vm *vm) evalSub(e Env, as []Node) Node {
	t := NewAssertion()
	var diff int32
	if len(as) == 1 {
		if t.AssertType(as[0], NUM_NODE, "Argument of -") {
			return t.Error()
		}
		v := as[0].(*numNode)
		return NewNumNode(-v.val)
	}
	if len(as) > 1 {
		if t.AssertType(as[0], NUM_NODE, "Argument of -") {
			return t.Error()
		}
		v := as[0].(*numNode)
		diff = v.val
	}
	for i := 1; i < len(as); i++ {
		if t.AssertType(as[i], NUM_NODE, "Argument of -") {
			return t.Error()
		}
		v := as[i].(*numNode)
		diff -= v.val
	}
	return NewNumNode(diff)
}

func (vm *vm) evalMul(e Env, as []Node) Node {
	t := NewAssertion()
	var prod int32 = 1
	for i := 0; i < len(as); i++ {
		if t.AssertType(as[i], NUM_NODE, "Argument of *") {
			return t.Error()
		}
		v := as[i].(*numNode)
		prod *= v.val
	}
	return NewNumNode(prod)
}

func (vm *vm) evalDiv(e Env, as []Node) Node {
	t := NewAssertion()
	var div int32 = 1
	if len(as) == 1 {
		if t.AssertType(as[0], NUM_NODE, "Argument of /") {
			return t.Error()
		}
		v := as[0].(*numNode)
		if t.AssertDivisble(v.val) {
			return t.Error()
		}
		return NewNumNode(1 / v.val)
	}
	if len(as) > 1 {
		if t.AssertType(as[0], NUM_NODE, "Argument of /") {
			return t.Error()
		}
		v := as[0].(*numNode)
		if t.AssertDivisble(v.val) {
			return t.Error()
		}
		div = v.val
	}
	for i := 1; i < len(as); i++ {
		if t.AssertType(as[i], NUM_NODE, "Argument of /") {
			return t.Error()
		}
		v := as[i].(*numNode)
		if t.AssertDivisble(v.val) {
			return t.Error()
		}
		div /= v.val
	}
	return NewNumNode(div)
}
