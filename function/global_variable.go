package main
import "fmt"

var sum int

func sum_func() {
    sum = 10+15

}

func main() {
    sum_func()
    fmt.Printf("Result : %d", sum)

}

