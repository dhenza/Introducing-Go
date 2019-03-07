package main

import "fmt"

// Panic and Recover

func main() {
  defer func () {
    str := recover() // this will never happen
    fmt.Println(str)
  }()
  panic("PANIC")
}
