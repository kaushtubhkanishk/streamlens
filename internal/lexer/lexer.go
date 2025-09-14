package lexer

import (
	"strings"
)

type TokenType string

type Token struct {
	TokenType TokenType
	Value     string
}

type Lexer struct {
	currentPosition int
	nextPosition    int
	currentChar     byte
	input           string
}

func NewLexer(input string) *Lexer {
	if len(input) == 0 {
		return &Lexer{currentPosition: 0, nextPosition: 0, currentChar: 0}
	}
	return &Lexer{
		input:           input,
		currentPosition: 0,
		nextPosition:    1,
		currentChar:     input[0],
	}
}

// ReadChar reads the next character and increments the current and the next pointer to the input buffer
func (l *Lexer) ReadChar() {
	if l.nextPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.nextPosition]
	}
	l.currentPosition = l.nextPosition
	l.nextPosition++
}

func (l *Lexer) PeekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}

// ReadIdentifier reads an identifier if possible from the identified lexeme
func (l *Lexer) ReadIdentifier() string {
	position := l.currentPosition
	for l.isLetter(l.currentChar) {
		l.ReadChar()
	}
	return l.input[position:l.currentPosition]
}

// ReadNumber reads a number until the end
func (l *Lexer) ReadNumber() string {
	position := l.currentPosition
	for l.isDigit(l.currentChar) {
		l.ReadChar()
	}
	return l.input[position:l.currentPosition]
}

func (l *Lexer) isLetter(ch byte) bool {
	return ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' || ch == '_'
}

func (l *Lexer) isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\r' || l.currentChar == '\n' {
		l.ReadChar()
	}
}

func (l *Lexer) NextToken() Token {
	var token Token
	l.skipWhitespace()
	switch l.currentChar {
	case '+':
		token = Token{PLUS, string(l.currentChar)}
	case '-':
		token = Token{MINUS, string(l.currentChar)}
	case '=':
		if char := l.PeekChar(); char == '=' {
			l.ReadChar()
			token = Token{EQUAL, "=="}
		} else {
			token.TokenType = ASSIGN
			token.Value = string(l.currentChar)
		}
	case '{':
		token = Token{LEFTBRACKET, string(l.currentChar)}
	case '}':
		token = Token{RIGHTBRACKET, string(l.currentChar)}
	case '!':
		if char := l.PeekChar(); char == '=' {
			l.ReadChar()
			token = Token{NEQUAL, "!="}
		} else {
			token = Token{BANG, string(l.currentChar)}
		}
	case '(':
		token = Token{LEFTPARENTHESE, string(l.currentChar)}
	case ')':
		token = Token{RIGHTPARENTHESE, string(l.currentChar)}
	case '*':
		token = Token{ASTERISK, string(l.currentChar)}
	case '/':
		token = Token{SLASH, string(l.currentChar)}
	case 0:
		token = Token{EOF, ""}
	default:
		if l.isLetter(l.currentChar) {
			token.Value = l.ReadIdentifier()
			if tokenType := isKeyword(token.Value); tokenType != "" {
				token.TokenType = tokenType
			} else {
				token.TokenType = IDENTIFIER
			}
		} else if l.isDigit(l.currentChar) {
			token.TokenType = NUMBER
			token.Value = l.ReadNumber()
		} else {
			token.TokenType = ILLEGAL
			token.Value = string(l.currentChar)
		}
	}
	l.ReadChar()
	return token
}

func isKeyword(value string) TokenType {
	value = strings.ToUpper(value)
	switch value {
	case "SELECT":
		return SELECT
	case "FROM":
		return FROM
	case "WHERE":
		return WHERE
	}
	return ""
}
