package main

import "fmt"

func update(name *string) {
    *name = "Pooja"
}

func main() {
    var name = "aashu"
    update(&name)
    fmt.Println(name)
}


