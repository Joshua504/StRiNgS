package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(countI("Hello wOrld"))
}

func countI(s string) int{
	// vowels := "aeiouAEIOU"
	count := 0
	for _, c := range s{
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}

	return count
}

func countII(s string) int{
	vowels := "aeiouAEIOU"
	count := 0
	for _, c := range s{
		for _, v := range vowels{
			if c == v {
				count++
			}
		}
	}

	return count
}

func countIII(s string) int{
	vowels := "aeiouAEIOU"
	count := 0
	for _, c := range s{
		if strings.Contains(vowels, string(c)){
			count++
		}
	}

	return count
}