package main

import "fmt"

func main() {
   type rectangle struct {
       length int
       bredth int
      }
   Rect :=rectangle {20,10}
   fmt.Printf("rectangle length is: %d \n", Rect.length)
   fmt.Printf("rectange bredth is %d \n", Rect.bredth)
   area1 := (Rect.length * Rect.bredth)
   fmt.Printf("Area of rectangle is %d", area1)
}
