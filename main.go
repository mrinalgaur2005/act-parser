package main

import (
	"fmt"
	"os"

	"github.com/mrinalgaur2005/act-parser/js/ts/lexer"
	"github.com/mrinalgaur2005/act-parser/js/ts/parser"
	"github.com/sanity-io/litter"
)

func main() {
	fmt.Println("hello")

	bytes, _ := os.ReadFile("js/ts/examples/05.ts")
	source := string(bytes)

	tokens := lexer.Tokenize(source)

	// for _, token := range tokens {
	// 	token.Debug()
	// }

	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
