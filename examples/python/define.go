package main

import "github.com/x0y14/dylexer"

var DefIllegal = dylexer.Define{
	Kind:        Illegal,
	Regex:       `0^`,
	Consecutive: false,
}
var DefNewLine = dylexer.Define{
	Kind:        NewLine,
	Regex:       `^\n$`,
	Consecutive: true,
}
var DefTab = dylexer.Define{
	Kind:        Tab,
	Regex:       `^\t$`,
	Consecutive: true,
}
var DefWhite = dylexer.Define{
	Kind:        White,
	Regex:       `^ $`,
	Consecutive: true,
}

var DefColon = dylexer.Define{
	Kind:        Colon,
	Regex:       `^:$`,
	Consecutive: false,
}
var DefPeriod = dylexer.Define{
	Kind:        Period,
	Regex:       `^\.$`,
	Consecutive: false,
}
var DefLParent = dylexer.Define{
	Kind:        LParenthesis,
	Regex:       `^\($`,
	Consecutive: false,
}
var DefRParent = dylexer.Define{
	Kind:        RParenthesis,
	Regex:       `^\)$`,
	Consecutive: false,
}
var DefLSquare = dylexer.Define{
	Kind:        LSquareBracket,
	Regex:       `^\[$`,
	Consecutive: false,
}
var DefRSquare = dylexer.Define{
	Kind:        RSquareBracket,
	Regex:       `^\]$`,
	Consecutive: false,
}
var DefPlus = dylexer.Define{
	Kind:        Plus,
	Regex:       `^\+$`,
	Consecutive: false,
}
var DefMinus = dylexer.Define{
	Kind:        Minus,
	Regex:       `^\-$`,
	Consecutive: false,
}
var DefStar = dylexer.Define{
	Kind:        Star,
	Regex:       `^\*$`,
	Consecutive: false,
}
var DefAt = dylexer.Define{
	Kind:        At,
	Regex:       `^\@$`,
	Consecutive: false,
}
var DefSlash = dylexer.Define{
	Kind:        Slash,
	Regex:       `^\/$`,
	Consecutive: false,
}
var DefPercent = dylexer.Define{
	Kind:        Percent,
	Regex:       `^\%$`,
	Consecutive: false,
}
var DefEqual = dylexer.Define{
	Kind:        Equal,
	Regex:       `^\=$`,
	Consecutive: false,
}

var DefStringDq = dylexer.Define{
	Kind:        StringDq,
	Regex:       `^".*"$`,
	Consecutive: false,
}
var DefStringSq = dylexer.Define{
	Kind:        StringSq,
	Regex:       `^'.*'$`,
	Consecutive: false,
}
var DefNumeric = dylexer.Define{
	Kind:        Numeric,
	Regex:       `^-?[0-9\.]+$`,
	Consecutive: true,
}
var DefIdent = dylexer.Define{
	Kind:        Ident,
	Regex:       `^[A-Za-z_][A-Za-z0-9_]*$`,
	Consecutive: true,
}

var DefFunctionDef = dylexer.Define{
	Kind:        FunctionDef,
	Regex:       `^def$`,
	Consecutive: false,
}
var DefClassDef = dylexer.Define{
	Kind:        ClassDef,
	Regex:       `^class$`,
	Consecutive: false,
}
var DefImport = dylexer.Define{
	Kind:        Import,
	Regex:       `^import$`,
	Consecutive: false,
}
var DefNone = dylexer.Define{
	Kind:        None,
	Regex:       `^None$`,
	Consecutive: false,
}
var DefTrue = dylexer.Define{
	Kind:        True,
	Regex:       `^True$`,
	Consecutive: false,
}
var DefFalse = dylexer.Define{
	Kind:        False,
	Regex:       `^False$`,
	Consecutive: false,
}
