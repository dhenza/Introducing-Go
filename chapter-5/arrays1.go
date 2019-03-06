package main

import "fmt"

func main(){
  // Arrays
  fmt.Println("\n")
  fmt.Println("Basic Arrays")
  fmt.Println("============")
  // Create array composed of 5 ints
  var x [5]int
  // Set the 5th element of the array x to 444
  x[4] = 444
  // Set the 2nd element of the array x to 111
  x[1] = 111
  fmt.Println(x)

  fmt.Println("\n")
  fmt.Println("Example of using arrays and a for loop")
  fmt.Println("======================================")

  // Create Array and assign values to all elements
  var y [5]float64
  y[0] = 98
  y[1] = 93
  y[2] = 77
  y[3] = 82
  y[4] = 83

  // Initialize total
  var total float64 = 0

  // Itterate through all values in y, adding them to the total
  for i:= 0; i < len(y); i++ {
    total += y[i]
  }
  // Calculate arthimetic mean and output value
  fmt.Println("Mean value is:", total / float64(len(y)))

  fmt.Println("\n")
  fmt.Println("Changing the loop s structure")
  fmt.Println("=============================")

  z := [5]float64{
    98,
    93,
    77,
    82,
    83,
  }
  
  total = 0

  /* Where _ is the current position in the arary
  and value is the same as y[position].

  _ tells the compiler that we don't need this
  */
  for _, value := range z {
    total += value
  }
  fmt.Println("Mean value is:", total / float64(len(y)))
}
