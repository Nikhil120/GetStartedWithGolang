package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("Tick")
		time.Sleep(1* time.Second)
		break
	}

	for _, i := range []int{1,2,3,4} {
		fmt.Println(i)
	}

	a := 1
	for a < 4 {
		fmt.Println(a)
		a = a + 1
	}
}