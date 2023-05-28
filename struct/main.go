package main

import "fmt"

func main() {
     type student struct {
          name string
          age int
          marks float64
       }
      person := student{"aashu", 38, 500.60}
      fmt.Println(person)
}

