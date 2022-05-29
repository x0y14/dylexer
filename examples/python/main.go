package main

import (
	"fmt"
	"github.com/x0y14/dylexer"
	"os"
)

func main() {
	bytes, err := os.ReadFile("snapshot.py")
	if err != nil {
		panic(err)
	}

	defs := []dylexer.Define{
		DefIllegal,

		DefNewLine,
		DefTab,
		DefWhite,

		DefColon,   // :
		DefPeriod,  // .
		DefLParent, // (
		DefRParent, // )
		DefLSquare, // [
		DefRSquare, // ]
		DefPlus,    // +
		DefMinus,   // -
		DefStar,    // *
		DefAt,      // @
		DefSlash,   // /
		DefPercent, // %
		DefEqual,   // =

		// 上の方が優先度が高い。
		// IdentとKeywordは被るので、Keywordを上に持ってきてる。

		DefFunctionDef,
		DefClassDef,
		DefImport,
		DefNone,
		DefTrue,
		DefFalse,

		DefStringDq, // "abc"
		DefStringSq, // 'abc'
		DefNumeric,  // 12
		DefIdent,    // abc
	}

	lexer := dylexer.NewLexer(defs)
	if err := lexer.CompileRegex(); err != nil {
		panic(err)
	}
	tokens := lexer.Tokenize(string(bytes))
	for _, tok := range tokens {
		fmt.Println(tok.String())
	}
}
