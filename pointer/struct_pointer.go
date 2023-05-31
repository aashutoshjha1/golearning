package main

import "fmt"

func main() {
    type Person struct {
         name string
         age int
    }
    
    person1 := Person{"aashu", 25}
    var ptr *Person
    ptr = &person1
    fmt.Println(person1)
    fmt.Println(ptr)
    fmt.Println(*ptr)
    fmt.Println(ptr.name)
    fmt.Println(ptr.age)
    //fmt.Println(*ptr.age)

}
