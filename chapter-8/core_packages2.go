package main

import (
  "fmt"
  "os"
)

func main() {
  file, err := os.Open("test.txt")
  if err != nil {
    // handle the error here
    fmt.Println("Could not open file")
    return
  }
  defer file.Close()

  // get the file size
  stat, err := file.Stat()
  fmt.Println("\nFile size is", stat, "\n")
  if err != nil {
    fmt.Println("Could retrieve file size")
    return
  }

  // read the file
  content := make([]byte, stat.Size())
  _, err = file.Read(content)
  if err != nil {
    return
  }

  str := string(content)
  fmt.Println(str)
}
