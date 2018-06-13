package main

type Env interface {
	Get(name string) Node
	Set(name string, val Node)
	SetNum(name string, num int32)
	SetFun(name string, fun Fun)
	All() map[string]Node
}

type env struct {
	parent Env
	pool   map[string]Node
}

func NewEnvironment() Env {
	return NewSubEnvironment(nil)
}

func NewSubEnvironment(parent Env) Env {
	pool := map[string]Node{}
	return &env{parent, pool}
}

func (e *env) Get(name string) Node {
	val, ok := e.pool[name]
	if ok {
		return val
	}
	if e.parent != nil {
		return e.parent.Get(name)
	}
	return NewErrNode("Unbound symbol [%s].\n", name)
	// fmt.Printf("Unbound symbol [%s].\n", name)
	// return nil
}

func (e *env) Set(name string, val Node) {
	//fmt.Printf("def [%s] = [%v].\n", name, val)
	e.pool[name] = val
}

func (e *env) SetNum(name string, num int32) {
	e.Set(name, &numNode{num})
}

func (e *env) SetFun(name string, fun Fun) {
	e.Set(name, NewFunNode(fun))
}

func (e *env) All() map[string]Node {
	return e.pool
}
