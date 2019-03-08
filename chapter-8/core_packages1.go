package main

import (
  "fmt"
  "strings"
)

func main() {
  // Confirm if a smaller string is contained within a larger string
  fmt.Println(strings.Contains("Bob", "ob")) // returns true

  // Count number of occurances of the smaller string in a larger string
  fmt.Println(strings.Count("BobobobobB", "ob")) // returns 4

  // Confirm if larger string starts with a given prefix
  fmt.Println(strings.HasPrefix("BobobobobB", "ob")) // returns false
  fmt.Println(strings.HasPrefix("BobobobobB", "Bo")) // returns false

  // Confirm if larger string starts with a given suffix
  fmt.Println(strings.HasSuffix("BobobobobB", "ob")) // returns false
  fmt.Println(strings.HasSuffix("BobobobobB", "bB")) // returns false

  // Confirm the position of a smaller string in a larger string (First occurance only)
  fmt.Println(strings.Index("BobobobobB", "ob")) // returns 1 as this is the first occurance

  // String concatenation
  fmt.Println("Bob" + "-" + "Ted") // returns Bob-Ted
  fmt.Println(strings.Join([]string{"Bob", "Ted"}, "-")) // returns Bob-Ted

  // Repeating a string
  fmt.Println(strings.Repeat("Bob", 3)) // returns BobBobBob

  // Replacing characters in a string
  fmt.Println(strings.Replace("aaabbb", "a", "b", 2)) // returns bbabbb
  fmt.Println(strings.Replace("aaabbb", "a", "b", -1)) // returns bbbbbb, -1 repalces all instances

  // Splitting strings
  fmt.Println(strings.Split("a,b,c,d,e,f,g", ",")) // returns [a b c d e f g]

  // Changing all chacters to lower case
  fmt.Println(strings.ToLower("AAAA")) // returns aaaa

  // Changing all chacters to upper case
  fmt.Println(strings.ToUpper("aaaa")) // returns AAAA
}
