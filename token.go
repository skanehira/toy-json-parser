package json

type TokenType string

const (
	ILLEGAL  TokenType = "ILLEGAL"
	NUMBER             = "NUMBER"
	STRING             = "STRING"
	LBRACKET           = "LBRACKET"
	RBRACKET           = "RBRACKET"
	LBRACE             = "LBRACE"
	RBRACE             = "RBRACE"
	COLON              = "COLON"
	COMMA              = "COMMA"
)

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, literal byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(literal),
	}
}
