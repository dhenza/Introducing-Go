package main

import "fmt"

func main(){

  // Prompt User for Input
  fmt.Print("Enter a number: ")

  // Declare Variable
  var input float64

  // Read user input
  fmt.Scanf("%f", &input)

  // Double Number
  output := input * 2

  // Print Output
  fmt.Println("You entered the number:\t", input, "\nThe number doubled is:\t", output)
}
