package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	s = append(s, 5)
	fmt.Println(s)

	s1 := make([]int, 10)
	s1[4] = 1
	fmt.Println(s1)

	s2 := [4]int{1, 2, 3, 4}
	fmt.Println(s2)
}