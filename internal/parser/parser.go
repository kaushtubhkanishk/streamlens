package parser

import "github.com/kaushtubhkanishk/streamlens/internal/token"

type selectStatement struct {
	literal token.Token
	table   string
	cols    []string
}

//func NewSelectStatement(tok token.Token, table string) *selectStatement {
//	return &selectStatement{
//		tok:   literal,
//		table: table,
//	}
//}

func (s *selectStatement) TokenLiteral() string {
	return s.literal.Value
}

func (s *selectStatement) statementNode() {

}
