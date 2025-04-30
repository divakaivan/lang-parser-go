package main

import (
	"os"

	"github.com/divakaivan/lang-parser-go/src/lexer"
	"github.com/divakaivan/lang-parser-go/src/parser"
	"github.com/sanity-io/litter"
)

func main() {
	bytes, _ := os.ReadFile("./examples/03.lang")
	tokens := lexer.Tokenize(string(bytes))

	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
