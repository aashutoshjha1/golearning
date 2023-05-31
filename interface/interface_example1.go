package main

import "fmt"
 
type Shape interface {
     area() float32
     square() float32
}

type Rectangle struct {
    length, breadth float32

}

func (r Rectangle) area() float32 {
    return r.length * r.breadth
}

func (r Rectangle) square() float32 {
     return r.length * r.breadth
}

func calculate(s Shape) {
    fmt.Println("Area : " , s.area())
    fmt.Println("Square: ", s.square())

}
func main() {
    rect := Rectangle{7, 7}
    calculate(rect)

}




