package main

import "fmt"

// A pointer is expressed using a * followed by the data type
// A * is also used to dereference pointer variables.

// Dereferencing a pointer gives us access to the value the pointer points to

func twenty(aptr *int) {
    // Dereferencing the pointer
    fmt.Println("aptr is", *aptr)
    // The following is the same as saying "store the int 20 in the memory location aptr refers to"
  *aptr = 20
}

func main() {
  x := 2
// The & operator is used to find the memory address of the variable, so &x return a pointer to the memory address
  twenty(&x)
  fmt.Println(x) // x is now 20
}
