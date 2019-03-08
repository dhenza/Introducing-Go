package main

import (
  "fmt"
  "io/ioutil"
)

// An alternative way to read the contents of a file

func main(){
  content, err := ioutil.ReadFile("test.txt")
  if err != nil {
    fmt.Println("Could not open file")
    return
  }
  fmt.Println(content) // prints byte values
  str := string(content) // converts to string
  fmt.Println(str) // prints string
}
