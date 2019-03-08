package main

import "fmt"

// Creating Structs and using Methods

type Car struct {
  wheels int
  doors int
  manufacturer string
}

func (car *Car) sports_package() {
  car.wheels = 4
  car.doors = 2
  car.manufacturer = "Super Manufacturer"
}

func (car *Car) truck_package() {
  car.wheels = 16
  car.doors = 2
  car.manufacturer = "Truck Manufacturer"
}

func main() {
  mycar := Car{4,4,"Regular Manufacturer"}
  fmt.Println(mycar.wheels, mycar.doors, mycar.manufacturer)
  mycar.sports_package()
  fmt.Println(mycar.wheels, mycar.doors, mycar.manufacturer)
  mycar.truck_package()
  fmt.Println(mycar.wheels, mycar.doors, mycar.manufacturer)
}
