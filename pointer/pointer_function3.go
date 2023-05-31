package main

import "fmt"

func main() {
    //Declare struct for a person
    type person struct {
         name string 
         age int
       }
    person1 := person{"john", 27}
    fmt.Println(person1)
    fmt.Println(person1.name)
}


