package main

import "fmt"

func main() {
    slice1 := []int{1,3,4}
    slice2 := []int{2,5,6}
    slice1 = append(slice1, slice2...)
    slice3 := append(slice2, 4)
    fmt.Println(slice1) 
    fmt.Println(slice3)
}

