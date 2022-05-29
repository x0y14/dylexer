package main

type TokenKind int

const (
	Illegal TokenKind = iota
	NewLine           // \n
	Tab               // \t
	White             // " "

	Colon          // :
	Period         // .
	LParenthesis   // (
	RParenthesis   // )
	LSquareBracket // [
	RSquareBracket // ]
	Plus           // +
	Minus          // -
	Star           // *
	At             // @
	Slash          // /
	Percent        // %
	Equal          // =

	StringDq // "aa"
	StringSq // 'aa'
	Numeric  // -12.3
	Ident    //

	FunctionDef // def
	ClassDef    // class
	Import      // import
	None        //
	True        //
	False       //
)

var kinds = [...]string{
	Illegal: "Illegal",
	NewLine: "NewLine",
	Tab:     "Tab",
	White:   "White",

	Colon:          "Colon",
	Period:         "Period",
	LParenthesis:   "LParenthesis",
	RParenthesis:   "RParenthesis",
	LSquareBracket: "LSquareBracket",
	RSquareBracket: "RSquareBracket",
	Plus:           "Plus",
	Minus:          "Minus",
	Star:           "Star",
	At:             "At",
	Slash:          "Slash",
	Percent:        "Percent",
	Equal:          "Equal",

	StringDq: "StringDq",
	StringSq: "StringSq",
	Numeric:  "Numeric",
	Ident:    "Ident",

	FunctionDef: "FunctionDef",
	ClassDef:    "ClassDef",
	Import:      "Import",
	None:        "None",
	True:        "True",
	False:       "False",
}

func (t TokenKind) String() string {
	return kinds[t]
}
