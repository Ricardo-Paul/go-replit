package main

import (
  "fmt"
)

// using pointer to pass array by reference
// zero sets all bytes in the array to 0
func zero(ptr *[32]byte){
  for i := range ptr {
    ptr[i] = 0
  }
}

func main(){
  // var a [3]int = [3]int{ 6, 9, 12 } // OR
  // using the elipsis, the length is dertermined by the number of initializers
  var a = [...]int{ 6, 9, 12 }
  a[len(a) - 1] = 90 // re-assign last element
  fmt.Println(a[len(a) - 1]) // print last element

  // range is used to iterate over the array 
  for i, v := range a {
    fmt.Printf("%d %d\n", i, v)
  }
  fmt.Printf("%T", a)

  // the size of the array is part of its type
  // trying to assign type [4]int to [3]int will error
  // var q = [3]int{1, 2, 3}
  // q = [4]int{1, 2, 3, 4}

  type Currency int
  const (
    USD Currency = iota
    EUR
  )
  symbol := [...]string{USD :"Dollar$"}
  fmt.Println("\nSymbol",symbol[USD], "\n")

  // an array of 100 elements, last element value is -1
  r := [...]int{99: -1}
  fmt.Println(len(r), r[len(r) - 1])

  arr1 := [2]int{1, 2}
  arr2 := [...]int{1, 2}
  fmt.Println(arr1 == arr2)

  byteArray := [32]byte{}
  zero(&byteArray)
  fmt.Println(byteArray)
}

