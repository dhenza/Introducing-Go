package main

import "fmt"

func main(){

  // Integers and Floating-Point Numbers
  fmt.Println("\n")
  fmt.Println("Integers and Floating-Point Numbers")
  fmt.Println("======================================")
  fmt.Println("\nUsing type int, 1 + 1 = ", 1 + 1)
  fmt.Println("Using type float, 1.0 + 1.0 = ", 1.0 + 1.0)
  fmt.Println("Using type float, 1.1 + 1.2 = ", 1.1 + 1.2)
  fmt.Println("\n")

  // Strings
  fmt.Println("Strings")
  fmt.Println("========")
  fmt.Println("The character length of 'Hello World' is", len("Hello World"))
  fmt.Println("The first character of 'Hello World' is", ("Hello World"[0]), "(byte value)")
  fmt.Println("The first character of 'Hello World' is", string("Hello World"[0]))
  fmt.Println("The remaining characters are", string("Hello World"[1:]))
  fmt.Println("The last character of 'Hello World' is", string("Hello World"[10]))
  fmt.Println("Lets " + "concatenate " + "Hello " + "World")
  fmt.Println("\n")

  // Booleans
  fmt.Println("Booleans")
  fmt.Println("========")
  fmt.Println("True AND True:\t\t", true && true)
  fmt.Println("True AND False:\t\t", true && false)
  fmt.Println("True OR True:\t\t", true || true)
  fmt.Println("True OR False:\t\t", true || false)
  fmt.Println("NOT True:\t\t", !true)
  fmt.Println("\n")
}
