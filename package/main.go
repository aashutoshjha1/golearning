package main

import ("fmt"
        "Packages/calculator"
       )

func main() {
    num1 := 20
    num2 :=4
    fmt.Println(calculator.Add(num1, num2))
    fmt.Println(calculator.Subtract(num1, num2))
}
