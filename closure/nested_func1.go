package main
import "fmt"

func greetings(name string) {
     fmt.Printf("welcome %s, wish you all the best \n", name)
     func internal_func() {
         fmt.Println("Hi", name) 
     }
     internal_func()
}

func main() {
    greetings("aashu")
   }
//Nested function like this doesn't work in GO, we have to assign function to variable to make it work

