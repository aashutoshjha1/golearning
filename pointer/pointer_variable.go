package main

import "fmt"

func main() {

   name := "Aashu"
   var ptr *string
   
   ptr = &name
   fmt.Println("variable memory address ", &name)
   fmt.Println("name address is : ", ptr)
   fmt.Println("name ", *ptr)
}

  
