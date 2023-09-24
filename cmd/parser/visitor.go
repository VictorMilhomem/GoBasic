package parser

import (
	"fmt"
	"strings"

	"github.com/VictorMilhomem/Basic/cmd/compiler"
	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

type Visitor struct {
	BaseBasicVisitor
	env   *Environment
	Lines *Environment
	comp  *compiler.Compiler
}

func NewVisitor(comp *compiler.Compiler) Visitor {
	return Visitor{
		env:   NewEnvironment(),
		Lines: NewEnvironment(),
		comp:  comp,
	}
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	// fmt.Printf("visit input type: %v\n", reflect.TypeOf(tree))
	switch t := tree.(type) {
	case *antlr.ErrorNodeImpl:
		fmt.Printf("Error: - %v\n", t.GetText())
		return nil
	default:
		return tree.Accept(v)
	}
}

func (v *Visitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, n := range node.GetChildren() {
		v.Visit(n.(antlr.ParseTree))
	}
	return nil
}

func (v *Visitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLine(ctx *LineContext) interface{} {
	v.Lines.Set(ctx.Linenumber().GetText(), ctx.GetChild(1))
	return v.Visit(ctx.Amprstmt(0))
}

func (v *Visitor) VisitAmperoper(ctx *AmperoperContext) interface{} {
	return nil
}

func (v *Visitor) VisitLinenumber(ctx *LinenumberContext) interface{} {
	return nil
}

func (v *Visitor) VisitAmprstmt(ctx *AmprstmtContext) interface{} {
	return v.VisitStatement(ctx.Statement().(*StatementContext))
}

func (v *Visitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitPrintstmt1(ctx.Printstmt1().(*Printstmt1Context))
}

func (v *Visitor) VisitVardecl(ctx *VardeclContext) interface{} {
	return nil
}

func (v *Visitor) VisitPrintstmt1(ctx *Printstmt1Context) interface{} {
	return v.VisitPrintlist(ctx.Printlist().(*PrintlistContext))
}

func (v *Visitor) VisitPrintlist(ctx *PrintlistContext) interface{} {
	val := v.VisitExpression(ctx.Expression(0).(*ExpressionContext)).(*constant.ExprGetElementPtr)
	v.comp.MainBlock.NewCall(v.comp.Functions["puts"], val)
	return nil
}

func (v *Visitor) VisitGetstmt(ctx *GetstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitLetstmt(ctx *LetstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitVariableassignment(ctx *VariableassignmentContext) interface{} {
	return nil
}

func (v *Visitor) VisitRelop(ctx *RelopContext) interface{} {
	return nil
}

func (v *Visitor) VisitNeq(ctx *NeqContext) interface{} {
	return nil
}

func (v *Visitor) VisitIfstmt(ctx *IfstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitForstmt1(ctx *Forstmt1Context) interface{} {
	return nil
}

func (v *Visitor) VisitForstmt2(ctx *Forstmt2Context) interface{} {
	return nil
}

func (v *Visitor) VisitNextstmt(ctx *NextstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitInputstmt(ctx *InputstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitReadstmt(ctx *ReadstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitDimstmt(ctx *DimstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitGotostmt(ctx *GotostmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitGosubstmt(ctx *GosubstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitPokestmt(ctx *PokestmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitCallstmt(ctx *CallstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitHplotstmt(ctx *HplotstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitVplotstmt(ctx *VplotstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitPlotstmt(ctx *PlotstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitOngotostmt(ctx *OngotostmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitOngosubstmt(ctx *OngosubstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitVtabstmnt(ctx *VtabstmntContext) interface{} {
	return nil
}

func (v *Visitor) VisitHtabstmnt(ctx *HtabstmntContext) interface{} {
	return nil
}

func (v *Visitor) VisitHimemstmt(ctx *HimemstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitLomemstmt(ctx *LomemstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitDatastmt(ctx *DatastmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitDatum(ctx *DatumContext) interface{} {
	return nil
}

func (v *Visitor) VisitWaitstmt(ctx *WaitstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitXdrawstmt(ctx *XdrawstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitDrawstmt(ctx *DrawstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitDefstmt(ctx *DefstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitTabstmt(ctx *TabstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitSpeedstmt(ctx *SpeedstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitRotstmt(ctx *RotstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitScalestmt(ctx *ScalestmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitColorstmt(ctx *ColorstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitHcolorstmt(ctx *HcolorstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitHlinstmt(ctx *HlinstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitVlinstmt(ctx *VlinstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitOnerrstmt(ctx *OnerrstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitPrstmt(ctx *PrstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitInstmt(ctx *InstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitStorestmt(ctx *StorestmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitRecallstmt(ctx *RecallstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitListstmt(ctx *ListstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitPopstmt(ctx *PopstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitAmptstmt(ctx *AmptstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitIncludestmt(ctx *IncludestmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitEndstmt(ctx *EndstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitReturnstmt(ctx *ReturnstmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitRestorestmt(ctx *RestorestmtContext) interface{} {
	return nil
}

func (v *Visitor) VisitNumber(ctx *NumberContext) interface{} {
	return nil
}

func (v *Visitor) VisitFunc_(ctx *Func_Context) interface{} {
	return nil
}

func (v *Visitor) VisitSignExpression(ctx *SignExpressionContext) interface{} {
	return nil
}

func (v *Visitor) VisitExponentExpression(ctx *ExponentExpressionContext) interface{} {
	return nil
}

func (v *Visitor) VisitMultiplyingExpression(ctx *MultiplyingExpressionContext) interface{} {
	return nil
}

func (v *Visitor) VisitAddingExpression(ctx *AddingExpressionContext) interface{} {
	return nil
}

func (v *Visitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return nil
}

func (v *Visitor) VisitExpression(ctx *ExpressionContext) interface{} {
	text := ctx.Func_().STRINGLITERAL().GetText()
	trimmedTxt := strings.Trim(text, `"`)
	val := constant.NewCharArrayFromString(trimmedTxt + "\x00")
	str := v.comp.Module.NewGlobalDef("str", val)
	zero := constant.NewInt(types.I64, 0)
	gep := constant.NewGetElementPtr(val.Typ, str, zero, zero)
	return gep
}

func (v *Visitor) VisitVar_(ctx *Var_Context) interface{} {
	return nil
}

func (v *Visitor) VisitVarname(ctx *VarnameContext) interface{} {
	return nil
}

func (v *Visitor) VisitVarsuffix(ctx *VarsuffixContext) interface{} {
	return nil
}

func (v *Visitor) VisitVarlist(ctx *VarlistContext) interface{} {
	return nil
}

func (v *Visitor) VisitExprlist(ctx *ExprlistContext) interface{} {
	return nil
}

func (v *Visitor) VisitSqrfunc(ctx *SqrfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitChrfunc(ctx *ChrfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitLenfunc(ctx *LenfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitAscfunc(ctx *AscfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitMidfunc(ctx *MidfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitPdlfunc(ctx *PdlfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitPeekfunc(ctx *PeekfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitIntfunc(ctx *IntfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitSpcfunc(ctx *SpcfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitFrefunc(ctx *FrefuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitPosfunc(ctx *PosfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitUsrfunc(ctx *UsrfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitLeftfunc(ctx *LeftfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitRightfunc(ctx *RightfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitStrfunc(ctx *StrfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitFnfunc(ctx *FnfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitValfunc(ctx *ValfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitScrnfunc(ctx *ScrnfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitSinfunc(ctx *SinfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitCosfunc(ctx *CosfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitTanfunc(ctx *TanfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitAtnfunc(ctx *AtnfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitRndfunc(ctx *RndfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitSgnfunc(ctx *SgnfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitExpfunc(ctx *ExpfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitLogfunc(ctx *LogfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitAbsfunc(ctx *AbsfuncContext) interface{} {
	return nil
}

func (v *Visitor) VisitTabfunc(ctx *TabfuncContext) interface{} {
	return nil
}
