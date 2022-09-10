package main

import (
	"fmt"
  "os"
  "log"
)

func f() {}
var g = "g"

var cwd string

func init(){
  //WRONG: the short var declaration will re-declare cwd as a local variable, the package-level cwd is not initialiezed
  // since we are not using the second cwd declaration, we'll get a compiler error
  // cwd, err := os.Getwd()
  // if err != nil {
  //   log.Fatalf("os.Getwd failed: %v", err, cwd)
  // }

  // RIGHT: instead avoid := 
  var err error
  cwd, err = os.Getwd()
  if err != nil {
    log.Fatalf("os.Getwd failed: %v", err)
  }
}

func main() {
  f := "f" // overrides/shadows func f
  fmt.Println(g, f)

  // scope
  x := "hello!"
  // the for statement creates an implicit block for var i
  // the for loop body is an explicit block
  for i := 0; i < len(x); i++ {
    x := x[i]
    if x != '!' {
      x := x + 'A' - 'a'
      fmt.Printf("%c", x)
    }
  }

  for _, x := range x {
    x := x + 'A' - 'a'
    fmt.Printf("%c", x)
  }
}