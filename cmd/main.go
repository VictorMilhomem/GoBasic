package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/VictorMilhomem/Basic/cmd/compiler"
	"github.com/VictorMilhomem/Basic/cmd/parser"
	"github.com/antlr4-go/antlr/v4"
)

func stringToLlvm(source string) error {
	// Define the file path
	outputPath := filepath.Join("bin", "out.ll")

	// Convert the string to bytes
	data := []byte(source)

	// Create the bin directory if it doesn't exist
	if err := os.MkdirAll("bin", os.ModePerm); err != nil {
		return err
	}

	// Create the file and write the bytes
	err := ioutil.WriteFile(outputPath, data, 0o644)
	if err != nil {
		return err
	}

	return nil
}

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
	cpl := compiler.NewCompiler()
	cpl.Set()
	visitor := parser.NewVisitor(cpl)
	cpl.Main()
	_ = visitor.Visit(tree)
	cpl.MainRet()
	fmt.Println(cpl.Module)
	err = stringToLlvm(cpl.Module.String())
	if err != nil {
		log.Panic("could not create ll file")
	}
}

func main() {
	// args := os.Args[1:]
	runFile("examples\\hello.bas")
	/* if len(args) > 1 {
		fmt.Println("Usage: basic [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile("examples\\paint.bas")
	} else {
		os.Exit(64)
	} */
}
