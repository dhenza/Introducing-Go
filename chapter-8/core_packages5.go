package main

import (
  "fmt"
  "container/list"
)

// Implementing a doubly linked list

func main() {
  var x list.List
  // Adding elements to the List
  x.PushBack(1)
  x.PushBack(2)
  x.PushBack(3)

  // Reading the elements in the list
  for element := x.Front(); element != nil ; element=element.Next() {
    fmt.Println(element.Value.(int))
  }
}
