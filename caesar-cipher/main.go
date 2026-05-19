package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(caesarCipher("abc", 3))
}

func caesarCipher(s string, n int) string {
	r := ""
	for _, c := range s {
		if unicode.IsLower(c) {
			r += string(((c - 'a' - rune(n) + 26) % 26) + 'a')
		} else if unicode.IsUpper(c) {
			r += string(((c - 'A' - rune(n) + 26) % 26) + 'A')
		} else {
			r += string(c)
		}
	}
	return r
}

func caesarCipherII(s string, n int) string {
	r := ""
	for _, c := range s {
		switch {
		case unicode.IsLower(c):
			r += string(((c - 'a' - rune(n) + 26) % 26) + 'a')
		case unicode.IsUpper(c):
			r += string(((c - 'A' - rune(n) + 26) % 26) + 'A')
		default:
			r += string(c)
		}
	}
	return r
}

func caesarCipherIII(s string, n int) string {
	r := ""
	for _, c := range s {
		if c == 'z' {
			c = 'a' - 1
		}
		if c == 'Z' {
			c = 'A' - 1
		}
		r += string(c + rune(n))
	}
	return r
}
