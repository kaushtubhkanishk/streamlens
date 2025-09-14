package lexer

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
	EOF             TokenType = ""
	SELECT          TokenType = "SELECT"
	STAR            TokenType = "*"
	FROM            TokenType = "FROM"
)
