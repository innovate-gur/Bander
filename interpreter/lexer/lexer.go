package lexer

import "github.com/Noreen-py/Bander/interpreter/token"

type Lexer struct {
	source   string
	position int
	current  int
	ch       byte
}

func Lex(input string) *Lexer {
	l := &Lexer{source: input}
	l.Advance()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var ret token.Token
	l.Space()
	switch l.ch {

	case 0:
		ret.Literal = ""
		ret.Type = token.EOF
	default:
		if IsLetter(l.ch) {
			ret.Literal = l.Identifier()
			ret.Type = token.LookUpIdent(ret.Literal)
			return ret
		} else if IsDigit(l.ch) {
			ret.Literal = l.Number()
			ret.Type = token.Number
			return ret
		} else {
			ret = AddToken(token.Undefined, l.ch)

		}

	}
	l.Advance()
	return ret
}

func AddToken(tokentype token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokentype, Literal: string(ch)}

}

func (l *Lexer) Advance() {
	if l.current >= len(l.source) {
		l.ch = 0
	} else {
		l.ch = l.source[l.current]
	}
	l.position = l.current
	l.current += 1
}

func (l *Lexer) Identifier() string {
	fix := l.position
	for IsLetter(l.ch) {
		l.Advance()
	}
	return l.source[fix:l.position]
}

func (l *Lexer) Number() string {
	fix := l.position
	for IsDigit(l.ch) {
		l.Advance()
	}
	return l.source[fix:l.position]
}

func IsLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) Space() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.Advance()
	}
}
