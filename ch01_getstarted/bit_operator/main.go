package main

import "fmt"

func main() {

    n :=212
    for i :=0; i <=4; i++ {
        fmt.Printf("Right Shift by %d: %d\n", i, n>>i)
    }
  //Small test without loop
  a :=8
  fmt.Println(a>>1)
}

 
