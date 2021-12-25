package main

import (
	"fmt"
	"strings"
)

func main() {
	// boolean
	a := true
	b := true
	fmt.Printf("a = %v, b = %v, and a == b = %v\n", a, b, a == b)


	// number
	i := 100
	var j int = 1234
	fmt.Printf("%v + %v = %v\n", i, j, i+j)

	f := 1.5
	fmt.Printf("f = %v\n", f)


	// string
	s := "I'm a string"
	fmt.Printf("Ends with string? %v\n", strings.HasSuffix(s, "string"))
}