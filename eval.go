package main

import "fmt"

type VM interface {
	Eval(n Node) Node
}

type vm struct {
	builtins map[string]func(args []Node) Node
}

func NewVM() VM {
	m := map[string]func(args []Node) Node{}
	vm := &vm{m}
	vm.builtins["+"] = evalAdd
	return vm
}

func (vm *vm) Eval(n Node) Node {
	switch v := n.(type) {
	// case NumNode:
	// 	return v
	// case SymNode:
	// 	return v
	case *sExprNode:
		return vm.evalSExpr(v)
	}
	// Returns Numbers, Symbols and Q-Expressions as-is.
	return n
}

func (vm *vm) evalSExpr(n *sExprNode) Node {
	if n.Len() == 0 {
		return n
	}

	f := n.Cell(0)
	f1, ok := f.(SymNode)
	if !ok {
		fmt.Printf("First element of a S-Expr must be a Symbol but is [%s].", f1)
		return n
	}

	args := []Node{}
	for i := 1; i < n.Len(); i++ {
		args = append(args, vm.Eval(n.Cell(i)))
	}

	b := vm.builtins[f1.Name()]
	if b == nil {
		fmt.Printf("Unupported builtin symbol [%s].", f1.Name())
		return n
	}
	return b(args)
}

func evalAdd(args []Node) Node {
	var sum int32
	for _, arg := range args {
		switch v := arg.(type) {
		case NumNode:
			sum += v.Value()
			break
		default:
			fmt.Printf("Cannot add the non-number [%s]", v)
			break
		}
	}
	return &numNode{sum}
}
