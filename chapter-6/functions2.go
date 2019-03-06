package main

import "fmt"

func doubleme (x int) (y int) {
  y = x*2
  return y
}

func name_me () (string, string) {
  name := "My Name"
  surname := "My Surname"
  return name, surname
}

func main(){
  x:= 5
  double_value := doubleme(x)
  fmt.Println(double_value)

  name, surname := name_me()
  fmt.Println(name, surname)
}
