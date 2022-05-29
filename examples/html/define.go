package main

import "github.com/x0y14/dylexer"

var DefIllegal = dylexer.Define{
	Kind:      Illegal,
	Regex:     `0^`,
	AllowJoin: false,
}
var DefIdent = dylexer.Define{
	Kind:      Ident,
	Regex:     `^[/A-Za-z0-9_,]$`,
	AllowJoin: true,
}
var DefEqual = dylexer.Define{
	Kind:      Equal,
	Regex:     `^=$`,
	AllowJoin: false,
}
var DefString = dylexer.Define{
	Kind:      String,
	Regex:     `^".*"$`,
	AllowJoin: false,
}
var DefStartTag = dylexer.Define{
	Kind:      StartTag,
	Regex:     `^<$`,
	AllowJoin: false,
}
var DefEndTag = dylexer.Define{
	Kind:      EndTag,
	Regex:     `^>$`,
	AllowJoin: false,
}
var DefWhite = dylexer.Define{
	Kind:      White,
	Regex:     `^\s$`,
	AllowJoin: true,
}
