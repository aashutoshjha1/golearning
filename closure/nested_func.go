package main
import "fmt"

func greetings(name string) {
     fmt.Printf("welcome %s, wish you all the best \n", name)
     var display_name = func() {
               fmt.Println("hi", name)
     }
     display_name()
     
}

func main() {
    greetings("aashu")
   }

