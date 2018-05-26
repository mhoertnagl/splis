package main

import "fmt"

type Env interface {
	Get(name string) Node
	Set(name string, val Node)
}

type env struct {
	pool map[string]Node
}

func NewEnvironment() Env {
	pool := map[string]Node{}
	return &env{pool}
}

func (e *env) Get(name string) Node {
	val, ok := e.pool[name]
	if !ok {
		fmt.Printf("Undefined name [%s].\n", name)
		return nil
	}
	return val
}

func (e *env) Set(name string, val Node) {
	//fmt.Printf("def [%s] = [%v].\n", name, val)
	e.pool[name] = val
}
