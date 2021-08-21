package json

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseString(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   `"gorilla"`,
			want: "gorilla",
		},
		{
			in:   `"hello world"`,
			want: "hello world",
		},
	}

	for _, tt := range tests {
		p := NewParser(tt.in)
		got := p.Parse()
		if tt.want != got {
			t.Errorf("unexpected result. want: %v, got: %v", tt.want, got)
		}
	}
}

func TestParseInteger(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{
			in:   `99`,
			want: 99,
		},
		{
			in:   `191`,
			want: 191,
		},
	}

	for _, tt := range tests {
		p := NewParser(tt.in)
		got := p.Parse()
		if tt.want != got {
			t.Errorf("unexpected result. want: %v, got: %v", tt.want, got)
		}
	}
}

func TestParseFloat(t *testing.T) {
	tests := []struct {
		in   string
		want float64
	}{
		{
			in:   `10.1`,
			want: 10.1,
		},
		{
			in:   `191.5`,
			want: 191.5,
		},
	}

	for _, tt := range tests {
		p := NewParser(tt.in)
		got := p.Parse()
		if tt.want != got {
			t.Errorf("unexpected result. want: %v, got: %v", tt.want, got)
		}
	}
}

func TestParseObject(t *testing.T) {
	tests := []struct {
		in   string
		want map[string]interface{}
	}{
		{
			in: `{"hello": "world"}`,
			want: map[string]interface{}{
				"hello": "world",
			},
		},
		{
			in: `{"hello": "world", "number": 10, "float": 1.1}`,
			want: map[string]interface{}{
				"hello":  "world",
				"number": 10,
				"float":  1.1,
			},
		},
		{
			in: `{"hello": "world", "obj": {"inner": "obj"}, "array": [1, 2, 10.5]}`,
			want: map[string]interface{}{
				"hello": "world",
				"obj":   map[string]interface{}{"inner": "obj"},
				"array": []interface{}{
					1,
					2,
					10.5,
				},
			},
		},
	}

	for _, tt := range tests {
		p := NewParser(tt.in)
		got := p.Parse()
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("unexpected result. want: %#v, got: %#v", tt.want, got)
		}
	}
}

func TestParseArray(t *testing.T) {
	tests := []struct {
		in   string
		want []interface{}
	}{
		{
			in:   `[1, 2, 3]`,
			want: []interface{}{1, 2, 3},
		},
		{
			in:   `["a", "b", "c"]`,
			want: []interface{}{"a", "b", "c"},
		},
		{
			in: `[{"hello": "world"}, 1, 10.5, "gorilla", [1, 2, 3]]`,
			want: []interface{}{
				map[string]interface{}{
					"hello": "world",
				},
				1,
				10.5,
				"gorilla",
				[]interface{}{
					1, 2, 3,
				},
			},
		},
	}

	for _, tt := range tests {
		p := NewParser(tt.in)
		got := p.Parse()
		diff := cmp.Diff(tt.want, got)
		if diff != "" {
			t.Errorf("unexpected array (-want +got)\n%s", diff)
		}
	}
}
