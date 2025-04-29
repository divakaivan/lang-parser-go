package main

import (
	"os"

	"github.com/divakaivan/lang-parser-go/src/lexer"
)

func main() {
	bytes, _ := os.ReadFile("./examples/02.lang")
	tokens := lexer.Tokenize(string(bytes))

	for _, token := range tokens {
		token.Debug()
	}
}
