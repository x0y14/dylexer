package json

import (
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
	lexer.Tokenize(`{"Key_String":"Value", "Key_Number1": -12.3, "Key_Number2": 1200}`)
}
