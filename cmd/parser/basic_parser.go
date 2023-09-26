// Code generated from ./cmd/parser/Basic.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Basic

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type BasicParser struct {
	*antlr.BaseParser
}

var BasicParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func basicParserInit() {
	staticData := &BasicParserStaticData
	staticData.LiteralNames = []string{
		"", "'PRINT'", "'IF'", "'THEN'", "'GOTO'", "'INPUT'", "'LET'", "'='",
		"'GOSUB'", "'RETURN'", "'CLEAR'", "'LIST'", "'RUN'", "'END'", "','",
		"'+'", "'-'", "'*'", "'/'", "'<'", "'>'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "STRING", "DIGIT", "VAR", "CR", "WS",
	}
	staticData.RuleNames = []string{
		"program", "line", "statement", "exprlist", "varlist", "expression",
		"term", "factor", "unary", "mutlop", "vara", "number", "relop",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 143, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 5, 0, 28, 8, 0, 10, 0, 12, 0, 31,
		9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 42, 8,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 51, 8, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 60, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 73, 8, 2, 1, 3, 1, 3, 3, 3,
		77, 8, 3, 1, 3, 1, 3, 1, 3, 3, 3, 82, 8, 3, 5, 3, 84, 8, 3, 10, 3, 12,
		3, 87, 9, 3, 1, 4, 1, 4, 1, 4, 5, 4, 92, 8, 4, 10, 4, 12, 4, 95, 9, 4,
		1, 5, 3, 5, 98, 8, 5, 1, 5, 1, 5, 1, 5, 5, 5, 103, 8, 5, 10, 5, 12, 5,
		106, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 112, 8, 6, 10, 6, 12, 6, 115,
		9, 6, 1, 7, 1, 7, 3, 7, 119, 8, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 11, 4, 11, 128, 8, 11, 11, 11, 12, 11, 129, 1, 12, 1, 12, 3, 12, 134,
		8, 12, 1, 12, 1, 12, 3, 12, 138, 8, 12, 1, 12, 3, 12, 141, 8, 12, 1, 12,
		0, 0, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 6, 1, 1, 24,
		24, 1, 0, 15, 16, 1, 0, 17, 18, 2, 0, 21, 21, 23, 23, 2, 0, 7, 7, 20, 20,
		2, 0, 7, 7, 19, 19, 156, 0, 29, 1, 0, 0, 0, 2, 41, 1, 0, 0, 0, 4, 72, 1,
		0, 0, 0, 6, 76, 1, 0, 0, 0, 8, 88, 1, 0, 0, 0, 10, 97, 1, 0, 0, 0, 12,
		107, 1, 0, 0, 0, 14, 118, 1, 0, 0, 0, 16, 120, 1, 0, 0, 0, 18, 122, 1,
		0, 0, 0, 20, 124, 1, 0, 0, 0, 22, 127, 1, 0, 0, 0, 24, 140, 1, 0, 0, 0,
		26, 28, 3, 2, 1, 0, 27, 26, 1, 0, 0, 0, 28, 31, 1, 0, 0, 0, 29, 27, 1,
		0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 32, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 32,
		33, 5, 0, 0, 1, 33, 1, 1, 0, 0, 0, 34, 35, 3, 22, 11, 0, 35, 36, 3, 4,
		2, 0, 36, 37, 7, 0, 0, 0, 37, 42, 1, 0, 0, 0, 38, 39, 3, 4, 2, 0, 39, 40,
		7, 0, 0, 0, 40, 42, 1, 0, 0, 0, 41, 34, 1, 0, 0, 0, 41, 38, 1, 0, 0, 0,
		42, 3, 1, 0, 0, 0, 43, 44, 5, 1, 0, 0, 44, 73, 3, 6, 3, 0, 45, 46, 5, 2,
		0, 0, 46, 47, 3, 10, 5, 0, 47, 48, 3, 24, 12, 0, 48, 50, 3, 10, 5, 0, 49,
		51, 5, 3, 0, 0, 50, 49, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52, 1, 0, 0,
		0, 52, 53, 3, 4, 2, 0, 53, 73, 1, 0, 0, 0, 54, 55, 5, 4, 0, 0, 55, 73,
		3, 22, 11, 0, 56, 57, 5, 5, 0, 0, 57, 73, 3, 8, 4, 0, 58, 60, 5, 6, 0,
		0, 59, 58, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 62,
		3, 20, 10, 0, 62, 63, 5, 7, 0, 0, 63, 64, 3, 10, 5, 0, 64, 73, 1, 0, 0,
		0, 65, 66, 5, 8, 0, 0, 66, 73, 3, 10, 5, 0, 67, 73, 5, 9, 0, 0, 68, 73,
		5, 10, 0, 0, 69, 73, 5, 11, 0, 0, 70, 73, 5, 12, 0, 0, 71, 73, 5, 13, 0,
		0, 72, 43, 1, 0, 0, 0, 72, 45, 1, 0, 0, 0, 72, 54, 1, 0, 0, 0, 72, 56,
		1, 0, 0, 0, 72, 59, 1, 0, 0, 0, 72, 65, 1, 0, 0, 0, 72, 67, 1, 0, 0, 0,
		72, 68, 1, 0, 0, 0, 72, 69, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 71, 1,
		0, 0, 0, 73, 5, 1, 0, 0, 0, 74, 77, 5, 21, 0, 0, 75, 77, 3, 10, 5, 0, 76,
		74, 1, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77, 85, 1, 0, 0, 0, 78, 81, 5, 14,
		0, 0, 79, 82, 5, 21, 0, 0, 80, 82, 3, 10, 5, 0, 81, 79, 1, 0, 0, 0, 81,
		80, 1, 0, 0, 0, 82, 84, 1, 0, 0, 0, 83, 78, 1, 0, 0, 0, 84, 87, 1, 0, 0,
		0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 7, 1, 0, 0, 0, 87, 85, 1,
		0, 0, 0, 88, 93, 3, 20, 10, 0, 89, 90, 5, 14, 0, 0, 90, 92, 3, 20, 10,
		0, 91, 89, 1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94,
		1, 0, 0, 0, 94, 9, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 96, 98, 3, 16, 8, 0,
		97, 96, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 104, 3,
		12, 6, 0, 100, 101, 7, 1, 0, 0, 101, 103, 3, 12, 6, 0, 102, 100, 1, 0,
		0, 0, 103, 106, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0,
		105, 11, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 107, 113, 3, 14, 7, 0, 108,
		109, 3, 18, 9, 0, 109, 110, 3, 14, 7, 0, 110, 112, 1, 0, 0, 0, 111, 108,
		1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0,
		0, 0, 114, 13, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 119, 3, 20, 10, 0,
		117, 119, 3, 22, 11, 0, 118, 116, 1, 0, 0, 0, 118, 117, 1, 0, 0, 0, 119,
		15, 1, 0, 0, 0, 120, 121, 7, 1, 0, 0, 121, 17, 1, 0, 0, 0, 122, 123, 7,
		2, 0, 0, 123, 19, 1, 0, 0, 0, 124, 125, 7, 3, 0, 0, 125, 21, 1, 0, 0, 0,
		126, 128, 5, 22, 0, 0, 127, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129,
		127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 23, 1, 0, 0, 0, 131, 133, 5,
		19, 0, 0, 132, 134, 7, 4, 0, 0, 133, 132, 1, 0, 0, 0, 133, 134, 1, 0, 0,
		0, 134, 141, 1, 0, 0, 0, 135, 137, 5, 20, 0, 0, 136, 138, 7, 5, 0, 0, 137,
		136, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 141,
		5, 7, 0, 0, 140, 131, 1, 0, 0, 0, 140, 135, 1, 0, 0, 0, 140, 139, 1, 0,
		0, 0, 141, 25, 1, 0, 0, 0, 17, 29, 41, 50, 59, 72, 76, 81, 85, 93, 97,
		104, 113, 118, 129, 133, 137, 140,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// BasicParserInit initializes any static state used to implement BasicParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBasicParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BasicParserInit() {
	staticData := &BasicParserStaticData
	staticData.once.Do(basicParserInit)
}

// NewBasicParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBasicParser(input antlr.TokenStream) *BasicParser {
	BasicParserInit()
	this := new(BasicParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &BasicParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Basic.g4"

	return this
}

// BasicParser tokens.
const (
	BasicParserEOF    = antlr.TokenEOF
	BasicParserT__0   = 1
	BasicParserT__1   = 2
	BasicParserT__2   = 3
	BasicParserT__3   = 4
	BasicParserT__4   = 5
	BasicParserT__5   = 6
	BasicParserT__6   = 7
	BasicParserT__7   = 8
	BasicParserT__8   = 9
	BasicParserT__9   = 10
	BasicParserT__10  = 11
	BasicParserT__11  = 12
	BasicParserT__12  = 13
	BasicParserT__13  = 14
	BasicParserT__14  = 15
	BasicParserT__15  = 16
	BasicParserT__16  = 17
	BasicParserT__17  = 18
	BasicParserT__18  = 19
	BasicParserT__19  = 20
	BasicParserSTRING = 21
	BasicParserDIGIT  = 22
	BasicParserVAR    = 23
	BasicParserCR     = 24
	BasicParserWS     = 25
)

// BasicParser rules.
const (
	BasicParserRULE_program    = 0
	BasicParserRULE_line       = 1
	BasicParserRULE_statement  = 2
	BasicParserRULE_exprlist   = 3
	BasicParserRULE_varlist    = 4
	BasicParserRULE_expression = 5
	BasicParserRULE_term       = 6
	BasicParserRULE_factor     = 7
	BasicParserRULE_unary      = 8
	BasicParserRULE_mutlop     = 9
	BasicParserRULE_vara       = 10
	BasicParserRULE_number     = 11
	BasicParserRULE_relop      = 12
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllLine() []ILineContext
	Line(i int) ILineContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BasicParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(BasicParserEOF, 0)
}

func (s *ProgramContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BasicVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BasicParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BasicParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14696310) != 0 {
		{
			p.SetState(26)
			p.Line()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
		p.Match(BasicParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() INumberContext
	Statement() IStatementContext
	CR() antlr.TerminalNode
	EOF() antlr.TerminalNode

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_line
	return p
}

func InitEmptyLineContext(p *LineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_line
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BasicParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *LineContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *LineContext) CR() antlr.TerminalNode {
	return s.GetToken(BasicParserCR, 0)
}

func (s *LineContext) EOF() antlr.TerminalNode {
	return s.GetToken(BasicParserEOF, 0)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BasicVisitor:
		return t.VisitLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BasicParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BasicParserRULE_line)
	var _la int

	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BasicParserDIGIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Number()
		}
		{
			p.SetState(35)
			p.Statement()
		}
		{
			p.SetState(36)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BasicParserEOF || _la == BasicParserCR) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case BasicParserT__0, BasicParserT__1, BasicParserT__3, BasicParserT__4, BasicParserT__5, BasicParserT__7, BasicParserT__8, BasicParserT__9, BasicParserT__10, BasicParserT__11, BasicParserT__12, BasicParserSTRING, BasicParserVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.Statement()
		}
		{
			p.SetState(39)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BasicParserEOF || _la == BasicParserCR) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exprlist() IExprlistContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	Relop() IRelopContext
	Statement() IStatementContext
	Number() INumberContext
	Varlist() IVarlistContext
	Vara() IVaraContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BasicParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Exprlist() IExprlistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprlistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprlistContext)
}

func (s *StatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *StatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) Relop() IRelopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *StatementContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *StatementContext) Varlist() IVarlistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarlistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarlistContext)
}

func (s *StatementContext) Vara() IVaraContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVaraContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVaraContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BasicVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BasicParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BasicParserRULE_statement)
	var _la int

	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BasicParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)
			p.Match(BasicParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.Exprlist()
		}

	case BasicParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Match(BasicParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.Expression()
		}
		{
			p.SetState(47)
			p.Relop()
		}
		{
			p.SetState(48)
			p.Expression()
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BasicParserT__2 {
			{
				p.SetState(49)
				p.Match(BasicParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(52)
			p.Statement()
		}

	case BasicParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(54)
			p.Match(BasicParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Number()
		}

	case BasicParserT__4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)
			p.Match(BasicParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.Varlist()
		}

	case BasicParserT__5, BasicParserSTRING, BasicParserVAR:
		p.EnterOuterAlt(localctx, 5)
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BasicParserT__5 {
			{
				p.SetState(58)
				p.Match(BasicParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(61)
			p.Vara()
		}
		{
			p.SetState(62)
			p.Match(BasicParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Expression()
		}

	case BasicParserT__7:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(65)
			p.Match(BasicParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Expression()
		}

	case BasicParserT__8:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(67)
			p.Match(BasicParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BasicParserT__9:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(68)
			p.Match(BasicParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BasicParserT__10:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(69)
			p.Match(BasicParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BasicParserT__11:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(70)
			p.Match(BasicParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BasicParserT__12:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(71)
			p.Match(BasicParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprlistContext is an interface to support dynamic dispatch.
type IExprlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsExprlistContext differentiates from other interfaces.
	IsExprlistContext()
}

type ExprlistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprlistContext() *ExprlistContext {
	var p = new(ExprlistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_exprlist
	return p
}

func InitEmptyExprlistContext(p *ExprlistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_exprlist
}

func (*ExprlistContext) IsExprlistContext() {}

func NewExprlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprlistContext {
	var p = new(ExprlistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BasicParserRULE_exprlist

	return p
}

func (s *ExprlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprlistContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(BasicParserSTRING)
}

func (s *ExprlistContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(BasicParserSTRING, i)
}

func (s *ExprlistContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExprlistContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExprlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprlistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BasicVisitor:
		return t.VisitExprlist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BasicParser) Exprlist() (localctx IExprlistContext) {
	localctx = NewExprlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BasicParserRULE_exprlist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(74)
			p.Match(BasicParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(75)
			p.Expression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BasicParserT__13 {
		{
			p.SetState(78)
			p.Match(BasicParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(79)
				p.Match(BasicParserSTRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case 2:
			{
				p.SetState(80)
				p.Expression()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarlistContext is an interface to support dynamic dispatch.
type IVarlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVara() []IVaraContext
	Vara(i int) IVaraContext

	// IsVarlistContext differentiates from other interfaces.
	IsVarlistContext()
}

type VarlistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarlistContext() *VarlistContext {
	var p = new(VarlistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_varlist
	return p
}

func InitEmptyVarlistContext(p *VarlistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_varlist
}

func (*VarlistContext) IsVarlistContext() {}

func NewVarlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarlistContext {
	var p = new(VarlistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BasicParserRULE_varlist

	return p
}

func (s *VarlistContext) GetParser() antlr.Parser { return s.parser }

func (s *VarlistContext) AllVara() []IVaraContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVaraContext); ok {
			len++
		}
	}

	tst := make([]IVaraContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVaraContext); ok {
			tst[i] = t.(IVaraContext)
			i++
		}
	}

	return tst
}

func (s *VarlistContext) Vara(i int) IVaraContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVaraContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVaraContext)
}

func (s *VarlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarlistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BasicVisitor:
		return t.VisitVarlist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BasicParser) Varlist() (localctx IVarlistContext) {
	localctx = NewVarlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BasicParserRULE_varlist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Vara()
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BasicParserT__13 {
		{
			p.SetState(89)
			p.Match(BasicParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Vara()
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext
	Unary() IUnaryContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BasicParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExpressionContext) Unary() IUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BasicVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BasicParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BasicParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BasicParserT__14 || _la == BasicParserT__15 {
		{
			p.SetState(96)
			p.Unary()
		}

	}
	{
		p.SetState(99)
		p.Term()
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BasicParserT__14 || _la == BasicParserT__15 {
		{
			p.SetState(100)
			_la = p.GetTokenStream().LA(1)

			if !(_la == BasicParserT__14 || _la == BasicParserT__15) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(101)
			p.Term()
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFactor() []IFactorContext
	Factor(i int) IFactorContext
	AllMutlop() []IMutlopContext
	Mutlop(i int) IMutlopContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BasicParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFactorContext); ok {
			len++
		}
	}

	tst := make([]IFactorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFactorContext); ok {
			tst[i] = t.(IFactorContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) AllMutlop() []IMutlopContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMutlopContext); ok {
			len++
		}
	}

	tst := make([]IMutlopContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMutlopContext); ok {
			tst[i] = t.(IMutlopContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Mutlop(i int) IMutlopContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMutlopContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMutlopContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BasicVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BasicParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BasicParserRULE_term)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Factor()
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BasicParserT__16 || _la == BasicParserT__17 {
		{
			p.SetState(108)
			p.Mutlop()
		}
		{
			p.SetState(109)
			p.Factor()
		}

		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Vara() IVaraContext
	Number() INumberContext

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BasicParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Vara() IVaraContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVaraContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVaraContext)
}

func (s *FactorContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BasicVisitor:
		return t.VisitFactor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BasicParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BasicParserRULE_factor)
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BasicParserSTRING, BasicParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Vara()
		}

	case BasicParserDIGIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Number()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryContext is an interface to support dynamic dispatch.
type IUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsUnaryContext differentiates from other interfaces.
	IsUnaryContext()
}

type UnaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryContext() *UnaryContext {
	var p = new(UnaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_unary
	return p
}

func InitEmptyUnaryContext(p *UnaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_unary
}

func (*UnaryContext) IsUnaryContext() {}

func NewUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryContext {
	var p = new(UnaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BasicParserRULE_unary

	return p
}

func (s *UnaryContext) GetParser() antlr.Parser { return s.parser }
func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BasicVisitor:
		return t.VisitUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BasicParser) Unary() (localctx IUnaryContext) {
	localctx = NewUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BasicParserRULE_unary)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BasicParserT__14 || _la == BasicParserT__15) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMutlopContext is an interface to support dynamic dispatch.
type IMutlopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMutlopContext differentiates from other interfaces.
	IsMutlopContext()
}

type MutlopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMutlopContext() *MutlopContext {
	var p = new(MutlopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_mutlop
	return p
}

func InitEmptyMutlopContext(p *MutlopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_mutlop
}

func (*MutlopContext) IsMutlopContext() {}

func NewMutlopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MutlopContext {
	var p = new(MutlopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BasicParserRULE_mutlop

	return p
}

func (s *MutlopContext) GetParser() antlr.Parser { return s.parser }
func (s *MutlopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutlopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MutlopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BasicVisitor:
		return t.VisitMutlop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BasicParser) Mutlop() (localctx IMutlopContext) {
	localctx = NewMutlopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BasicParserRULE_mutlop)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BasicParserT__16 || _la == BasicParserT__17) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVaraContext is an interface to support dynamic dispatch.
type IVaraContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsVaraContext differentiates from other interfaces.
	IsVaraContext()
}

type VaraContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVaraContext() *VaraContext {
	var p = new(VaraContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_vara
	return p
}

func InitEmptyVaraContext(p *VaraContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_vara
}

func (*VaraContext) IsVaraContext() {}

func NewVaraContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VaraContext {
	var p = new(VaraContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BasicParserRULE_vara

	return p
}

func (s *VaraContext) GetParser() antlr.Parser { return s.parser }

func (s *VaraContext) VAR() antlr.TerminalNode {
	return s.GetToken(BasicParserVAR, 0)
}

func (s *VaraContext) STRING() antlr.TerminalNode {
	return s.GetToken(BasicParserSTRING, 0)
}

func (s *VaraContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VaraContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VaraContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BasicVisitor:
		return t.VisitVara(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BasicParser) Vara() (localctx IVaraContext) {
	localctx = NewVaraContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BasicParserRULE_vara)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BasicParserSTRING || _la == BasicParserVAR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDIGIT() []antlr.TerminalNode
	DIGIT(i int) antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BasicParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(BasicParserDIGIT)
}

func (s *NumberContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(BasicParserDIGIT, i)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BasicVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BasicParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BasicParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == BasicParserDIGIT {
		{
			p.SetState(126)
			p.Match(BasicParserDIGIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelopContext is an interface to support dynamic dispatch.
type IRelopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRelopContext differentiates from other interfaces.
	IsRelopContext()
}

type RelopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelopContext() *RelopContext {
	var p = new(RelopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_relop
	return p
}

func InitEmptyRelopContext(p *RelopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BasicParserRULE_relop
}

func (*RelopContext) IsRelopContext() {}

func NewRelopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelopContext {
	var p = new(RelopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BasicParserRULE_relop

	return p
}

func (s *RelopContext) GetParser() antlr.Parser { return s.parser }
func (s *RelopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BasicVisitor:
		return t.VisitRelop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BasicParser) Relop() (localctx IRelopContext) {
	localctx = NewRelopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BasicParserRULE_relop)
	var _la int

	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BasicParserT__18:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.Match(BasicParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BasicParserT__6 || _la == BasicParserT__19 {
			{
				p.SetState(132)
				_la = p.GetTokenStream().LA(1)

				if !(_la == BasicParserT__6 || _la == BasicParserT__19) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	case BasicParserT__19:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Match(BasicParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BasicParserT__6 || _la == BasicParserT__18 {
			{
				p.SetState(136)
				_la = p.GetTokenStream().LA(1)

				if !(_la == BasicParserT__6 || _la == BasicParserT__18) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	case BasicParserT__6:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(139)
			p.Match(BasicParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
