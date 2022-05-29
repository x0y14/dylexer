package main

import (
	"fmt"
	"github.com/x0y14/dylexer"
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
	lexer := dylexer.NewLexer(defines)
	if err := lexer.CompileRegex(); err != nil {
		panic(err)
	}
	tokens := lexer.Tokenize(`{"Key_String":"Value", "Key_Number1": -12.3, "Key_Number2": 1200}`)
	for _, tok := range tokens {
		fmt.Println(tok.String())
	}
}
