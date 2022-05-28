package main

import (
	"fmt"
	"github.com/x0y14/dylexer"
)

func main() {
	defines := []dylexer.Define{
		DefIllegal,
		DefIdent,
		DefEqual,
		DefString,
		DefStartTag,
		DefEndTag,
		DefWhite,
	}
	lexer := dylexer.NewLexer(defines)
	if err := lexer.CompileRegex(); err != nil {
		panic(err)
	}
	tokens := lexer.Tokenize("<div>hello, world</div>")
	for _, tok := range tokens {
		fmt.Println(tok.String())
	}
}
