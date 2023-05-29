package main

import "fmt"

func recurse(n int) {
     if  n > 0 {
         fmt.Printf( "countdown is %d \n", n)
         recurse(n - 1)
     }
}
func main() {
    recurse(5)

}

