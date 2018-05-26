package main

import "fmt"

type VM interface {
	Eval(n Node) Node
}

type vm struct {
	env Env
}

func NewVM() VM {
	e := NewEnvironment()
	vm := &vm{e}
	e.Set("+", NewFunNode(vm.evalAdd))
	e.Set("def", NewFunNode(vm.defVar))
	return vm
}

func (vm *vm) Eval(n Node) Node {
	return vm.eval(vm.env, n)
}

func (vm *vm) eval(e Env, n Node) Node {
	switch v := n.(type) {
	case SymNode:
		return e.Get(v.Name())
	case *sExprNode:
		return vm.evalSExpr(e, v)
	}
	// Return Numbers, Symbols and Q-Expressions as-is.
	return n
}

func (vm *vm) evalSExpr(e Env, n *sExprNode) Node {
	if n.Len() == 0 {
		return n
	}

	f := vm.eval(e, n.Cell(0))
	f1, ok := f.(FunNode)
	if !ok {
		fmt.Printf("First element of a S-Expr must be a Function-Symbol but is [%v].\n", f1)
		return n
	}

	args := []Node{}
	for i := 1; i < n.Len(); i++ {
		args = append(args, vm.Eval(n.Cell(i)))
	}

	return f1.Apply(e, args)
}

func (vm *vm) evalAdd(e Env, args []Node) Node {
	var sum int32
	for _, arg := range args {
		switch v := arg.(type) {
		case NumNode:
			sum += v.Value()
			break
		default:
			fmt.Printf("Cannot add the non-number [%v].\n", v)
			break
		}
	}
	return &numNode{sum}
}

func (vm *vm) defVar(e Env, args []Node) Node {
	len := len(args)
	if len < 2 {
		fmt.Printf("Define requires at least 2 arguments.\n")
		return NewSExprNode()
	}

	params, ok := args[0].(*qExprNode)
	if !ok {
		fmt.Printf("Define requires the first argument to be a Q-Expression.\n")
		return NewSExprNode()
	}

	if params.Len() != len-1 {
		fmt.Printf("Number of defined names [%d] and definitions [%d] must be the same.\n", params.Len(), len-1)
		return NewSExprNode()
	}

	for i := 0; i < params.Len(); i++ {
		param, ok := params.Cell(i).(SymNode)
		if !ok {
			fmt.Printf("Parameter name [%d] is not a Symbol. Definition names can only be Symbols.\n", i+1)
			return NewSExprNode()
		}
		e.Set(param.Name(), vm.eval(e, args[i+1]))
	}

	return NewSExprNode()
}
