package main

import "fmt"

func main(){
  // Using a constant
  const x string = "abc"
  fmt.Println(x)

  // Defining multiple variables
  var (
    a = "alpha"
    b = 10
    c = "beta"
  )
  fmt.Println(a,b,c)
}
