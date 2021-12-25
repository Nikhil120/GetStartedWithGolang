package main

import "fmt"

func main() {
    var x interface{}
    x = "Hello, World"
    s, ok := x.(string)
    if !ok {
        panic("No!")
    }
    fmt.Println(s)

    switch x.(type) {
    case int:
        fmt.Println("Integer!")
    case string:
        fmt.Println("String!")
    }
}