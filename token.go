package json

import "fmt"

type TokenType string

const (
	ILLEGAL  TokenType = "ILLEGAL"
	NUMBER             = "NUMBER"
	STRING             = "STRING"
	NULL               = "NULL"
	BOOL               = "BOOL"
	LBRACKET           = "LBRACKET"
	RBRACKET           = "RBRACKET"
	LBRACE             = "LBRACE"
	RBRACE             = "RBRACE"
	COLON              = "COLON"
	COMMA              = "COMMA"
	EOF                = "EOF"
)

type Token struct {
	Type    TokenType
	Literal string
}

func (t *Token) String() string {
	return fmt.Sprintf("token type: %s, literal: %s", t.Type, t.Literal)
}

func NewToken(tokenType TokenType, literal byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(literal),
	}
}
