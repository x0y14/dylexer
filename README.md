dylexer  
正規表現に従って文字列を分解するレクサ

### usage

1. 種類を定義してください
```go
type TokenKind int
const (
    Illegal TokenKind = iota
    //
)
```

2. Defineを作ってください(任意)  
`AllowJoin` 数字など連続している可能性があるが、単体でも意味を持つもので、連続して出現した場合それを結合したい場合にtrueしてください  
例えばfalseにした場合、数字120は、1、2、0として独立して処理されますが、trueにした場合、120と扱われます。
```go
var DefIllegal = dylexer.Define{
	Kind:      Illegal,
	Regex:     `0^`,
	AllowJoin: false, // when continuous
}
```

3. レクサに登録して、コンパイルしてください
```go
lexer := dylexer.NewLexer(defs)
if err := lexer.CompileRegex(); err != nil {
    panic(err)
}
```

4. トークナイズしてください
```go
tokens := lexer.Tokenize("text here")
```