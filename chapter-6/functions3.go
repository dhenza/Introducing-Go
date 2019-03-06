package main

import "fmt"

// Variadic Functions

func var_func(args ...int) (int) {
  total := 0
  for _, v := range args {
    total += v
  }
  return total
}

func main() {
  fmt.Println(var_func (1,2,3))

// Closures are also possible
 add := func(x,y int) int {
   return x + y
   }
   fmt.Println(add(5,9))
}
