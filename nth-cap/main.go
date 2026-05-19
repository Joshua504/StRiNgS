package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(nthCap("hello world am back", 3))
}

func nthCap(s string, n int) string {
	sp := strings.Fields(s)
	r := []string{}

	for i := 0; i < len(sp)-n; i++ {
		r = append(r, sp[i])
	}
	for i := len(sp) - n; i < len(sp); i++ {
		r = append(r, caps(sp[i]))
	}

	return strings.Join(r, " ")
}

func caps(s string) string {
	fChar := ""
	firstRuneSize := 0

	for _, c := range s {
		fChar = strings.ToUpper(string(c))
		firstRuneSize = utf8.RuneLen(c)
		break
	}
	rChar := strings.ToLower(s[firstRuneSize:])
	return fChar + rChar
}
