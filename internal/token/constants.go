package token

const (
	ASSIGN          TokenType = "="
	EQUAL           TokenType = "=="
	NEQUAL          TokenType = "!="
	PLUS            TokenType = "+"
	MINUS           TokenType = "-"
	BANG            TokenType = "!"
	NUMBER          TokenType = "NUMBER"
	LEFTBRACKET     TokenType = "{"
	RIGHTBRACKET    TokenType = "}"
	LEFTPARENTHESE  TokenType = "("
	RIGHTPARENTHESE TokenType = ")"
	ILLEGAL         TokenType = "ILLEGAL"
	IDENTIFIER      TokenType = "IDENTIFIER"
	EOF             TokenType = "EOF"
	SELECT          TokenType = "SELECT"
	ASTERISK        TokenType = "*"
	FROM            TokenType = "FROM"
	WHERE           TokenType = "WHERE"
	SLASH           TokenType = "/"
)
