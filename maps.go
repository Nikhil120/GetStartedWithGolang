package main

import "fmt"

func main() {
	m := map[string]string{}
	m["first"] = "John"
	m["last"] = "Doe"
	fmt.Println(m)

	
	m1 := map[string]string{
		"first": "John",
		"last": "Doe",
	}
	fmt.Println(m1)

	m2 := make(map[string]string)
	m2["first"] = "John"
	m2["last"] = "Doe"
	fmt.Println(m2)
	fmt.Println(m2["last"])
}