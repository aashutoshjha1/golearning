package main

import "fmt"

func multiply(n1 int, n2 int) int {
     mult := n1 * n2
     return mult

}

func main() {
    //multiple := multiply(10,12)
    //fmt.Printf("Result : %d \n", multiple)
    fmt.Printf("Result : %d \n", multiply(10,12))
}

