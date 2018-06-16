package main

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

func (vm *vm) evalEQ(e Env, as []Node) Node {
	t := NewAssertion()
	t.AssertLen(as, 2, "==")

	if t.Failed() {
		return t.Error()
	}

	if vm.evalEQ2(e, as[0], as[1]) {
		return &numNode{1}
	}
	return &numNode{0}
}

func (vm *vm) evalNE(e Env, as []Node) Node {
	t := NewAssertion()
	t.AssertLen(as, 2, "!=")

	if t.Failed() {
		return t.Error()
	}

	if vm.evalEQ2(e, as[0], as[1]) {
		return &numNode{0}
	}
	return &numNode{1}
}

func (vm *vm) evalEQ2(e Env, x Node, y Node) bool {
	if x.Type() != y.Type() {
		return false
	}
	switch x.Type() {
	case ERR_NODE:
		return vm.evalErrEQ(e, x.(*errNode), y.(*errNode))
	case NUM_NODE:
		return vm.evalNumEQ(e, x.(*numNode), y.(*numNode))
	case SYM_NODE:
		return vm.evalSymEQ(e, x.(*symNode), y.(*symNode))
	case SXP_NODE:
		return vm.evalSeqEQ(e, x.(*seqNode), y.(*seqNode))
	case QXP_NODE:
		return vm.evalSeqEQ(e, x.(*seqNode), y.(*seqNode))
	case FUN_NODE:
		return vm.evalFunEQ(e, x.(*funNode), y.(*funNode))
	case LBD_NODE:
		return vm.evalLbdEQ(e, x.(*lambdaNode), y.(*lambdaNode))
	}
	return false
}

func (vm *vm) evalErrEQ(e Env, x *errNode, y *errNode) bool {
	return x.msg == y.msg
}

func (vm *vm) evalNumEQ(e Env, x *numNode, y *numNode) bool {
	return x.val == y.val
}

func (vm *vm) evalSymEQ(e Env, x *symNode, y *symNode) bool {
	return x.name == y.name
}

func (vm *vm) evalSeqEQ(e Env, x *seqNode, y *seqNode) bool {
	if len(x.cells) != len(y.cells) {
		return false
	}
	for i := 0; i < len(x.cells); i++ {
		if vm.evalEQ2(e, x.cells[i], y.cells[i]) == false {
			return false
		}
	}
	return true
}

func (vm *vm) evalFunEQ(e Env, x *funNode, y *funNode) bool {
	return &x.fun == &y.fun
}

func (vm *vm) evalLbdEQ(e Env, x *lambdaNode, y *lambdaNode) bool {
	if len(x.ps) != len(y.ps) {
		return false
	}
	for i := 0; i < len(x.ps); i++ {
		if x.ps[i].name != y.ps[i].name {
			return false
		}
	}
	return vm.evalSeqEQ(e, x.body, y.body)
}
