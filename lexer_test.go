package json

import "testing"

func TestReadNumber(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "0",
			want: "0",
		},
		{
			in:   "1",
			want: "1",
		},
		{
			in:   "0.0",
			want: "0.0",
		},
		{
			in:   "10.19",
			want: "10.19",
		},
		{
			in:   "-1",
			want: "-1",
		},
		{
			in:   "-1.1",
			want: "-1.1",
		},
	}

	for _, tt := range tests {
		l := NewLexer(tt.in)
		got := l.readNumber()
		if tt.want != got {
			t.Errorf("unexpected number. want: %s, got: %s", tt.want, got)
		}
	}
}

func TestReadString(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   `"hello\n"`,
			want: `hello\n`,
		},
		{
			in:   `"1\r"`,
			want: `1\r`,
		},
		{
			in:   `"@a:o1\t"`,
			want: `@a:o1\t`,
		},
		{
			in:   `"a\\b"`,
			want: `a\\b`,
		},
		{
			in:   `"a\n\b"`,
			want: `a\n\b`,
		},
	}

	for _, tt := range tests {
		l := NewLexer(tt.in)
		got := l.readString()
		if tt.want != got {
			t.Errorf("unexpected string. want: %s, got: %s", tt.want, got)
		}
	}
}

func TestNextToken(t *testing.T) {
	tests := []struct {
		in   string
		want Token
	}{
		{
			in: "0",
			want: Token{
				Type:    NUMBER,
				Literal: "0",
			},
		},
		{
			in: "99",
			want: Token{
				Type:    NUMBER,
				Literal: "99",
			},
		},
		{
			in: "0.0",
			want: Token{
				Type:    NUMBER,
				Literal: "0.0",
			},
		},
		{
			in: "9.9",
			want: Token{
				Type:    NUMBER,
				Literal: "9.9",
			},
		},
		{
			in: "-1",
			want: Token{
				Type:    NUMBER,
				Literal: "-1",
			},
		},
		{
			in: "-0.0",
			want: Token{
				Type:    NUMBER,
				Literal: "-0.0",
			},
		},
		{
			in: "+1",
			want: Token{
				Type:    ILLEGAL,
				Literal: "+",
			},
		},
		{
			in: `"hello"`,
			want: Token{
				Type:    STRING,
				Literal: "hello",
			},
		},
		{
			in: `"1"`,
			want: Token{
				Type:    STRING,
				Literal: "1",
			},
		},
		{
			in: `"@a:o1"`,
			want: Token{
				Type:    STRING,
				Literal: "@a:o1",
			},
		},
		{
			in: `{`,
			want: Token{
				Type:    LBRACE,
				Literal: "{",
			},
		},
		{
			in: `}`,
			want: Token{
				Type:    RBRACE,
				Literal: "}",
			},
		},
		{
			in: `[`,
			want: Token{
				Type:    LBRACKET,
				Literal: "[",
			},
		},
		{
			in: `]`,
			want: Token{
				Type:    RBRACKET,
				Literal: "]",
			},
		},
		{
			in: `:`,
			want: Token{
				Type:    COLON,
				Literal: ":",
			},
		},
		{
			in: `,`,
			want: Token{
				Type:    COMMA,
				Literal: ",",
			},
		},
	}

	for _, tt := range tests {
		l := NewLexer(tt.in)

		tok := l.NextToken()
		if tt.want.Type != tok.Type {
			t.Fatalf("unexpected token type. want: %s, got: %s", tt.want.Type, tok.Type)
		}

		if tt.want.Literal != tok.Literal {
			t.Errorf("unexpected token literal. want: %s, got: %s", tt.want.Literal, tok.Literal)
		}
	}
}

func TestToken(t *testing.T) {
	in := `
10
"hello"
9.99
[1, "hello"]
{"hello": 1}
{"pk": [1, 2], {"ck": "val"}}
`

	l := NewLexer(in)

	tests := []struct {
		want Token
	}{
		{
			want: Token{
				Type:    NUMBER,
				Literal: "10",
			},
		},
		{
			want: Token{
				Type:    STRING,
				Literal: "hello",
			},
		},
		{
			want: Token{
				Type:    NUMBER,
				Literal: "9.99",
			},
		},
		{
			want: Token{
				Type:    LBRACKET,
				Literal: "[",
			},
		},
		{
			want: Token{
				Type:    NUMBER,
				Literal: "1",
			},
		},
		{
			want: Token{
				Type:    COMMA,
				Literal: ",",
			},
		},
		{
			want: Token{
				Type:    STRING,
				Literal: "hello",
			},
		},
		{
			want: Token{
				Type:    RBRACKET,
				Literal: "]",
			},
		},
		{
			want: Token{
				Type:    LBRACE,
				Literal: "{",
			},
		},
		{
			want: Token{
				Type:    STRING,
				Literal: "hello",
			},
		},
		{
			want: Token{
				Type:    COLON,
				Literal: ":",
			},
		},
		{
			want: Token{
				Type:    NUMBER,
				Literal: "1",
			},
		},
		{
			want: Token{
				Type:    RBRACE,
				Literal: "}",
			},
		},
		{
			want: Token{
				Type:    LBRACE,
				Literal: "{",
			},
		},
		{
			want: Token{
				Type:    STRING,
				Literal: "pk",
			},
		},
		{
			want: Token{
				Type:    COLON,
				Literal: ":",
			},
		},
		{
			want: Token{
				Type:    LBRACKET,
				Literal: "[",
			},
		},
		{
			want: Token{
				Type:    NUMBER,
				Literal: "1",
			},
		},
		{
			want: Token{
				Type:    COMMA,
				Literal: ",",
			},
		},
		{
			want: Token{
				Type:    NUMBER,
				Literal: "2",
			},
		},
		{
			want: Token{
				Type:    RBRACKET,
				Literal: "]",
			},
		},
		{
			want: Token{
				Type:    COMMA,
				Literal: ",",
			},
		},
		{
			want: Token{
				Type:    LBRACE,
				Literal: "{",
			},
		},
		{
			want: Token{
				Type:    STRING,
				Literal: "ck",
			},
		},
		{
			want: Token{
				Type:    COLON,
				Literal: ":",
			},
		},
		{
			want: Token{
				Type:    STRING,
				Literal: "val",
			},
		},
		{
			want: Token{
				Type:    RBRACE,
				Literal: "}",
			},
		},
		{
			want: Token{
				Type:    RBRACE,
				Literal: "}",
			},
		},
	}

	for i, tt := range tests {
		tok := l.NextToken()
		if tt.want.Literal != tok.Literal {
			t.Errorf("test[%d] unexpected token literal. want: %s, got: %s", i, tt.want.Literal, tok.Literal)
		}
		if tt.want.Type != tok.Type {
			t.Fatalf("test[%d] unexpected token type. want: %s, got: %s", i, tt.want.Type, tok.Type)
		}
	}

}
