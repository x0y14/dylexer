package json

import (
	"fmt"
	"github.com/x0y14/dylexer"
	"testing"
)

func Test(t *testing.T) {
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
		t.Fatal(err)
	}
	tokens := lexer.Tokenize(`{"Key_String":"Value", "Key_Number1": -12.3, "Key_Number2": 1200}`)
	for _, tok := range tokens {
		fmt.Println(tok.String())
	}
}
