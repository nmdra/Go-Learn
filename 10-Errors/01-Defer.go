package main
import "fmt"

func main() {

  // defer the execution of Println() function
  defer fmt.Println("Three")
  defer fmt.Println("Four")

  fmt.Println("One")
  defer fmt.Println("Five")
  fmt.Println("Two")
}
