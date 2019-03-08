package main

import "fmt"

type Circle struct {
  x float64
  y float64
  z float64
}

type Rectangle struct {
  height, width float64
}

// Creating a new instance of the Circle type

// Simply creating a new instance
// var c Circle

// Specifying initial values
// c:= Circle{x:1, y:2, z:3}

// Field names may be omitted, but order must be maintained
// c:= Circle{1,2,3}

// If we want the pointer we can do the following
// c:= &Circle{1,2,3}

// Allocating memeory for all fields, setting them to 0 and returning a pointer to the struct
// (returns *Circle)

// c:= new(Circle)


// Initializing the new instance of Rectangle

func double_size(rectangle *Rectangle) {
  rectangle.height = rectangle.height * 2
  rectangle.width = rectangle.width * 2
}
func main() {
  rectangle_1 := Rectangle{height:10, width:20}

  // Accessing the fields
  fmt.Println(rectangle_1.height, rectangle_1.width)
  // Modifying field values
  double_size(&rectangle_1)
  fmt.Println(rectangle_1.height, rectangle_1.width)
}
