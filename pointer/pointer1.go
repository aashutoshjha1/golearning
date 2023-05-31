package main
//Creating pointers with new() function 
import  "fmt"

func main() {
   var ptr = new(int)
   *ptr = 6
   fmt.Println(*ptr)

}

