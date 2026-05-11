package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(truncate("Hello world from Go programming", 3))
}

func truncate(s string, n int) string{
	split := strings.Fields(s)

	if len(split) <= n{
		return strings.Join(split, " ")
	}

	return strings.Join(split[:n], " ") + "..."
}
