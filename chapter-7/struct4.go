package main

import "fmt"
import "math"

// Using Interfaces

// Like a struct, an interface is created using the type keyword, followed by a name and
// the keyword interface . But instead of defining fields, we define a method set. A
// method set is a list of methods that a type must have in order to implement the inter‐face

type Shape interface {
  calculate_area() float64
}

type Square struct {
  side_length int
  area float64
}

type Circle struct {
  radius int
  area float64
}

func totalArea(shapes ...Shape) float64 {
  var total_area float64
  for _, s := range shapes {
    total_area += s.calculate_area()
  }
  return total_area
}

func (object *Square) calculate_area() float64 {
  return float64(object.side_length*2)
}

func (object *Circle) calculate_area() float64 {
  return float64(math.Pi*float64(object.radius^2))
}

func main() {
    // Create objects
    my_square := Square{2,0}
    my_circle := Circle{5,0}
    // Calculate total area
    my_total := totalArea(&my_square, &my_circle)
    // Print total area
    fmt.Println(my_total)

}
