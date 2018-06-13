package main

func (vm *vm) evalAdd(e Env, as []Node) Node {
	var sum int32
	for _, a := range as {
		switch v := a.(type) {
		case *numNode:
			sum += v.val
			break
		default:
			return NewErrNode("Cannot add non-number [%s].\n", printAst(v))
		}
	}
	return &numNode{sum}
}
