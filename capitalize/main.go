package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(capitalize("éllo"))
}

func capitalize(s string) string {
	split := strings.Fields(s)
	r := []string{}

	for _, c := range split {
		r = append(r, capIII(c))
	}
	return strings.Join(r, " ")
}

func cap(s string) string {
	fChar := strings.ToUpper(string(s[0]))
	rChar := strings.ToLower(s[1:])

	return fChar + rChar
}

func capII(s string) string {
	fChar := ""
	rChar := ""

	for _, c := range s {
		fChar = strings.ToUpper(string(c))
		break
	}
	rChar = strings.ToLower(s[1:])

	return fChar + rChar
}
func capIII(s string)string{
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