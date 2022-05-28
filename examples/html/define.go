package main

import "github.com/x0y14/dylexer"

var DefIllegal = dylexer.Define{
	Kind:        Illegal,
	Regex:       `0^`,
	Consecutive: false,
}
var DefIdent = dylexer.Define{
	Kind:        Ident,
	Regex:       `^[/A-Za-z0-9_,]$`,
	Consecutive: true,
}
var DefEqual = dylexer.Define{
	Kind:        Equal,
	Regex:       `^=$`,
	Consecutive: false,
}
var DefString = dylexer.Define{
	Kind:        String,
	Regex:       `^".*"$`,
	Consecutive: false,
}
var DefStartTag = dylexer.Define{
	Kind:        StartTag,
	Regex:       `^<$`,
	Consecutive: false,
}
var DefEndTag = dylexer.Define{
	Kind:        EndTag,
	Regex:       `^>$`,
	Consecutive: false,
}
var DefWhite = dylexer.Define{
	Kind:        White,
	Regex:       `^\s$`,
	Consecutive: true,
}
