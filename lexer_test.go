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
			in:   `"\"\r\n\t"`,
			want: `\"\r\n\t`,
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
				Type:    Number,
				Literal: "0",
			},
		},
		{
			in: "99",
			want: Token{
				Type:    Number,
				Literal: "99",
			},
		},
		{
			in: "0.0",
			want: Token{
				Type:    Number,
				Literal: "0.0",
			},
		},
		{
			in: "9.9",
			want: Token{
				Type:    Number,
				Literal: "9.9",
			},
		},
		{
			in: "-1",
			want: Token{
				Type:    Number,
				Literal: "-1",
			},
		},
		{
			in: "-0.0",
			want: Token{
				Type:    Number,
				Literal: "-0.0",
			},
		},
		{
			in: "+1",
			want: Token{
				Type:    Illegal,
				Literal: "+",
			},
		},
		{
			in: `"hello"`,
			want: Token{
				Type:    String,
				Literal: "hello",
			},
		},
		{
			in: `"1"`,
			want: Token{
				Type:    String,
				Literal: "1",
			},
		},
		{
			in: `"@a:o1"`,
			want: Token{
				Type:    String,
				Literal: "@a:o1",
			},
		},
		{
			in: `"\""`,
			want: Token{
				Type:    String,
				Literal: `\"`,
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
