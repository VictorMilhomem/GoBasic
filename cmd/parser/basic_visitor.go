// Code generated from ./cmd/parser/Basic.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Basic

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by BasicParser.
type BasicVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BasicParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by BasicParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by BasicParser#amperoper.
	VisitAmperoper(ctx *AmperoperContext) interface{}

	// Visit a parse tree produced by BasicParser#linenumber.
	VisitLinenumber(ctx *LinenumberContext) interface{}

	// Visit a parse tree produced by BasicParser#amprstmt.
	VisitAmprstmt(ctx *AmprstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by BasicParser#vardecl.
	VisitVardecl(ctx *VardeclContext) interface{}

	// Visit a parse tree produced by BasicParser#printstmt1.
	VisitPrintstmt1(ctx *Printstmt1Context) interface{}

	// Visit a parse tree produced by BasicParser#printlist.
	VisitPrintlist(ctx *PrintlistContext) interface{}

	// Visit a parse tree produced by BasicParser#getstmt.
	VisitGetstmt(ctx *GetstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#letstmt.
	VisitLetstmt(ctx *LetstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#variableassignment.
	VisitVariableassignment(ctx *VariableassignmentContext) interface{}

	// Visit a parse tree produced by BasicParser#relop.
	VisitRelop(ctx *RelopContext) interface{}

	// Visit a parse tree produced by BasicParser#neq.
	VisitNeq(ctx *NeqContext) interface{}

	// Visit a parse tree produced by BasicParser#ifstmt.
	VisitIfstmt(ctx *IfstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#forstmt1.
	VisitForstmt1(ctx *Forstmt1Context) interface{}

	// Visit a parse tree produced by BasicParser#forstmt2.
	VisitForstmt2(ctx *Forstmt2Context) interface{}

	// Visit a parse tree produced by BasicParser#nextstmt.
	VisitNextstmt(ctx *NextstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#inputstmt.
	VisitInputstmt(ctx *InputstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#readstmt.
	VisitReadstmt(ctx *ReadstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#dimstmt.
	VisitDimstmt(ctx *DimstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#gotostmt.
	VisitGotostmt(ctx *GotostmtContext) interface{}

	// Visit a parse tree produced by BasicParser#gosubstmt.
	VisitGosubstmt(ctx *GosubstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#pokestmt.
	VisitPokestmt(ctx *PokestmtContext) interface{}

	// Visit a parse tree produced by BasicParser#callstmt.
	VisitCallstmt(ctx *CallstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#hplotstmt.
	VisitHplotstmt(ctx *HplotstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#vplotstmt.
	VisitVplotstmt(ctx *VplotstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#plotstmt.
	VisitPlotstmt(ctx *PlotstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#ongotostmt.
	VisitOngotostmt(ctx *OngotostmtContext) interface{}

	// Visit a parse tree produced by BasicParser#ongosubstmt.
	VisitOngosubstmt(ctx *OngosubstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#vtabstmnt.
	VisitVtabstmnt(ctx *VtabstmntContext) interface{}

	// Visit a parse tree produced by BasicParser#htabstmnt.
	VisitHtabstmnt(ctx *HtabstmntContext) interface{}

	// Visit a parse tree produced by BasicParser#himemstmt.
	VisitHimemstmt(ctx *HimemstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#lomemstmt.
	VisitLomemstmt(ctx *LomemstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#datastmt.
	VisitDatastmt(ctx *DatastmtContext) interface{}

	// Visit a parse tree produced by BasicParser#datum.
	VisitDatum(ctx *DatumContext) interface{}

	// Visit a parse tree produced by BasicParser#waitstmt.
	VisitWaitstmt(ctx *WaitstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#xdrawstmt.
	VisitXdrawstmt(ctx *XdrawstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#drawstmt.
	VisitDrawstmt(ctx *DrawstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#defstmt.
	VisitDefstmt(ctx *DefstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#tabstmt.
	VisitTabstmt(ctx *TabstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#speedstmt.
	VisitSpeedstmt(ctx *SpeedstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#rotstmt.
	VisitRotstmt(ctx *RotstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#scalestmt.
	VisitScalestmt(ctx *ScalestmtContext) interface{}

	// Visit a parse tree produced by BasicParser#colorstmt.
	VisitColorstmt(ctx *ColorstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#hcolorstmt.
	VisitHcolorstmt(ctx *HcolorstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#hlinstmt.
	VisitHlinstmt(ctx *HlinstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#vlinstmt.
	VisitVlinstmt(ctx *VlinstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#onerrstmt.
	VisitOnerrstmt(ctx *OnerrstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#prstmt.
	VisitPrstmt(ctx *PrstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#instmt.
	VisitInstmt(ctx *InstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#storestmt.
	VisitStorestmt(ctx *StorestmtContext) interface{}

	// Visit a parse tree produced by BasicParser#recallstmt.
	VisitRecallstmt(ctx *RecallstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#liststmt.
	VisitListstmt(ctx *ListstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#popstmt.
	VisitPopstmt(ctx *PopstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#amptstmt.
	VisitAmptstmt(ctx *AmptstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#includestmt.
	VisitIncludestmt(ctx *IncludestmtContext) interface{}

	// Visit a parse tree produced by BasicParser#endstmt.
	VisitEndstmt(ctx *EndstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#returnstmt.
	VisitReturnstmt(ctx *ReturnstmtContext) interface{}

	// Visit a parse tree produced by BasicParser#restorestmt.
	VisitRestorestmt(ctx *RestorestmtContext) interface{}

	// Visit a parse tree produced by BasicParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by BasicParser#func_.
	VisitFunc_(ctx *Func_Context) interface{}

	// Visit a parse tree produced by BasicParser#signExpression.
	VisitSignExpression(ctx *SignExpressionContext) interface{}

	// Visit a parse tree produced by BasicParser#exponentExpression.
	VisitExponentExpression(ctx *ExponentExpressionContext) interface{}

	// Visit a parse tree produced by BasicParser#multiplyingExpression.
	VisitMultiplyingExpression(ctx *MultiplyingExpressionContext) interface{}

	// Visit a parse tree produced by BasicParser#addingExpression.
	VisitAddingExpression(ctx *AddingExpressionContext) interface{}

	// Visit a parse tree produced by BasicParser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by BasicParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by BasicParser#var_.
	VisitVar_(ctx *Var_Context) interface{}

	// Visit a parse tree produced by BasicParser#varname.
	VisitVarname(ctx *VarnameContext) interface{}

	// Visit a parse tree produced by BasicParser#varsuffix.
	VisitVarsuffix(ctx *VarsuffixContext) interface{}

	// Visit a parse tree produced by BasicParser#varlist.
	VisitVarlist(ctx *VarlistContext) interface{}

	// Visit a parse tree produced by BasicParser#exprlist.
	VisitExprlist(ctx *ExprlistContext) interface{}

	// Visit a parse tree produced by BasicParser#sqrfunc.
	VisitSqrfunc(ctx *SqrfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#chrfunc.
	VisitChrfunc(ctx *ChrfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#lenfunc.
	VisitLenfunc(ctx *LenfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#ascfunc.
	VisitAscfunc(ctx *AscfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#midfunc.
	VisitMidfunc(ctx *MidfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#pdlfunc.
	VisitPdlfunc(ctx *PdlfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#peekfunc.
	VisitPeekfunc(ctx *PeekfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#intfunc.
	VisitIntfunc(ctx *IntfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#spcfunc.
	VisitSpcfunc(ctx *SpcfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#frefunc.
	VisitFrefunc(ctx *FrefuncContext) interface{}

	// Visit a parse tree produced by BasicParser#posfunc.
	VisitPosfunc(ctx *PosfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#usrfunc.
	VisitUsrfunc(ctx *UsrfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#leftfunc.
	VisitLeftfunc(ctx *LeftfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#rightfunc.
	VisitRightfunc(ctx *RightfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#strfunc.
	VisitStrfunc(ctx *StrfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#fnfunc.
	VisitFnfunc(ctx *FnfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#valfunc.
	VisitValfunc(ctx *ValfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#scrnfunc.
	VisitScrnfunc(ctx *ScrnfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#sinfunc.
	VisitSinfunc(ctx *SinfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#cosfunc.
	VisitCosfunc(ctx *CosfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#tanfunc.
	VisitTanfunc(ctx *TanfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#atnfunc.
	VisitAtnfunc(ctx *AtnfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#rndfunc.
	VisitRndfunc(ctx *RndfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#sgnfunc.
	VisitSgnfunc(ctx *SgnfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#expfunc.
	VisitExpfunc(ctx *ExpfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#logfunc.
	VisitLogfunc(ctx *LogfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#absfunc.
	VisitAbsfunc(ctx *AbsfuncContext) interface{}

	// Visit a parse tree produced by BasicParser#tabfunc.
	VisitTabfunc(ctx *TabfuncContext) interface{}
}
