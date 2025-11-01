package main

import (
	"Language/cli/repl"
	"os"
)

func main() {

	repl.Start(os.Stdin, os.Stdout)

	// if len(os.Args) < 2 {
	// 	fmt.Println("Error: no input provided")
	// 	return
	// }

	// src, err := os.ReadFile(os.Args[1])
	// if err != nil {
	// 	fmt.Println("Error: couldn't open file")
	// 	panic(err)
	// }

	// l := lexer.NewLexer(string(src))
	// tokens := l.Tokenize()

	// p := parser.NewParser(tokens)
	// program := p.ParseProgram()

}
