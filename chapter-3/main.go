package main

import "fmt"

func main (){

  fmt.Println("\n")
  fmt.Println("Variable Basics")
  fmt.Println("===============")
  // Basic Variable assignment
  var a string
  a = "Hello World"
  fmt.Println(a)

  // Using +=
  var alpha = "alpha"
  alpha += " go"
  fmt.Println(alpha)

  fmt.Println("\n")
  fmt.Println("Evaluating Equality")
  fmt.Println("===============")
  // Evaluating Equality
  var x string = "cat"
  var y string = "dog"
  fmt.Println("Are the strings 'dog' and 'cat' equal?\t", x == y)
    // y will now be assigned the string 'cat'
    y = "cat"
    fmt.Println("Are the strings 'cat' and 'cat' equal?\t", x == y)

  fmt.Println("\n")
  fmt.Println("Allowing the compiler to determine variable type")
  fmt.Println("================================================")
  // Allowing the compiler to determine type
  variable_string := "Hello World"
  variable_int := 2
  variable_float := 3.2
  fmt.Println(variable_string, variable_int, variable_float)
}
