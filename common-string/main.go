package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(commonString("abcdef", "zcdemf"))
	fmt.Println(commonStringI("abcdef", "zcdemf"))
}

func commonString(s, x string) string {
	var r strings.Builder
	for _, c := range s {
		for _, v := range x {
			if c == v {
				r.WriteString(string(c))
			}
		}
	}
	return r.String()
}

func commonStringI(s, x string) string {
	var r string
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			sub := s[i:j]

			if strings.Contains(x, sub) && len(sub) > len(r) {
				r = sub
			}
		}
	}
	return r
}
