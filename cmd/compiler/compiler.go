package compiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

type Compiler struct {
	Module    *ir.Module
	Functions BasicFunc
	Entry     *ir.Block
	MainFunc  *ir.Func
}

type BasicFunc map[string]*ir.Func

func (c *Compiler) SetFunc(name string, val *ir.Func) {
	c.Functions[name] = val
}

func (c *Compiler) GetFunc(name string) (*ir.Func, bool) {
	val, ok := c.Functions[name]
	return val, ok
}

func NewCompiler() *Compiler {
	return &Compiler{
		Module:    ir.NewModule(),
		Functions: make(map[string]*ir.Func),
	}
}

func (c *Compiler) Set() {
	print := c.Module.NewFunc("printf", types.I32, ir.NewParam("", types.NewPointer(types.I8)))
	c.SetFunc("printf", print)
}

func (c *Compiler) Main() {
	c.MainFunc = c.Module.NewFunc("main", types.I32)
	c.Entry = c.MainFunc.NewBlock("entry")
}

func (c *Compiler) MainRet() {
	c.Entry.NewRet(constant.NewInt(types.I32, 0))
}
