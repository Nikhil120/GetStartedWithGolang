package main

import "fmt"

type Phone struct {
	AreaCode string
	Prefix string
	Suffix string
}

type Person struct {
	First string
	Last string
	Age int
	Phone Phone
}

func main() {
	p := Person {
		First: "John",
		Last: "Doe",
		Age: 30,
		Phone: Phone {
			AreaCode: "0123",
			Prefix: "456",
			Suffix: "789",
		},
	}
	fmt.Println(p)
	fmt.Println(p.Age)
	fmt.Println(p.Phone.AreaCode)

	// q := &Person{"Jane", "Doe", 25}
	// fmt.Println(q)
	// fmt.Println(*q)

	pt := struct {
		X int
		Y int
	}{
		X: 10,
		Y: 20,
	}
	fmt.Println(pt)
}