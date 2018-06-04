package main

import "fmt"

type Env interface {
	Get(name string) Node
	Set(name string, val Node)
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
	fmt.Printf("Undefined name [%s].\n", name)
	return nil
}

func (e *env) Set(name string, val Node) {
	//fmt.Printf("def [%s] = [%v].\n", name, val)
	e.pool[name] = val
}
