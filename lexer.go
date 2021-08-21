package json

import (
	"unicode"
)

type Lexer struct {
	input        string
	ch           byte
	position     int
	readPosition int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) Skip() {
	for unicode.IsSpace(rune(l.ch)) {
		l.readChar()
	}
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.Skip()

	switch l.ch {
	case '-':
		if unicode.IsNumber(rune(l.input[l.readPosition])) {
			lit := l.readNumber()
			tok = Token{Type: NUMBER, Literal: lit}
		} else {
			tok = Token{
				Type:    ILLEGAL,
				Literal: string(l.ch),
			}
		}
	case '{':
		tok = NewToken(LBRACE, l.ch)
	case '}':
		tok = NewToken(RBRACE, l.ch)
	case '[':
		tok = NewToken(LBRACKET, l.ch)
	case ']':
		tok = NewToken(RBRACKET, l.ch)
	case ':':
		tok = NewToken(COLON, l.ch)
	case ',':
		tok = NewToken(COMMA, l.ch)
	case '"':
		lit := l.readString()
		tok = Token{Type: STRING, Literal: lit}
	case 'n':
		lit := l.readNull()
		if lit == "null" {
			tok = Token{
				Type:    NULL,
				Literal: "null",
			}
		} else {
			tok = Token{
				Type:    ILLEGAL,
				Literal: lit,
			}
		}
		return tok
	case 't', 'f':
		lit := l.readBool()
		if lit == "true" || lit == "false" {
			tok = Token{
				Type:    BOOL,
				Literal: lit,
			}
		} else {
			tok = Token{
				Type:    ILLEGAL,
				Literal: lit,
			}
		}

		return tok
	case 0:
		tok = NewToken(EOF, l.ch)
	default:
		switch {
		case unicode.IsNumber(rune(l.ch)):
			lit := l.readNumber()
			tok = Token{Type: NUMBER, Literal: lit}
			return tok
		default:
			tok = Token{
				Type:    ILLEGAL,
				Literal: string(l.ch),
			}
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readBool() string {
	pos := l.position
	end := pos
	if l.ch == 't' {
		end += 4
	} else {
		end += 5
	}

	for l.position < end {
		l.readChar()
		if l.ch == 0 {
			break
		}
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNull() string {
	pos := l.position
	end := pos + 4
	for l.position < end {
		l.readChar()
		if l.ch == 0 {
			break
		}
	}

	return l.input[pos:l.position]
}

func (l *Lexer) readString() string {
	pos := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position
	var dot int
	for unicode.IsNumber(rune(l.ch)) || isDot(l.ch) || l.ch == '-' {
		if dot > 1 {
			break
		}

		if isDot(l.ch) {
			dot++
		}
		l.readChar()
	}

	return l.input[pos:l.position]
}

func isDot(char byte) bool {
	return char == '.'
}
