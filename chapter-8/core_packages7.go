package main

import (
  "fmt"
  "doubler"
)

func main() {
  x := 5
  y := doubler.Double(x)
  fmt.Println(x, y)
}

/*

The folllowing was created in $GOPATH/src/doubler/doubler.go

package doubler

func Double(x int) int {
  return x*2
}

*/
