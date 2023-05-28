package main

import "fmt"

func main() {
    students := map[int]string{1: "John", 2: "Lily"}
    fmt.Println(students)
    students[2] = "Robin"
    students[3] = "Aashutosh"
    fmt.Println(students)
    students[8] = "Pooja"
    fmt.Println(students)

}

