package main

import "fmt"

func main(){

  // For loops
  fmt.Println("\n")
  fmt.Println("For Loops")
  fmt.Println("=========")

  i := 1
  for i <= 10 {
    fmt.Println(i)
    i += 1
  }

  // If Statements
  fmt.Println("\n")
  fmt.Println("If Statements")
  fmt.Println("=============")

  j := 1
  for j <= 10 {
    if j % 2 == 0 {
    fmt.Println(j, "\t\tEven")
    } else {
    fmt.Println(j, "\t\tOdd")
    }
    j += 1
  }

// Using the Switch Statement
  fmt.Println("\n")
  fmt.Println("Switch Statements")
  fmt.Println("=================")

  k := 1
  for k <= 10 {
    switch k {
    case 1: fmt.Println(k, "\t\tOne")
    case 2: fmt.Println(k, "\t\tTwo")
    case 3: fmt.Println(k, "\t\tThree")
    case 4: fmt.Println(k, "\t\tFour")
    default: fmt.Println(k, "\t\tAbove 5")
    }
    k += 1
  }
}
