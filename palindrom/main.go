package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(palindrome("go"))
}

func palindrome(s string) bool {
	l := strings.ToLower(s)
	r := ""
	for _, c := range l {
		r = string(c) + r
	}
	return r == l
	// return false
}

func palindromeII(s string) bool{
	l := strings.ToLower(s)
	r := ""
	for _, c := range l{
		r = string(c) + r
	}
	if r == l{
		return true
	}
	return false
}