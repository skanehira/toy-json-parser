package main

import (
	"fmt"
	"json"
)

func main() {
	input := `{"hello": "world", "array": [{"name": "gorilla"}]}`
	obj := json.NewParser(input).Parse().(map[string]interface{})
	fmt.Println(obj["hello"])                    // world
	fmt.Println(obj["array"].([]interface{})[0]) // map[name:gorilla]
}
