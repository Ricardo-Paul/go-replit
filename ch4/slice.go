package main

import (
  "fmt"
  "bytes"
)

func main(){
  // byte slice
  var b []byte
  b = []byte("A")
  fmt.Println("byte", b)
  fmt.Printf("%T", b)

  fmt.Printf("Slice: %#v\n", b)

  var months = []string{0: "Nothing", 1: "January", 2: "February", 3: "March", 4: "April"}

  Q1 := months[:3]
  Q2 := months[2:4]
  fmt.Println(Q1, Q2)
  // find months in common (inefficient)
  for _, i := range Q1{
    for _, j := range Q2{
      if i == j{
        fmt.Println("common", i)
      }
    }
  }

  // comparing byte slices
  eqBytes := bytes.Equal([]byte{}, []byte{})
  fmt.Println("bytes.Equal",eqBytes)

  // compare string slices
  eqString := equal(Q1, Q2)
  fmt.Println("String slices equal",eqString)

  // compare slice against nil
  var s []int
  s = nil // nil
  s = []int(nil) // init with nil
  s = []int{} // not nil but length is zero
  fmt.Println("slice nil?",s == nil)

  // create slices with make
  var sl = make([]string, 3, 6)
  fmt.Printf("slice = %v \nlen = %d \ncap = %d \n",sl, len(sl), cap(sl))

  // append elements to slices
  // Fun Fact: range returns a 32-bit unicode code point for each byte in the string - basically a rune
  var runes []rune // a rune slice
  var strings []string
  for _, r := range "Hello World" {
    runes = append(runes, r)
    strings = append(strings, string(r))
  }

  fmt.Printf("rune slice = %v \nchars = %c \nstring slice = %v", runes, runes, strings)
  
}

// implement func for string slices comparison
// equal compares string slices
// a slice can contain itself: https://stackoverflow.com/questions/36077566/how-can-a-slice-contain-itself
func equal(x, y []string) bool {
  if len(x) != len(y) {
    return false
  }
  for i := range x {
    if x[i] != y[i] {
      return false
    }
  }
  return true
}