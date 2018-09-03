package main

type Env interface {
	Get(name string) Node
	Set(name string, val Node)
	SetNum(name string, num float64)
	SetFun(name string, fun Fun)
	All() map[string]Node
	Copy() Env
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
		return val.Copy()
	}
	if e.parent != nil {
		return e.parent.Get(name)
	}
	return NewErrNode("Unbound symbol [%s].", name)
}

func (e *env) Set(name string, val Node) {
	e.pool[name] = val.Copy()
}

func (e *env) SetNum(name string, num float64) {
	e.Set(name, &numNode{num})
}

func (e *env) SetFun(name string, fun Fun) {
	e.Set(name, NewFunNode(fun))
}

func (e *env) All() map[string]Node {
	return e.pool
}

func (e *env) Copy() Env {
	pool := map[string]Node{}
	f := &env{e.parent, pool}
	for k, v := range e.pool {
		f.pool[k] = v.Copy()
	}
	return f
}
