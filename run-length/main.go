package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(runLength("aabaa"))
}

func runLength(s string) string {
	r := ""
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			r += strconv.Itoa(count) + string(s[i-1])
			count = 1
		}
	}

	r += strconv.Itoa(count) + string(s[len(s)-1])

	return r
}
