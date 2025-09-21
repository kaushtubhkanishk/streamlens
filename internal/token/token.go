package token

type TokenType string

type Token struct {
	TokenType TokenType
	Value     string
}

func NewToken(TokenType TokenType, value string) *Token {
	return &Token{
		TokenType: TokenType,
		Value:     value,
	}
}
