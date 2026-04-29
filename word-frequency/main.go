package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(wordFeq("the cat sat on the mat the cat"))
}

func wordFeq(s string) map[string]string{
	sS := strings.Fields(s)
}
