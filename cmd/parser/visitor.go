package parser

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	BaseBasicVisitor
	env *Environment
}

var currentLine string

func NewVisitor() Visitor {
	return Visitor{
		env: NewEnvironment(),
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

func (v *Visitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.Visit(ctx.Line(0))
}

func (v *Visitor) VisitLine(ctx *LineContext) interface{} {
	switch {
	case ctx.Statement() != nil && ctx.Number() != nil:
		currentLine = strings.Trim(ctx.Number().GetText(), "\"")
		return v.Visit(ctx.Statement())
	default:
		return v.Visit(ctx.Statement())
	}
}

func (v *Visitor) VisitStatement(ctx *StatementContext) interface{} {
	stmt := fmt.Sprint(ctx.GetChild(0))
	switch stmt {
	case "PRINT":
		return v.Visit(ctx.Exprlist())
	case "IF":
		return nil
	case "GOTO":
		return nil
	case "INPUT":
		return nil
	case "LET":
		return nil
	case "GOSUB":
		return nil
	case "RETURN":
		return nil
	case "CLEAR":
		return nil
	case "LIST":
		return nil
	case "RUN":
		return nil
	case "END":
		return nil
	}

	return nil
}

func (v *Visitor) VisitExprlist(ctx *ExprlistContext) interface{} {
	switch {
	case ctx.STRING(0) != nil:
		str := strings.Trim(ctx.STRING(0).GetText(), "\"")
		fmt.Print(str)
	}
	return nil
}
