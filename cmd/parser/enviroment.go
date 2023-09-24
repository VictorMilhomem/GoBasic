package parser

import "github.com/llir/llvm/ir"

type Environment struct {
	Values map[string]*ir.Block
}

func NewEnvironment() *Environment {
	return &Environment{
		Values: make(map[string]*ir.Block),
	}
}

func (e *Environment) Get(name string) (*ir.Block, bool) {
	val, ok := e.Values[name]
	return val, ok
}

func (e *Environment) Set(name string, val *ir.Block) interface{} {
	e.Values[name] = val
	return val
}

func CopyEnv(env *Environment) *Environment {
	newEnv := NewEnvironment()
	for k, v := range env.Values {
		newEnv.Values[k] = v
	}
	return newEnv
}
