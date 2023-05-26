package main

import "fmt"

func main() {
    var arrayOfString = [...]string{"Hello", "Programiz"}
    var integerarray = [...]int{1,2,3,4,10,20,21}
    var array3[4] int
    fmt.Println(arrayOfString)
    fmt.Println(integerarray)
    array3[0] = 5
    array3[1] = 10
    array3[2] = 20
    array3[3] = 30
    fmt.Println(array3)
    fmt.Println(integerarray[2])
}
