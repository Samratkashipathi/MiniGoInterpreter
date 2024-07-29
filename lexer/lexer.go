package lexer

import (
	"interpreter/token"
)

// TODO: Current implementation of lexer does not support unicode, only support ASCII code
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition = l.readPosition + 1
}

func (l Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpaces()

	if l.ch == 0 {
		tok.Type = token.EOF
		tok.Literal = ""
		return tok
	}

	// Check for character
	if tk, ok := token.Symbols[string(l.ch)]; ok {

		if l.ch == '=' && l.peekChar() == '=' {
			l.readChar()
			l.readChar()
			tok.Type = token.EQ
			tok.Literal = "=="
			return tok
		}

		if l.ch == '!' && l.peekChar() == '=' {
			l.readChar()
			l.readChar()
			tok.Type = token.NOT_EQ
			tok.Literal = "!="
			return tok
		}

		tok = newToken(tk, l.ch)
		l.readChar()
		return tok
	}

	// If it is a letter check for keywords else normal identifier
	if isLetter(l.ch) {
		tok.Literal = l.readIdentifier()

		if keyWordTokenType, ok := token.Keywords[tok.Literal]; ok {
			tok.Type = keyWordTokenType
			return tok
		}

		tok.Type = token.IDENT
		return tok
	}

	// If it is a digit return INT
	if isDigit(l.ch) {
		tok.Literal = l.readNumber()
		tok.Type = token.INT

		return tok
	}

	tok = newToken(token.ILLEGAL, l.ch)
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

// TODO: Currently supports only Integer
func (l *Lexer) readNumber() string {
	startPosition := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

func (l *Lexer) skipWhiteSpaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
