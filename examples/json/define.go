package json

import "github.com/x0y14/dylexer"

var DefIllegal = dylexer.Define{
	Kind:        Illegal,
	Regex:       "0^",
	Consecutive: false,
}
var DefWhite = dylexer.Define{
	Kind:        White,
	Regex:       `^\s$`,
	Consecutive: false,
}
var DefLCurlyBracket = dylexer.Define{
	Kind:        LCurlyBracket,
	Regex:       `^{$`,
	Consecutive: false,
}
var DefRCurlyBracket = dylexer.Define{
	Kind:        RCurlyBracket,
	Regex:       `^}$`,
	Consecutive: false,
}
var DefColon = dylexer.Define{
	Kind:        Colon,
	Regex:       `^:$`,
	Consecutive: false,
}
var DefComma = dylexer.Define{
	Kind:        RCurlyBracket,
	Regex:       `^,$`,
	Consecutive: false,
}
var DefString = dylexer.Define{
	Kind:        String,
	Regex:       `^".*"$`,
	Consecutive: false,
}

var DefNumber = dylexer.Define{
	Kind:        Number,
	Regex:       `^-?[0-9\.]+$`,
	Consecutive: true,
}
