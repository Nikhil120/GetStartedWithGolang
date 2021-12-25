package main

import (
	"fmt"
)

func main() {
	a := 1

	// if else
	if a < 0 {
		fmt.Println("Negative Value")
	} else if a == 0 {
		fmt.Println("Zero Value")
	} else {
		fmt.Println("Positive Value")
	}

	// switch
	switch a {
	case 1:
		fmt.Println("One!")
	case 2:
		fmt.Println("Two!")
	default:
		fmt.Println("None of these")
	}
}