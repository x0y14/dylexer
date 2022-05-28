package dylexer

type Define struct {
	Kind        TokenKind
	Regex       string
	Consecutive bool
}
