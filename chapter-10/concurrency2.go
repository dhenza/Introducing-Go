package main

// Channels provide a way for two goroutines to communicate with each other and syn‐ chronize their execution.
// channel messages are recieved by all display functions initialized using c in the main function. Blocking occurs such
// that the single display in this case with print "ping" "pong"

import (
  "fmt"
  "time"
)

func streamer (c chan string) {
  for i := 0; ; i++ {
    c <- "ping"
  }
}

func streamer2 (c chan string) {
  for i := 0; ; i++ {
    c <- "pong"
  }
}

func display(c chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second *1)
  }
}

func main() {
  var c chan string = make(chan string)
  go streamer(c)
  go streamer2(c)
  go display(c)

  var input string
  fmt.Scanln(&input)
}
