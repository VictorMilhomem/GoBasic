package main

import (
	"io/ioutil"
	"log"
	"path"

	"github.com/VictorMilhomem/Basic/cmd/parser"
	"github.com/antlr4-go/antlr/v4"
)

func runFile(fpath string) {
	bytes, err := ioutil.ReadFile(path.Base(fpath))
	source := string(bytes)
	if err != nil {
		log.Fatalf("could not open file %s", fpath)
		return
	}

	input := antlr.NewInputStream(source)
	lexer := parser.NewBasicLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewBasicParser(stream)
	p.BuildParseTrees = true
	tree := p.Prog()
	visitor := parser.Visitor{}
	_ = visitor.VisitProg(tree.(*parser.ProgContext))
}

func main() {
	// args := os.Args[1:]
	runFile("examples\\paint.bas")
	/* if len(args) > 1 {
		fmt.Println("Usage: basic [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile("examples\\paint.bas")
	} else {
		os.Exit(64)
	} */
}
