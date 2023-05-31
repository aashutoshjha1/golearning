package main

import "fmt"


func display() *string {
    message := "hello"
    return &message
}

func main() {

   r :=display()
   fmt.Println("hi ", *r)
 
}
   
