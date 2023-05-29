package main

import "fmt"

func nested() func() {
     return func() {
          fmt.Println("return function")
     }
   }

func main() {
    g1 :=nested()
    g1()
}
