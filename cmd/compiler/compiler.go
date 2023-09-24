package compiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type Compiler struct {
	module    *ir.Module
	functions BasicFunc
}

type BasicFunc map[string]*ir.Func

func (c *Compiler) SetFunc(name string, val *ir.Func) {
	c.functions[name] = val
}

func (c *Compiler) GetFunc(name string) (*ir.Func, bool) {
	val, ok := c.functions[name]
	return val, ok
}

func NewCompiler() *Compiler {
	return &Compiler{
		module:    ir.NewModule(),
		functions: make(map[string]*ir.Func),
	}
}

func (c *Compiler) Set() {
	print := c.module.NewFunc("print", types.I32, ir.NewParam("", types.NewPointer(types.I8)))
	c.SetFunc("print", print)
}
