package main

import (
	"fmt"
	"os"

	"github.com/mrinalgaur2005/act-parser/javascript"
)

func main() {
	fmt.Println("hello")

	bytes, _ := os.ReadFile("javascript/examples/00.js")
	source := string(bytes)

	tokens := javascript.Tokenize(source)

	for _, token := range tokens {
		token.Debug()
	}
}
