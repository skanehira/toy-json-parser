package json

import (
	"fmt"
	"strings"
)

type Value interface {
	String() string
}

type Number struct {
	Value interface{} // int or float
}

func (n *Number) String() string {
	return fmt.Sprintf("%v", n.Value)
}

type String struct {
	Value string
}

func (s *String) String() string {
	return fmt.Sprintf(`"%s"`, s.Value)
}

type Array struct {
	Values []Value
}

func (a Array) String() string {
	vals := make([]string, len(a.Values))
	for i, v := range a.Values {
		vals[i] = v.String()
	}
	return fmt.Sprintf(`[%s]`, strings.Join(vals, ", "))
}

type Object struct {
	Key   string
	Value Value
}

func (o *Object) String() string {
	return fmt.Sprintf(`{"%s": %s}`, o.Key, o.Value.String())
}
