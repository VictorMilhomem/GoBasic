package parser

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *Visitor) VisitProg(ctx *ProgContext) interface{} {
	fmt.Println("Visited")
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAmperoper(ctx *AmperoperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLinenumber(ctx *LinenumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAmprstmt(ctx *AmprstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVardecl(ctx *VardeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPrintstmt1(ctx *Printstmt1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPrintlist(ctx *PrintlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitGetstmt(ctx *GetstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLetstmt(ctx *LetstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVariableassignment(ctx *VariableassignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitRelop(ctx *RelopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitNeq(ctx *NeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitIfstmt(ctx *IfstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitForstmt1(ctx *Forstmt1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitForstmt2(ctx *Forstmt2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitNextstmt(ctx *NextstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInputstmt(ctx *InputstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitReadstmt(ctx *ReadstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitDimstmt(ctx *DimstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitGotostmt(ctx *GotostmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitGosubstmt(ctx *GosubstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPokestmt(ctx *PokestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitCallstmt(ctx *CallstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitHplotstmt(ctx *HplotstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVplotstmt(ctx *VplotstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPlotstmt(ctx *PlotstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitOngotostmt(ctx *OngotostmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitOngosubstmt(ctx *OngosubstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVtabstmnt(ctx *VtabstmntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitHtabstmnt(ctx *HtabstmntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitHimemstmt(ctx *HimemstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLomemstmt(ctx *LomemstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitDatastmt(ctx *DatastmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitDatum(ctx *DatumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitWaitstmt(ctx *WaitstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitXdrawstmt(ctx *XdrawstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitDrawstmt(ctx *DrawstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitDefstmt(ctx *DefstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTabstmt(ctx *TabstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSpeedstmt(ctx *SpeedstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitRotstmt(ctx *RotstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitScalestmt(ctx *ScalestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitColorstmt(ctx *ColorstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitHcolorstmt(ctx *HcolorstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitHlinstmt(ctx *HlinstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVlinstmt(ctx *VlinstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitOnerrstmt(ctx *OnerrstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPrstmt(ctx *PrstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInstmt(ctx *InstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStorestmt(ctx *StorestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitRecallstmt(ctx *RecallstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitListstmt(ctx *ListstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPopstmt(ctx *PopstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAmptstmt(ctx *AmptstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitIncludestmt(ctx *IncludestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEndstmt(ctx *EndstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitReturnstmt(ctx *ReturnstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitRestorestmt(ctx *RestorestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunc_(ctx *Func_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSignExpression(ctx *SignExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExponentExpression(ctx *ExponentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMultiplyingExpression(ctx *MultiplyingExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAddingExpression(ctx *AddingExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVar_(ctx *Var_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVarname(ctx *VarnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVarsuffix(ctx *VarsuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVarlist(ctx *VarlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExprlist(ctx *ExprlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSqrfunc(ctx *SqrfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitChrfunc(ctx *ChrfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLenfunc(ctx *LenfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAscfunc(ctx *AscfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMidfunc(ctx *MidfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPdlfunc(ctx *PdlfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPeekfunc(ctx *PeekfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitIntfunc(ctx *IntfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSpcfunc(ctx *SpcfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFrefunc(ctx *FrefuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPosfunc(ctx *PosfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUsrfunc(ctx *UsrfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLeftfunc(ctx *LeftfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitRightfunc(ctx *RightfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStrfunc(ctx *StrfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFnfunc(ctx *FnfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitValfunc(ctx *ValfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitScrnfunc(ctx *ScrnfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSinfunc(ctx *SinfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitCosfunc(ctx *CosfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTanfunc(ctx *TanfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAtnfunc(ctx *AtnfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitRndfunc(ctx *RndfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSgnfunc(ctx *SgnfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExpfunc(ctx *ExpfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLogfunc(ctx *LogfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAbsfunc(ctx *AbsfuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTabfunc(ctx *TabfuncContext) interface{} {
	return v.VisitChildren(ctx)
}
