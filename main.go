package main

import (
	"flag"
	"fmt"
	"reflect"
	"strings"
)

var n = flag.Bool("n", false, "omit trainling newline")
var sep = flag.String("s", " ", "seperator")

// type conversion
// Fahrenheit(x) is a type conversion to conform with the return value
// This is possible because Cecius and Farenheit have the same underlying type: float64
// 
  type Celcius float64
  type Fahrenheit float64
  const (
    AbsoluteZeroC Celcius = -88
    FreezingC Celcius = 0
    BoilingC Celcius = 100
  )

// conversion not real
func CToF(c Celcius) Fahrenheit { return Fahrenheit(c + 32) }
func FToC(f Fahrenheit) Celcius { return Celcius(f - 32) }

func init() {
  fmt.Println("Called Implicitly")
}

// notice how the flag variables are being accessed as pointers.
func main(){
  flag.Parse()
  fmt.Print(strings.Join(flag.Args(), *sep))
  if !*n {
    fmt.Println()
  }

  // yes we can use the new function to create variable, but it returns a pointer
  var p = new(int)
  var q int
// 255 is the biggest number that can be represented in 8 bits (0-255)
  var m uint8 = 255

  fmt.Println(m)
  fmt.Println(reflect.TypeOf(p))
  fmt.Println(&q)

  medals := []string{"gold", "silver", "bronze"}
  fmt.Println("medals", medals, reflect.TypeOf(medals))

  var fahren = CToF(6)
  fmt.Println(reflect.TypeOf(fahren))
  fmt.Println(BoilingC-FreezingC) // both are Celcius
  // fmt.Println(BoilingC-CToF(8)) // type mismatch Celcius & F

  var f Fahrenheit
  var c Celcius
  fmt.Println(f == 0, c == 0)
  fmt.Println(c == Celcius(f)) // value of f is 0, just changing its type to Celcius
}