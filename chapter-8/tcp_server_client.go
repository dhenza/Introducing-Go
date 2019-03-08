package main

import (
  "encoding/gob"
  "fmt"
  "net"
)

func server() {
  // Initialize Listener
  ln, err := net.Listen("tcp", ":50010")
  if err != nil {
    fmt.Println("An error was encountered:", err)
  }
  for {
    // Accept connection
    c, err := ln.Accept()
    if err != nil {
      fmt.Println("An error was encountered:", err)
      continue
    }
    // Handle the connection
    go handleServerConnection(c)
  }
}

func handleServerConnection(c net.Conn) {
  // recieve the message
  var msg string
  err := gob.NewDecoder(c).Decode(&msg)
  if err != nil {
    fmt.Println("An error was encountered:", err)
  } else {
    fmt.Println("Recieved the following...", msg)
  }
  c.Close()
}

func client() {
  // connect to the server
  c, err := net.Dial("tcp", "127.0.0.1:50010")
  if err != nil {
    fmt.Println("An error was encountered:", err)
    return
  }
  // send the message
  msg := "My super long message"
  fmt.Println("Sending:... ", msg)
  err = gob.NewEncoder(c).Encode(msg)
  if err != nil {
    fmt.Println(err)
  }
  c.Close()
}

func main() {
  go server()
  go client()

  var input string
  fmt.Scanln(&input)
}
