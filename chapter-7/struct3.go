package main

import "fmt"

type Person struct {
  name string
}

type Employee struct {
  Person
  Model string
}

func main() {
  i := new(Employee)
  i.Person.name = "Firstname Lastname"
  i.Model = "Super Manufacturer"
  fmt.Println(i.Person.name, i.Model)
}
