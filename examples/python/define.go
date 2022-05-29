package main

import "github.com/x0y14/dylexer"

var DefIllegal = dylexer.Define{
	Kind:      Illegal,
	Regex:     `0^`,
	AllowJoin: false,
}
var DefNewLine = dylexer.Define{
	Kind:      NewLine,
	Regex:     `^\n$`,
	AllowJoin: true,
}
var DefTab = dylexer.Define{
	Kind:      Tab,
	Regex:     `^\t$`,
	AllowJoin: true,
}
var DefWhite = dylexer.Define{
	Kind:      White,
	Regex:     `^ $`,
	AllowJoin: true,
}

var DefColon = dylexer.Define{
	Kind:      Colon,
	Regex:     `^:$`,
	AllowJoin: false,
}
var DefPeriod = dylexer.Define{
	Kind:      Period,
	Regex:     `^\.$`,
	AllowJoin: false,
}
var DefComma = dylexer.Define{
	Kind:      Comma,
	Regex:     `^,$`,
	AllowJoin: false,
}
var DefLParent = dylexer.Define{
	Kind:      LParenthesis,
	Regex:     `^\($`,
	AllowJoin: false,
}
var DefRParent = dylexer.Define{
	Kind:      RParenthesis,
	Regex:     `^\)$`,
	AllowJoin: false,
}
var DefLSquare = dylexer.Define{
	Kind:      LSquareBracket,
	Regex:     `^\[$`,
	AllowJoin: false,
}
var DefRSquare = dylexer.Define{
	Kind:      RSquareBracket,
	Regex:     `^\]$`,
	AllowJoin: false,
}
var DefLCurly = dylexer.Define{
	Kind:      LCurlyBracket,
	Regex:     `^\{$`,
	AllowJoin: false,
}
var DefRCurly = dylexer.Define{
	Kind:      RCurlyBracket,
	Regex:     `^\}$`,
	AllowJoin: false,
}
var DefPlus = dylexer.Define{
	Kind:      Plus,
	Regex:     `^\+$`,
	AllowJoin: false,
}
var DefMinus = dylexer.Define{
	Kind:      Minus,
	Regex:     `^\-$`,
	AllowJoin: false,
}
var DefStar = dylexer.Define{
	Kind:      Star,
	Regex:     `^\*$`,
	AllowJoin: false,
}
var DefAt = dylexer.Define{
	Kind:      At,
	Regex:     `^\@$`,
	AllowJoin: false,
}
var DefSlash = dylexer.Define{
	Kind:      Slash,
	Regex:     `^\/$`,
	AllowJoin: false,
}
var DefPercent = dylexer.Define{
	Kind:      Percent,
	Regex:     `^\%$`,
	AllowJoin: false,
}
var DefEqual = dylexer.Define{
	Kind:      Equal,
	Regex:     `^\=$`,
	AllowJoin: false,
}

var DefStringDq = dylexer.Define{
	Kind:      StringDq,
	Regex:     `^".*"$`,
	AllowJoin: false,
}
var DefStringSq = dylexer.Define{
	Kind:      StringSq,
	Regex:     `^'.*'$`,
	AllowJoin: false,
}
var DefNumeric = dylexer.Define{
	Kind:      Numeric,
	Regex:     `^-?[0-9\.]+$`,
	AllowJoin: true,
}
var DefIdent = dylexer.Define{
	Kind:      Ident,
	Regex:     `^[A-Za-z_][A-Za-z0-9_]*$`,
	AllowJoin: true,
}

var DefFunctionDef = dylexer.Define{
	Kind:      FunctionDef,
	Regex:     `^def$`,
	AllowJoin: false,
}
var DefClassDef = dylexer.Define{
	Kind:      ClassDef,
	Regex:     `^class$`,
	AllowJoin: false,
}
var DefImport = dylexer.Define{
	Kind:      Import,
	Regex:     `^import$`,
	AllowJoin: false,
}
var DefNone = dylexer.Define{
	Kind:      None,
	Regex:     `^None$`,
	AllowJoin: false,
}
var DefTrue = dylexer.Define{
	Kind:      True,
	Regex:     `^True$`,
	AllowJoin: false,
}
var DefFalse = dylexer.Define{
	Kind:      False,
	Regex:     `^False$`,
	AllowJoin: false,
}
