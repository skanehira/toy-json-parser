package json

import (
	"strconv"
	"strings"
)

type Parser struct {
	l         *Lexer
	curToken  Token
	peekToken Token
}

func NewParser(input string) *Parser {
	p := &Parser{
		l: NewLexer(input),
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Parse() interface{} {
	switch p.curToken.Type {
	case NUMBER:
		if strings.Contains(p.curToken.Literal, ".") {
			f, err := strconv.ParseFloat(p.curToken.Literal, 64)
			if err != nil {
				panic("token number is not integer: " + err.Error())
			}
			return f
		}
		num, err := strconv.Atoi(p.curToken.Literal)
		if err != nil {
			panic("token is not float: " + err.Error())
		}
		return num
	case STRING:
		return p.curToken.Literal
	case LBRACE:
		return p.buildObject()
	case LBRACKET:
		return p.buildArray()
	}
	panic("invalid token: " + p.curToken.String())
}

func (p *Parser) buildObject() map[string]interface{} {
	obj := map[string]interface{}{}
	p.nextToken()
	for p.curToken.Type != RBRACE {
		if p.curToken.Type != STRING {
			panic("object key is not string: " + p.curToken.String())
		}

		key := p.curToken.Literal

		if !p.expectPeek(COLON) {
			panic("peek token is not COLON: " + p.peekToken.String())
		}

		p.nextToken() // skip ":"

		obj[key] = p.Parse()

		p.expectPeek(COMMA) // skip if ","
		p.nextToken()
	}
	return obj
}

func (p *Parser) buildArray() []interface{} {
	var values []interface{}
	p.nextToken()
	for p.curToken.Type != RBRACKET {
		values = append(values, p.Parse())
		p.expectPeek(COMMA)
		p.nextToken()
	}
	return values
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) peekTokenIs(t TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	return false
}
