package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simpleTI("Hello {{name}}, you are {{age}} years old!"))
}

func simpleT(s string) string {
	f := strings.Fields(s)
	m := map[string]string{"name": "Bishop", "age": "20"}
	r := []string{}

	for _, marker := range f {
		switch {
		case marker == "{{name}},":
			r = append(r, m["name"])
		case marker == "{{age}}":
			r = append(r, m["age"])
		default:
			r = append(r, marker)
		}
	}
	return strings.Join(r, " ")
}

func simpleTI(s string) string {
	m := map[string]string{"name": "Bishop", "age": "20"}
	for key, val := range m{
		s = strings.ReplaceAll(s, "{{"+key+"}}", val)
	}
	return s
}