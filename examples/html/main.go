package main

import (
	"fmt"
	"github.com/x0y14/dylexer"
	"go.uber.org/zap"
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
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	lexer := dylexer.NewLexer(*logger.Sugar(), defines)
	if err := lexer.CompileRegex(); err != nil {
		panic(err)
	}
	tokens := lexer.Tokenize("<div>hello, world</div>")
	for _, tok := range tokens {
		fmt.Println(tok.String())
	}
}
