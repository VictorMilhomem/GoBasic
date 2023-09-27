package compiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

type LLVMType int

const (
	Int32 LLVMType = iota
	Int64
	Float64
	String
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

func (c *Compiler) Init() {
	c.Set()
	c.Main()
}

func (c *Compiler) Set() {
	print := c.Module.NewFunc("printf", types.I32, ir.NewParam("", types.NewPointer(types.I8)))
	c.SetFunc("printf", print)
}

func (c *Compiler) Main() {
	c.MainFunc = c.Module.NewFunc("main", types.I32)
	c.Entry = c.MainFunc.NewBlock("entry")
}

func (c *Compiler) NewBlock(line string) *ir.Block {
	return c.MainFunc.NewBlock(line)
}

func (c *Compiler) NewGlobal(name string, ttype LLVMType, value interface{}) *ir.Global {
	switch ttype {
	case Int32:
		val := constant.NewInt(types.I32, value.(int64))
		return c.Module.NewGlobalDef(name, val)
	case Int64:
		val := constant.NewInt(types.I64, value.(int64))
		return c.Module.NewGlobalDef(name, val)
	case Float64:
		val := constant.NewFloat(types.Float, value.(float64))
		return c.Module.NewGlobalDef(name, val)
	case String:
		val := constant.NewCharArrayFromString(value.(string) + "\x00")
		return c.Module.NewGlobalDef(name, val)
	default:
		return nil
	}
}

func (c *Compiler) GetStringPointer(name *ir.Global, value *constant.CharArray) *constant.ExprGetElementPtr {
	zero := constant.NewInt(types.I64, 0)
	ret := constant.NewGetElementPtr(value.Typ, name, zero, zero)
	return ret
}
