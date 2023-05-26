package main

import "fmt"

func main() {
    age := [...]int{10,18,38,20,14}
    for i:=0; i < len(age); i++ {
         fmt.Printf("age of %d is : %d \n", i,age[i])

  }
}

