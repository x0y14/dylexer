package dylexer

import "fmt"

type Token struct {
	TokenKind
	Raw string
	S   int
	E   int
}

func (t Token) String() string {
	return fmt.Sprintf("%s(%d-%d) %v", t.TokenKind.String(), t.S, t.E, t.Raw)
}
