package main

import "fmt"

func main() {
    b :=20
    var num *int = &b
    n := &b
    fmt.Println(&num)
    fmt.Println(&n)
}


