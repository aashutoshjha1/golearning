package main

import "fmt"

func main() {
    var ptr *int
    num := 4
    
    ptr = &num

    fmt.Println("Value of pointer:", ptr)
    //*ptr = 6
    fmt.Println(num)
    fmt.Println(*ptr)
    *ptr = 6
    fmt.Println(num)
    fmt.Println(*ptr)
    fmt.Println("Value of pointer:", ptr)
}


