package dylexer

import (
	"go.uber.org/zap"
	"testing"
)

type testTokenKind int

var kinds []string

func (t testTokenKind) String() string {
	return kinds[t]
}
func newTestTokenKind(alt string) TokenKind {
	var a = testTokenKind(len(kinds))
	kinds = append(kinds, alt)
	return a
}

func TestLexer_CompileRegex(t *testing.T) {

}

func TestLexer_Tokenize(t *testing.T) {

	var defs []Define
	defs = append(defs, Define{
		Kind:  newTestTokenKind("Unknown"),
		Regex: `0^`,
	})
	defs = append(defs, Define{
		Kind:  newTestTokenKind("LeftCurlyBracket"),
		Regex: `^{$`,
	})
	defs = append(defs, Define{
		Kind:  newTestTokenKind("RightCurlyBracket"),
		Regex: `^}$`,
	})
	defs = append(defs, Define{
		Kind:  newTestTokenKind("Null"),
		Regex: `^null$`,
	})

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatal(err)
	}
	lexer := NewLexer(*logger.Sugar(), defs)
	if err := lexer.CompileRegex(); err != nil {
		t.Fatal(err)
	}
	lexer.Tokenize("{}")
}
