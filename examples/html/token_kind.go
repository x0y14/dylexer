package main

type TokenKind int

const (
	Illegal  TokenKind = iota
	Ident              //
	Equal              // =
	String             // "abc"
	StartTag           // <
	EndTag             // >
	White
)

var kinds = [...]string{
	Illegal:  "Illegal",
	White:    "White",
	Ident:    "Ident",
	Equal:    "Equal",
	String:   "String",
	StartTag: "StartTag",
	EndTag:   "EndTag",
}

func (t TokenKind) String() string {
	return kinds[t]
}
