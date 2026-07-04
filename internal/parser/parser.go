package parser

import "github.com/kaushtubhkanishk/streamlens/internal/lexer"

type parser struct {
	l *lexer.Lexer
}

func newParser(l *lexer.Lexer) *parser {
	return &parser{l: l}
}

func (parser *parser) parseQuery() {
}
