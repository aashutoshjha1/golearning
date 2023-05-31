package main

import "fmt"

func main() {

    var a interface {}
        a = 10
        a = "hello world"
        interfaceValue, booleanValue := a.(string)
        fmt.Println(interfaceValue)
        fmt.Println("Boolean Value:", booleanValue)
        iValue, bValue := a.(int)
        fmt.Println("Interface Value:", iValue)
        fmt.Println("Boolean Value:", bValue)
}


