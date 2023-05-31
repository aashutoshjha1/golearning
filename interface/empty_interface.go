package main

import "fmt"


func displayvalue(i interface {}) {
    fmt.Println(i)
}

func main() { 
   a := "test program"
   b := 10
   c := 20
   displayvalue(a)
   displayvalue(b)
   displayvalue(c)

}

 

