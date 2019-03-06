package main

import "fmt"

func main (){
  // Example of initializing and using a map
  x := make(map[string]int)
  x["key"] = 10
  fmt.Println(x["key"])

  // This could be an int for the key and a string for the value
  // String/String, etc would also be allowed
  y := make(map[int]string)
  y[11] = "Eleven"
  y[12] = "Twelve"
  fmt.Println("Printing the value of a single element", y[11])
  fmt.Println("Printing the entire map",y)

  // Deleteting an element from a map
  // Note that 11 is the key value
  delete(y,11)
  fmt.Println("Printing the entire map",y)

  // One way of creating elements in a map
  alphabet := make(map[string]string)
  alphabet["A"] = "a"
  alphabet["B"] = "b"
  alphabet["C"] = "c"
  fmt.Println(alphabet)

  // Another way of creating elements in a map
  new_alphabet := map[string]string{
    "A": "a",
    "B": "b",
    "C": "c",
  }
    fmt.Println(new_alphabet)
}
