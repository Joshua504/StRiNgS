package main

import "fmt"

func main() {
	fmt.Println(longComPre([]string{"flower", "floor", "flaw"}))
}

func longComPre(s []string) string {
	prefix := ""

	for i := 0; i < len(s[0]); i++ {
		ch := s[0][i] // current character from first string

		for _, word := range s[1:] {
			if i >= len(word) || word[i] != ch {
				return prefix // stop!
			}
		}
		prefix += string(ch)
	}
	return prefix
}
