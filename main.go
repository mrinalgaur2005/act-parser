package main

import (
	"fmt"
	"os"

	"github.com/mrinalgaur2005/act-parser/javascript/lexer"
	"github.com/mrinalgaur2005/act-parser/javascript/parser"
	"github.com/sanity-io/litter"
)

func main() {
	fmt.Println("hello")

	bytes, _ := os.ReadFile("javascript/examples/04.js")
	source := string(bytes)

	tokens := lexer.Tokenize(source)

	// for _, token := range tokens {
	// 	token.Debug()
	// }

	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
