package main

import "fmt"

func main() {
    //sum for 1 to n , where n is user input in integer
    var n int
    fmt.Printf("enter value of n : \n")
    fmt.Scanf("%d", &n)
    fmt.Println(n)
    sum :=0 
    for i := 1; i <=n; i++ {
        sum +=i
    }
    fmt.Println(n)
    fmt.Println(sum)
}


    
