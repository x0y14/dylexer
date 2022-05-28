package json

type TokenKind int

const (
	Illegal       TokenKind = iota
	LCurlyBracket           // {
	RCurlyBracket           // }
	DQuotation              // "
	Colon                   // :
	Comma                   // ,

	LSquareBracket // [
	RSquareBracket // ]

	String // "abc"
	Number // 123
	True
	False
	Null

	White //
)

var kinds = [...]string{
	Illegal:       "Illegal",
	LCurlyBracket: "LCurlyBracket",
	RCurlyBracket: "RCurlyBracket",
	DQuotation:    "DQuotation",
	Colon:         "Colon",
	Comma:         "Comma",

	LSquareBracket: "LSquareBracket",
	RSquareBracket: "RSquareBracket",

	String: "String",
	Number: "Number",
	True:   "True",
	False:  "False",
	Null:   "Null",

	White: "White",
}

func (t TokenKind) String() string {
	return kinds[t]
}
