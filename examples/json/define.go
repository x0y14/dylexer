package main

import "github.com/x0y14/dylexer"

var DefIllegal = dylexer.Define{
	Kind:      Illegal,
	Regex:     "0^",
	AllowJoin: false,
}
var DefWhite = dylexer.Define{
	Kind:      White,
	Regex:     `^\s$`,
	AllowJoin: false,
}
var DefLCurlyBracket = dylexer.Define{
	Kind:      LCurlyBracket,
	Regex:     `^{$`,
	AllowJoin: false,
}
var DefRCurlyBracket = dylexer.Define{
	Kind:      RCurlyBracket,
	Regex:     `^}$`,
	AllowJoin: false,
}
var DefColon = dylexer.Define{
	Kind:      Colon,
	Regex:     `^:$`,
	AllowJoin: false,
}
var DefComma = dylexer.Define{
	Kind:      RCurlyBracket,
	Regex:     `^,$`,
	AllowJoin: false,
}
var DefString = dylexer.Define{
	Kind:      String,
	Regex:     `^".*"$`,
	AllowJoin: false,
}
var DefNumber = dylexer.Define{
	Kind:      Number,
	Regex:     `^-?[0-9\.]+$`,
	AllowJoin: true,
}
