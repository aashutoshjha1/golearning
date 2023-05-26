package main

import "fmt"

func main() {
     numbers := [5]int{21, 24, 27, 30, 33}
     for index, item := range numbers {
           fmt.Printf("numbers[%d] = %d \n", index, item)
}
}

