package main

import (
	"fmt"
	"os"
)

// 1.1 modified program that shows the first argument
// 1.2 - modified program that shows the index
func main(){
  echoCli2()
}

func echoCli2() {
	// can range only be used inside a loop?
	// range returns an index and the element
	element, sep := "", ""
  fmt.Println("caller command :", os.Args[0:])
	for idx, arg := range os.Args[1:] {
		element += sep + arg
		sep = " "
	  fmt.Println(idx, element)
	}
}

// this will error because short var declaration can only
// be used within function
// s := ""

