package main

import "fmt"


func main(){
	fmt.Println(reverse("héllo"))
}

func reverse (s string) string{
	r := ""
	for _, c := range s {
		r = string(c) + r
	}
	return r
}