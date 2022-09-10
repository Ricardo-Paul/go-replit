package main

import (
  "fmt"
)

func main(){
  var slice []int = []int{}
  appendInt(slice, 4)

  // slices are reference type
  sliceA := []int{2, 4, 6, 8}
  sliceB := sliceA

  // sliceA contains a reference to the slice not the slice itself
  // the same reference is copied to sliceB

  sliceA[0] = 90
  sliceB[1] = 100

  // the same slice has been modified
  fmt.Println("Slice A :", sliceA)
  fmt.Println("Slice B :", sliceB)

  // arrays are value types
  arrayA := [2]int{10, 20}
  arrayB := arrayA

  // arrayA containts the array itself
  // the array is copied to arrayB
  arrayA[0] = 33
  arrayB[0] = 55

  fmt.Println("ArrayA", arrayA) // [33 20]
  fmt.Println("ArrayB", arrayB) // [53 20]
}

// when adding the first element to an empty slice
// the slice will be at cap = 0 and len = 0
func appendInt(x []int, y int){
  var z []int
  zlen := len(x) + 1 // zlen = 1
  if zlen <= cap(x){
    fmt.Println("there is room")
    z = x[:zlen]
    fmt.Println(z)
  } else {
    zcap := zlen // 1 == 1
    fmt.Printf("zlen = %v \n", zlen)
    fmt.Printf("2 * len(x) =  %v \n", 2*len(x))
    fmt.Println("zcap", zcap)
    if zlen < 2*len(x) {
      zcap = 2*len(x)
      fmt.Println("run ?")
    }

    z = make([]int, zlen, zcap) // cap = len =1 => [0]
  }
    copy(z, x) // x => [0]
    z[len(x)] = y // insert y at the last pos
    fmt.Println("z ", z)
}