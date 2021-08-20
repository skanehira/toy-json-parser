package json

type Parser struct {
	input        string
	position     int
	readPosition int // next position
}

func NewParser(input string) *Parser {
	p := &Parser{
		input: input,
	}
	return p
}
