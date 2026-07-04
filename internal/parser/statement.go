package parser

import "github.com/kaushtubhkanishk/streamlens/internal/token"

type selectStatement struct {
	tok         token.Token
	whereClause Clause
	fromClause  Clause
}
