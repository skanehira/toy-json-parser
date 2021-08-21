package json

import "testing"

func TestValueString(t *testing.T) {
	tests := []struct {
		array Value
		want  string
	}{
		{
			array: Array{
				Values: []Value{
					&Number{
						Value: 1,
					},
					&String{
						Value: "hello",
					},
					&Object{
						Key: "unko",
						Value: &Array{
							Values: []Value{
								&String{
									Value: "str",
								},
								&Number{
									Value: 05.4,
								},
							},
						},
					},
					&Number{
						Value: 10.5,
					},
					&Object{
						Key: "gorilla",
						Value: &String{
							Value: "unko",
						},
					},
				},
			},
			want: `[1, "hello", {"unko": ["str", 5.4]}, 10.5, {"gorilla": "unko"}]`,
		},
	}

	for _, tt := range tests {
		got := tt.array.String()
		if got != tt.want {
			t.Errorf("unexpected strings. want: %s, got: %s", tt.want, got)
		}
	}
}
