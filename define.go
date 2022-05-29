package dylexer

type Define struct {
	Kind      TokenKind
	Regex     string
	AllowJoin bool // when continuous
}
