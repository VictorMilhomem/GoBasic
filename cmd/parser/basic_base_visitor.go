// Code generated from ./cmd/parser/Basic.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Basic

import "github.com/antlr4-go/antlr/v4"

type BaseBasicVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseBasicVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitAmperoper(ctx *AmperoperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitLinenumber(ctx *LinenumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitAmprstmt(ctx *AmprstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitVardecl(ctx *VardeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitPrintstmt1(ctx *Printstmt1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitPrintlist(ctx *PrintlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitGetstmt(ctx *GetstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitLetstmt(ctx *LetstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitVariableassignment(ctx *VariableassignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitRelop(ctx *RelopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitNeq(ctx *NeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitIfstmt(ctx *IfstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitForstmt1(ctx *Forstmt1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitForstmt2(ctx *Forstmt2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitNextstmt(ctx *NextstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitInputstmt(ctx *InputstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitReadstmt(ctx *ReadstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitDimstmt(ctx *DimstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitGotostmt(ctx *GotostmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitGosubstmt(ctx *GosubstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitPokestmt(ctx *PokestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitCallstmt(ctx *CallstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitHplotstmt(ctx *HplotstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitVplotstmt(ctx *VplotstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitPlotstmt(ctx *PlotstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitOngotostmt(ctx *OngotostmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitOngosubstmt(ctx *OngosubstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitVtabstmnt(ctx *VtabstmntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitHtabstmnt(ctx *HtabstmntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitHimemstmt(ctx *HimemstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitLomemstmt(ctx *LomemstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitDatastmt(ctx *DatastmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitDatum(ctx *DatumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitWaitstmt(ctx *WaitstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitXdrawstmt(ctx *XdrawstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitDrawstmt(ctx *DrawstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitDefstmt(ctx *DefstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitTabstmt(ctx *TabstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitSpeedstmt(ctx *SpeedstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitRotstmt(ctx *RotstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitScalestmt(ctx *ScalestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitColorstmt(ctx *ColorstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitHcolorstmt(ctx *HcolorstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitHlinstmt(ctx *HlinstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitVlinstmt(ctx *VlinstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitOnerrstmt(ctx *OnerrstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitPrstmt(ctx *PrstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitInstmt(ctx *InstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitStorestmt(ctx *StorestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitRecallstmt(ctx *RecallstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitListstmt(ctx *ListstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitPopstmt(ctx *PopstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitAmptstmt(ctx *AmptstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitIncludestmt(ctx *IncludestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitEndstmt(ctx *EndstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitReturnstmt(ctx *ReturnstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitRestorestmt(ctx *RestorestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitFunc_(ctx *Func_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitSignExpression(ctx *SignExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitExponentExpression(ctx *ExponentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitMultiplyingExpression(ctx *MultiplyingExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitAddingExpression(ctx *AddingExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitVar_(ctx *Var_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitVarname(ctx *VarnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitVarsuffix(ctx *VarsuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitVarlist(ctx *VarlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitExprlist(ctx *ExprlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitSqrfunc(ctx *SqrfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitChrfunc(ctx *ChrfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitLenfunc(ctx *LenfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitAscfunc(ctx *AscfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitMidfunc(ctx *MidfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitPdlfunc(ctx *PdlfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitPeekfunc(ctx *PeekfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitIntfunc(ctx *IntfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitSpcfunc(ctx *SpcfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitFrefunc(ctx *FrefuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitPosfunc(ctx *PosfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitUsrfunc(ctx *UsrfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitLeftfunc(ctx *LeftfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitRightfunc(ctx *RightfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitStrfunc(ctx *StrfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitFnfunc(ctx *FnfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitValfunc(ctx *ValfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitScrnfunc(ctx *ScrnfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitSinfunc(ctx *SinfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitCosfunc(ctx *CosfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitTanfunc(ctx *TanfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitAtnfunc(ctx *AtnfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitRndfunc(ctx *RndfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitSgnfunc(ctx *SgnfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitExpfunc(ctx *ExpfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitLogfunc(ctx *LogfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitAbsfunc(ctx *AbsfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBasicVisitor) VisitTabfunc(ctx *TabfuncContext) interface{} {
	return v.VisitChildren(ctx)
}
