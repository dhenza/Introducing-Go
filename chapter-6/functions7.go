package main

import "fmt"

func twenty (aptr *int) {
  *aptr = 20
}
func main() {
  aptr := new(int) // Creating a new pointer
  twenty(aptr)
  fmt.Println("aptr is", *aptr) // Remembering to dereference the pointer
}
