package main

import "fmt"

// Go routines allow multiple functions to be execured concurrently

func myfunction(i int) {
  fmt.Println(i)
}

func main() {
  for i := 0; i < 10; i++ {
      go myfunction(i)
  }
  var input string
  fmt.Scanln(&input)
}
