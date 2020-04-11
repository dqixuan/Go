package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "asdfa"
	s1 := strings.ToLower(s)
	strings.Compare(s, s1)
	fmt.Println(strings.Title(s))
}
