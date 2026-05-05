package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(anagram("Astronomer", "Moon starer"))
}

func anagram(s, v string) bool{
	fW := strings.ToLower(s)
	sW := strings.ToLower(v)
	m := make(map[string]int)
	m2 := make(map[string]int)

	for _, c := range fW {
		if c == ' ' {
			continue
		}
		m[string(c)]++
	}
	for _, b := range sW {
		if b == ' ' {
			continue
		}
		m2[string(b)]++
	}
	if len(m) != len(m2) {
		return false
	}
	for key, val := range m {
		if m2[key] != val {
			return false
		}
	}
	return true
}