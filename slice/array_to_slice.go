package main

import "fmt"

func main() {
    array1 := [8]int{2,3,4,5,6,7,8,9}
    slice1 := array1[0:5]
    fmt.Println(slice1)
    
}

