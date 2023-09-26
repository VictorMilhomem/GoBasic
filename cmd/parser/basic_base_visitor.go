// Code generated from ./cmd/parser/Basic.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Basic

import "github.com/antlr4-go/antlr/v4"

type BaseBasicVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseBasicVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitExprlist(ctx *ExprlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitVarlist(ctx *VarlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitUnary(ctx *UnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitMutlop(ctx *MutlopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitVara(ctx *VaraContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitRelop(ctx *RelopContext) interface{} {
	return v.VisitChildren(ctx)
}
