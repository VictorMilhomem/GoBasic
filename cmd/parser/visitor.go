package parser

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	BaseBasicVisitor
	env   *Environment
	lines *Environment // store all lines evaluation
}

var (
	currentLine string
	_goto       bool
)

func NewVisitor() Visitor {
	return Visitor{
		env:   NewEnvironment(),
		lines: NewEnvironment(),
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
		v.lines.Set(currentLine, ctx.Statement())
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
		line_stmt, _ := v.lines.Get(strings.Trim(ctx.Number().GetText(), "\""))
		return v.Visit(line_stmt.(antlr.ParseTree))
	case "INPUT":
		return nil
	case "LET":
		name := strings.Trim(ctx.Vara().GetText(), "\"")
		return v.env.Set(name, v.Visit(ctx.Expression(0)))
	case "GOSUB":
		return v.Visit(ctx.Expression(0))
	case "RETURN":
		return nil
	case "CLEAR":
		clearScreen()
		return nil
	case "LIST":
		return nil
	case "RUN":
		return nil
	case "END":
		return nil
	default:
		if ctx.Vara() != nil {
			name := strings.Trim(ctx.Vara().GetText(), "\"")
			if _, ok := v.env.Get(name); ok {
				val := v.Visit(ctx.Expression(0))
				return v.env.Set(name, val)
			}

		}
		return nil

	}
}

func (v *Visitor) VisitExprlist(ctx *ExprlistContext) interface{} {
	switch {
	case ctx.STRING(0) != nil:
		str := strings.Trim(ctx.STRING(0).GetText(), "\"")
		if strings.Contains(str, "\\n") {
			str = strings.ReplaceAll(str, "\\n", "\n")
		}
		fmt.Printf("%v", str)
		return str
	default:
		val := v.Visit(ctx.Expression(0))
		fmt.Printf("%v", val)
		return val
	}
}

func (v *Visitor) VisitExpression(ctx *ExpressionContext) interface{} {
	leftVal := v.Visit(ctx.Term(0)).(float64)

	if ctx.Unary() != nil {
		op := ctx.Unary().GetText()
		if op == "-" {
			leftVal = -leftVal
		}
	}

	for i := 1; i < len(ctx.AllTerm()); i++ {
		var op string
		if ctx.GetChild(2) != nil {
			op = fmt.Sprint(ctx.GetChildren()[1])
		}

		rightVal := v.Visit(ctx.Term(i)).(float64)

		switch op {
		case "+":
			leftVal += rightVal
		case "-":
			leftVal -= rightVal
		}
	}

	return leftVal
}

func (v *Visitor) VisitTerm(ctx *TermContext) interface{} {
	result := v.Visit(ctx.Factor(0)).(float64)

	for i := 1; i < len(ctx.AllFactor()); i++ {
		opCtx := ctx.Mutlop(i - 1)
		op := opCtx.GetText()
		right := v.Visit(ctx.Factor(i)).(float64)

		switch op {
		case "*":
			result *= right
		case "/":
			if right != 0 {
				result /= right
			} else {
			}
		}
	}

	return result
}

func (v *Visitor) VisitFactor(ctx *FactorContext) interface{} {
	switch {
	case ctx.Number() != nil:
		return v.Visit(ctx.Number())
	case ctx.Vara() != nil:
		return v.Visit(ctx.Vara())
	default:
		return nil
	}
}

func (v *Visitor) VisitVara(ctx *VaraContext) interface{} {
	name := strings.Trim(ctx.GetText(), "\"")
	val, _ := v.env.Get(name)
	return val
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

func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Limpeza de tela não suportada neste sistema operacional.")
	}
}
