package main

import "fmt"

func displayn() func() int {
    num :=1
    return func() int {
        num++
        return num
    }
}
func main() {
   closure1 :=displayn()
   fmt.Println(closure1())
   fmt.Println(closure1())
   fmt.Println(closure1())
   fmt.Println(closure1())
   fmt.Println(closure1())
   
   closure2 := displayn()
   fmt.Println(closure2())
   fmt.Println(closure2())

}
