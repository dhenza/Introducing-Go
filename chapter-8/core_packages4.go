package main

import (
  "errors"
  "fmt"
  "io/ioutil"
  "os"
)

func main() {
  file, err := os.Create("test.txt")
  if err != nil {
    // handle the error
    
    return
  }
  defer file.Close()

  file.WriteString("Test Text Right Here")

  content, err := ioutil.ReadFile("test.txt")
  if err != nil {
    fmt.Println("Could not open file")
    return
  }
  str := string(content) // converts to string
  fmt.Println(str) // prints string
}
