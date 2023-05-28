package main

import "fmt"

func hello() {
    fmt.Println("hello function")
}

func main() {
    fmt.Println("this is main function")
    //calling in function.
    hello()
}


