package main

import "fmt"

// If you want to create a slice, you should use the built-in make function.
// This creates a slice that is associated with an underlying float64 array of length 5.

// Creating an empty slice
// x := make([]float64, 5)

// Creating an empty slice, length 5, capacity of 10
// x := make([]float64, 5, 10)

func main(){
  fmt.Println("\n")
  fmt.Println("Basic Slicing")
  fmt.Println("=============")
  // Another way of creating a slice is as follows:
  arr := [5]float64{1,2,3,4,5}
  x := arr[0:3]
  fmt.Println(arr, x)

  fmt.Println("\n")
  fmt.Println("Slicing a string")
  fmt.Println("================")

  name := "Hello World"
  y := name[3:5]
  fmt.Println(y)

  fmt.Println("\n")
  fmt.Println("Appending to a slice")
  fmt.Println("====================")
  // Appending to a slice
  slice1 := []int{1,2,3}
  slice2 := append(slice1, 4, 5)
  fmt.Println(slice1)
  fmt.Println(slice2)

  fmt.Println("\n")
  fmt.Println("Copying between slices")
  fmt.Println("====================")

  // Copying an array
  // After running this program slice1 has [1,2,3] and slice2 has [1,2] . The contents
  // of slice1 are copied into slice2 , but because slice2 has room for only two
  // elements, only the first two elements of slice1 are copied.

  slice10 := []int{1,2,3}
  slice11 := make([]int, 2)

  // note that the ordering is copy(dst, src)
  copy(slice11, slice10)
  fmt.Println(slice10, slice11)
}
