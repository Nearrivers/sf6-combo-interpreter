package lexer

import (
	"slices"

	"github.com/Nearrivers/combo-parser/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
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
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '.'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func (l *Lexer) skipWhitespace() {
	for isWhiteSpace(l.ch) {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readIdentifierUntilChar(chars ...byte) string {
	position := l.position
	for l.readPosition < len(l.input) {
		l.readChar()
		if slices.Contains(chars, l.ch) {
			l.readChar()
			break
		}
	}

	return l.input[position:l.position]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '>':
		tok = newToken(token.GT, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '~':
		tok = newToken(token.TILDA, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '(':
		// On veut lire que jusqu'aux parenthèses fermantes
		tok.Literal = l.readIdentifierUntilChar(')')
		// (CA) pour critical art
		if l.peekChar() == 'C' {
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}

		tok.Type = token.UNKOWN
		return tok
	case '[':
		// [] pour [CH] ou [PC] ou indiquer qu'il faut maintenir l'input
		// ou le bouton
		tok.Literal = l.readIdentifierUntilChar(']')
		if token.IsIdent(tok.Literal) {
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}

		// [2] par exemple pour maintenir bas
		if isDigit(l.peekChar()) {
			tok.Type = token.INPUT
			return tok
		}

		// [HK] par exemple pour maitenir Heavy Kick
		if isLetter(l.peekChar()) {
			tok.Type = token.BUTTON
			return tok
		}

		tok.Type = token.UNKOWN
	case 'D':
		// DI pour Drive Impact
		tok.Literal = l.readIdentifier()
		if token.IsIdent(tok.Literal) {
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}

		tok.Type = token.UNKOWN
		return tok
	case 't':
		if l.peekChar() == 'k' {
			tok.Literal = l.readIdentifierUntilChar('k')
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}

		tok = newToken(token.UNKOWN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
		return tok
	}

	l.readChar()
	return tok
}
