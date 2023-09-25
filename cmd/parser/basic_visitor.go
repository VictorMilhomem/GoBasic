// Code generated from ./cmd/parser/Basic.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Basic

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by BasicParser.
type BasicVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BasicParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by BasicParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by BasicParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by BasicParser#exprlist.
	VisitExprlist(ctx *ExprlistContext) interface{}

	// Visit a parse tree produced by BasicParser#varlist.
	VisitVarlist(ctx *VarlistContext) interface{}

	// Visit a parse tree produced by BasicParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by BasicParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by BasicParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by BasicParser#vara.
	VisitVara(ctx *VaraContext) interface{}

	// Visit a parse tree produced by BasicParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by BasicParser#relop.
	VisitRelop(ctx *RelopContext) interface{}
}
