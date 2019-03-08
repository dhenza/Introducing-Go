package main

// Direction added to channels
// Buffered channel used

import (
  "fmt"
  "time"
)

func streamer (c chan<- string) {
  for i := 0; ; i++ {
    c <- "ping"
  }
}

func streamer2 (c chan<- string) {
  for i := 0; ; i++ {
    c <- "pong"
  }
}

func display(c<- chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second *1)
  }
}

func main() {
  var c chan string = make(chan string, 2)
  go streamer(c)
  go streamer2(c)
  go display(c)

  var input string
  fmt.Scanln(&input)
}
