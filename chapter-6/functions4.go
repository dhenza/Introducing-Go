package main

import "fmt"

// Using the defer statement
func one() {
  fmt.Println("One")
}

func two() {
  fmt.Println("Two")
}

func three() {
  fmt.Println("Three")
}

func main(){
  defer two()
  one()
}

// Example
// f, _ := os.Open(filename)
// defer f.Close()
