// Code generated from ./cmd/parser/Basic.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type BasicLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var BasicLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func basiclexerLexerInit() {
	staticData := &BasicLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "STRING", "DIGIT", "VAR", "CR", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 150, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 5, 20, 131, 8,
		20, 10, 20, 12, 20, 134, 9, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 23, 4, 23, 143, 8, 23, 11, 23, 12, 23, 144, 1, 24, 1, 24, 1, 24, 1,
		24, 0, 0, 25, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18,
		37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 1, 0, 3, 3, 0,
		10, 10, 13, 13, 34, 34, 2, 0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 151,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 1, 51, 1, 0, 0, 0, 3, 57, 1, 0, 0, 0,
		5, 60, 1, 0, 0, 0, 7, 65, 1, 0, 0, 0, 9, 70, 1, 0, 0, 0, 11, 76, 1, 0,
		0, 0, 13, 80, 1, 0, 0, 0, 15, 82, 1, 0, 0, 0, 17, 88, 1, 0, 0, 0, 19, 95,
		1, 0, 0, 0, 21, 101, 1, 0, 0, 0, 23, 106, 1, 0, 0, 0, 25, 110, 1, 0, 0,
		0, 27, 114, 1, 0, 0, 0, 29, 116, 1, 0, 0, 0, 31, 118, 1, 0, 0, 0, 33, 120,
		1, 0, 0, 0, 35, 122, 1, 0, 0, 0, 37, 124, 1, 0, 0, 0, 39, 126, 1, 0, 0,
		0, 41, 128, 1, 0, 0, 0, 43, 137, 1, 0, 0, 0, 45, 139, 1, 0, 0, 0, 47, 142,
		1, 0, 0, 0, 49, 146, 1, 0, 0, 0, 51, 52, 5, 80, 0, 0, 52, 53, 5, 82, 0,
		0, 53, 54, 5, 73, 0, 0, 54, 55, 5, 78, 0, 0, 55, 56, 5, 84, 0, 0, 56, 2,
		1, 0, 0, 0, 57, 58, 5, 73, 0, 0, 58, 59, 5, 70, 0, 0, 59, 4, 1, 0, 0, 0,
		60, 61, 5, 84, 0, 0, 61, 62, 5, 72, 0, 0, 62, 63, 5, 69, 0, 0, 63, 64,
		5, 78, 0, 0, 64, 6, 1, 0, 0, 0, 65, 66, 5, 71, 0, 0, 66, 67, 5, 79, 0,
		0, 67, 68, 5, 84, 0, 0, 68, 69, 5, 79, 0, 0, 69, 8, 1, 0, 0, 0, 70, 71,
		5, 73, 0, 0, 71, 72, 5, 78, 0, 0, 72, 73, 5, 80, 0, 0, 73, 74, 5, 85, 0,
		0, 74, 75, 5, 84, 0, 0, 75, 10, 1, 0, 0, 0, 76, 77, 5, 76, 0, 0, 77, 78,
		5, 69, 0, 0, 78, 79, 5, 84, 0, 0, 79, 12, 1, 0, 0, 0, 80, 81, 5, 61, 0,
		0, 81, 14, 1, 0, 0, 0, 82, 83, 5, 71, 0, 0, 83, 84, 5, 79, 0, 0, 84, 85,
		5, 83, 0, 0, 85, 86, 5, 85, 0, 0, 86, 87, 5, 66, 0, 0, 87, 16, 1, 0, 0,
		0, 88, 89, 5, 82, 0, 0, 89, 90, 5, 69, 0, 0, 90, 91, 5, 84, 0, 0, 91, 92,
		5, 85, 0, 0, 92, 93, 5, 82, 0, 0, 93, 94, 5, 78, 0, 0, 94, 18, 1, 0, 0,
		0, 95, 96, 5, 67, 0, 0, 96, 97, 5, 76, 0, 0, 97, 98, 5, 69, 0, 0, 98, 99,
		5, 65, 0, 0, 99, 100, 5, 82, 0, 0, 100, 20, 1, 0, 0, 0, 101, 102, 5, 76,
		0, 0, 102, 103, 5, 73, 0, 0, 103, 104, 5, 83, 0, 0, 104, 105, 5, 84, 0,
		0, 105, 22, 1, 0, 0, 0, 106, 107, 5, 82, 0, 0, 107, 108, 5, 85, 0, 0, 108,
		109, 5, 78, 0, 0, 109, 24, 1, 0, 0, 0, 110, 111, 5, 69, 0, 0, 111, 112,
		5, 78, 0, 0, 112, 113, 5, 68, 0, 0, 113, 26, 1, 0, 0, 0, 114, 115, 5, 44,
		0, 0, 115, 28, 1, 0, 0, 0, 116, 117, 5, 43, 0, 0, 117, 30, 1, 0, 0, 0,
		118, 119, 5, 45, 0, 0, 119, 32, 1, 0, 0, 0, 120, 121, 5, 42, 0, 0, 121,
		34, 1, 0, 0, 0, 122, 123, 5, 47, 0, 0, 123, 36, 1, 0, 0, 0, 124, 125, 5,
		60, 0, 0, 125, 38, 1, 0, 0, 0, 126, 127, 5, 62, 0, 0, 127, 40, 1, 0, 0,
		0, 128, 132, 5, 34, 0, 0, 129, 131, 8, 0, 0, 0, 130, 129, 1, 0, 0, 0, 131,
		134, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 135,
		1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 135, 136, 5, 34, 0, 0, 136, 42, 1, 0,
		0, 0, 137, 138, 2, 48, 57, 0, 138, 44, 1, 0, 0, 0, 139, 140, 2, 65, 90,
		0, 140, 46, 1, 0, 0, 0, 141, 143, 7, 1, 0, 0, 142, 141, 1, 0, 0, 0, 143,
		144, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 48, 1,
		0, 0, 0, 146, 147, 7, 2, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 6, 24, 0,
		0, 149, 50, 1, 0, 0, 0, 3, 0, 132, 144, 1, 6, 0, 0,
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

// BasicLexerInit initializes any static state used to implement BasicLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewBasicLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BasicLexerInit() {
	staticData := &BasicLexerLexerStaticData
	staticData.once.Do(basiclexerLexerInit)
}

// NewBasicLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewBasicLexer(input antlr.CharStream) *BasicLexer {
	BasicLexerInit()
	l := new(BasicLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &BasicLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Basic.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BasicLexer tokens.
const (
	BasicLexerT__0   = 1
	BasicLexerT__1   = 2
	BasicLexerT__2   = 3
	BasicLexerT__3   = 4
	BasicLexerT__4   = 5
	BasicLexerT__5   = 6
	BasicLexerT__6   = 7
	BasicLexerT__7   = 8
	BasicLexerT__8   = 9
	BasicLexerT__9   = 10
	BasicLexerT__10  = 11
	BasicLexerT__11  = 12
	BasicLexerT__12  = 13
	BasicLexerT__13  = 14
	BasicLexerT__14  = 15
	BasicLexerT__15  = 16
	BasicLexerT__16  = 17
	BasicLexerT__17  = 18
	BasicLexerT__18  = 19
	BasicLexerT__19  = 20
	BasicLexerSTRING = 21
	BasicLexerDIGIT  = 22
	BasicLexerVAR    = 23
	BasicLexerCR     = 24
	BasicLexerWS     = 25
)
