package main

import "fmt"

func main() {
	s := "My String"
	sptr1 := &s
	fmt.Println(s)
	fmt.Println(sptr1)
	fmt.Println(*sptr1)

	sptr2 := new(string)
	fmt.Println(sptr2)
	fmt.Println(*sptr2)
	*sptr2 = "Hello"
	fmt.Println(*sptr2)

	var sptr3 *string;
	fmt.Println(sptr3)
	sptr3 = &s
	fmt.Println(sptr3)
	fmt.Println(*sptr3)
}