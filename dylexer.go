package dylexer

import (
	"go.uber.org/zap"
	"regexp"
)

type Lexer struct {
	logger              zap.SugaredLogger
	startPos            int
	pos                 int
	runes               []rune
	defines             []Define
	regex               []*regexp.Regexp
	tmp                 string
	consecutiveStartPos int
	consecutiveTmp      string
	consecutiveKind     TokenKind
}

func NewLexer(logger zap.SugaredLogger, defs []Define) *Lexer {
	return &Lexer{
		logger:              logger,
		startPos:            0,
		pos:                 0,
		runes:               nil,
		defines:             defs,
		tmp:                 "",
		consecutiveStartPos: 0,
		consecutiveTmp:      "",
		consecutiveKind:     nil,
	}
}

func (l *Lexer) CompileRegex() error {
	for _, def := range l.defines {
		r, err := regexp.Compile(def.Regex)
		if err != nil {
			return err
		}
		l.regex = append(l.regex, r)
		l.logger.Debugf("compiled:%10v:%20v", def.Kind.String(), def.Regex)
	}

	return nil
}

func (l *Lexer) Tokenize(subject string) []Token {
	var result []Token

	l.runes = []rune(subject)
	for !l.isEof() {
		tmpPos := l.pos
		l.consumeAndAttachTmp()
		def, ok := l.match()
		if l.consecutiveKind != nil && l.consecutiveKind != def.Kind {
			//fmt.Printf("[%d-%d]%20s: %s\n", l.consecutiveStartPos, tmpPos, l.consecutiveTmp, l.consecutiveKind.String())
			result = append(result, Token{
				TokenKind: l.consecutiveKind,
				Raw:       l.consecutiveTmp,
				S:         l.consecutiveStartPos,
				E:         tmpPos,
			})
			l.consecutiveTmp = ""
			l.consecutiveKind = nil
		}

		if ok {
			if !def.AllowJoin {
				//fmt.Printf("[%d-%d]%20s: %s\n", l.startPos, l.pos, l.tmp, def.Kind.String())
				result = append(result, Token{
					TokenKind: def.Kind,
					Raw:       l.tmp,
					S:         l.startPos,
					E:         l.pos,
				})
				l.consecutiveStartPos = l.pos
			} else {
				l.consecutiveKind = def.Kind
				l.consecutiveTmp += l.tmp
			}
			l.resetTmp()
			l.startPos = l.pos
		}
	}

	return result
}

func (l *Lexer) isEof() bool {
	return l.pos >= len(l.runes)
}

func (l *Lexer) curt() rune {
	return l.runes[l.pos]
}

func (l *Lexer) goNext() {
	l.pos++
}

func (l *Lexer) consumeAndAttachTmp() {
	c := l.curt()
	l.goNext()
	l.tmp += string(c)
}

func (l *Lexer) resetTmp() {
	l.tmp = ""
}

func (l *Lexer) match() (Define, bool) {
	for i, reg := range l.regex {
		if reg.MatchString(l.tmp) {
			return l.defines[i], true
		}
	}

	return l.defines[0], false
}
