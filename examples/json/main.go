package main

import (
	"fmt"
	"github.com/x0y14/dylexer"
	"go.uber.org/zap"
)

func main() {
	defines := []dylexer.Define{
		DefIllegal,
		DefWhite,
		DefLCurlyBracket,
		DefRCurlyBracket,
		DefColon,
		DefComma,
		DefString,
		DefNumber,
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	lexer := dylexer.NewLexer(*logger.Sugar(), defines)
	if err := lexer.CompileRegex(); err != nil {
		panic(err)
	}
	tokens := lexer.Tokenize(`{"Key_String":"Value", "Key_Number1": -12.3, "Key_Number2": 1200}`)
	for _, tok := range tokens {
		fmt.Println(tok.String())
	}
}
