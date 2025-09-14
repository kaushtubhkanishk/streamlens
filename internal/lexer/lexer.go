package lexer

import "fmt"

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
	return &Lexer{
		input:           input,
		currentPosition: 0,
		nextPosition:    1,
		currentChar:     input[0],
	}
}

func (l *Lexer) ReadChar() {
	if l.nextPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.currentPosition]
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

// read an identifier if possible from the identified lexeme
func (l *Lexer) ReadIdentifier() string {
	position := l.currentPosition
	for l.isLetter(l.currentChar) {
		fmt.Printf("%v", string(l.currentChar))
		l.ReadChar()
	}
	return l.input[position:l.currentPosition]
}

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

func (l *Lexer) NextToken() Token {
	var token Token
	switch l.currentChar {
	case '+':
		token = Token{PLUS, string(l.currentChar)}
	case '-':
		token = Token{MINUS, string(l.currentChar)}
	case '=':
		if char := l.PeekChar(); char == '=' {
			token = Token{EQUAL, "=="}
		} else {
			token = Token{ASSIGN, string(l.currentChar)}
		}
	case '{':
		token = Token{LEFTBRACKET, string(l.currentChar)}
	case '}':
		token = Token{RIGHTBRACKET, string(l.currentChar)}
	case '!':
		if char := l.PeekChar(); char == '=' {
			token = Token{NEQUAL, "!="}
		} else {
			token = Token{BANG, string(l.currentChar)}
		}
	case '(':
		token = Token{LEFTPARENTHESE, string(l.currentChar)}
	case ')':
		token = Token{RIGHTPARENTHESE, string(l.currentChar)}
	case ' ':
	case 0:
		token = Token{EOF, ""}
	default:
		if l.isLetter(l.currentChar) {
			token.TokenType = IDENTIFIER
			token.Value = l.ReadIdentifier()
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
