package json

type TokenType string

const (
	Illegal TokenType = "Illegal"
	Number            = "Number"
)

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}
