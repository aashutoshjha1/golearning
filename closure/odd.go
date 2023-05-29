package main
import "fmt"

func calculate() func() int {
    num :=1
    return func() int {
       num = num +2
       return num
    }
  }

func main() {
    n :=calculate()
    fmt.Println(n())
    fmt.Println(n())
    n1 :=calculate()
    fmt.Println(n1())
}

