package parser

import (
	"fmt"
	"strconv"
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
	for _, line := range ctx.AllLine() {
		v.Visit(line)
	}
	return nil
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
		lhs := v.Visit(ctx.Expression(0))
		rhs := v.Visit(ctx.Expression(1))
		relop := strings.Trim(ctx.Relop().GetText(), "\"")
		val := ifExpression(lhs, rhs, relop)
		if val {
			return v.Visit(ctx.Statement())
		}
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

func (v *Visitor) VisitExpression(ctx *ExpressionContext) interface{} {
	prefix := fmt.Sprint(ctx.GetChild(0))
	switch prefix {
	case "-":
		return v.Visit(ctx.Term(0))
	default:
		return v.Visit(ctx.Term(0))
	}
}

func (v *Visitor) VisitTerm(ctx *TermContext) interface{} {
	switch {
	case ctx.Factor(0) != nil:
		return v.Visit(ctx.Factor(0))
	}
	return nil
}

func (v *Visitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.Visit(ctx.Number())
}

func (v *Visitor) VisitNumber(ctx *NumberContext) interface{} {
	num := strings.Trim(ctx.GetText(), "\"")
	val, _ := strconv.ParseFloat(num, 64)
	return val
}

func ifExpression(lhs, rhs interface{}, relop string) bool {
	lhval := lhs.(float64)
	rhval := rhs.(float64)
	switch relop {
	case "<":
		if lhval < rhval {
			return true
		}
		return false
	case ">":
		if lhval > rhval {
			return true
		}
		return false
	case "<=":
		if lhval <= rhval {
			return true
		}
		return false
	case ">=":
		if lhval >= rhval {
			return true
		}
		return false
	case "=":
		if lhval == rhval {
			return true
		}
		return false
	}
	return false
}
