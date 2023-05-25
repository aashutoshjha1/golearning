 package main

import "fmt"

func main() {
     // create an array
     numbers := [6] int {11, 22, 33, 44, 55, 70}
     
     // for loop with range
     for i, n := range numbers {
         fmt.Println(numbers[i])
         fmt.Println(numbers[n])

    }

}
 
