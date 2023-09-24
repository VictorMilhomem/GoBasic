package compiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

type Compiler struct {
	Module    *ir.Module
	Functions BasicFunc
	MainBlock *ir.Block
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
	print := c.Module.NewFunc("puts", types.I32, ir.NewParam("", types.NewPointer(types.I8)))
	c.SetFunc("puts", print)
}

func (c *Compiler) Main() {
	main := c.Module.NewFunc("main", types.I32)
	c.MainBlock = main.NewBlock("")
}

func (c *Compiler) MainRet() {
	c.MainBlock.NewRet(constant.NewInt(types.I32, 0))
}
