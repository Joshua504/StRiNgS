package main

import "fmt"

func main() {
	fmt.Println(permutation("abcde"))
}

func permutation(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}

	result := []string{}

	for i := 0; i < len(s); i++ {
		firstChar := string(s[i])
		remaining := s[:i] + s[i+1:]

		for _, perm := range permutation(remaining) {
			result = append(result, firstChar+perm)
		}
	}

	return result
}
