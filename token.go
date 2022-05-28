package dylexer

type Token struct {
	TokenKind
	Raw string
	S   int
	E   int
}
