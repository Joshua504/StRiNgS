package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(wordFeq("the cat sat on the mat the cat"))
}

func wordFeq(s string) map[string]int{
	sS := strings.Fields(s)
	m := make(map[string]int)
	for _, s := range sS {
		m[s]++
	}
	return m
}