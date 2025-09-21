package lexer

import (
	"strings"

	"github.com/kaushtubhkanishk/streamlens/internal/token"
)

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

func (l *Lexer) NextToken() token.Token {
	tok := token.NewToken("", "")
	l.skipWhitespace()
	switch l.currentChar {
	case '+':
		tok.TokenType = token.PLUS
		tok.Value = string(l.currentChar)
	case '-':
		tok.TokenType = token.MINUS
		tok.Value = string(l.currentChar)
	case '=':
		if char := l.PeekChar(); char == '=' {
			l.ReadChar()
			tok.TokenType = token.EQUAL
			tok.Value = string(l.currentChar)
		} else {
			tok.TokenType = token.ASSIGN
			tok.Value = string(l.currentChar)
		}
	case '{':
		tok.TokenType = token.LEFTBRACKET
		tok.Value = string(l.currentChar)
	case '}':
		tok.TokenType = token.RIGHTBRACKET
		tok.Value = string(l.currentChar)
	case '!':
		if char := l.PeekChar(); char == '=' {
			l.ReadChar()
			tok.TokenType = token.BANG
			tok.Value = string(l.currentChar)
		} else {
			tok.TokenType = token.NEQUAL
			tok.Value = string(l.currentChar)
		}
	case '(':
		tok.TokenType = token.LEFTPARENTHESE
		tok.Value = string(l.currentChar)
	case ')':
		tok.TokenType = token.RIGHTPARENTHESE
		tok.Value = string(l.currentChar)
	case '*':
		tok.TokenType = token.ASTERISK
		tok.Value = string(l.currentChar)
	case '/':
		tok.TokenType = token.SLASH
		tok.Value = string(l.currentChar)
	case 0:
		tok.TokenType = token.EOF
		tok.Value = string(l.currentChar)
	default:
		if l.isLetter(l.currentChar) {
			tok.Value = l.ReadIdentifier()
			if tokenType := isKeyword(tok.Value); tokenType != "" {
				tok.TokenType = tokenType
			} else {
				tok.TokenType = token.IDENTIFIER
			}
		} else if l.isDigit(l.currentChar) {
			tok.TokenType = token.NUMBER
			tok.Value = l.ReadNumber()
		} else {
			tok.TokenType = token.ILLEGAL
			tok.Value = string(l.currentChar)
		}
	}
	l.ReadChar()
	return *tok
}

func isKeyword(value string) token.TokenType {
	value = strings.ToUpper(value)
	switch value {
	case "SELECT":
		return token.SELECT
	case "FROM":
		return token.FROM
	case "WHERE":
		return token.WHERE
	}
	return ""
}
