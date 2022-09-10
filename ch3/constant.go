package main
import(
  "fmt"
  "time"
)

// declare a sequence of const
const (
  pi = 3.14
  e = 2.71
)
// 1 is assigned to b c, 4 to g f
const (
  a = 1
  b
  c
  d = 4
  g
  f
)

// using iota: a constant generator 0, 1, 2
type Weekday int
const (
  Sunday Weekday = iota
  Monday
  Tuesday
  Wednesday
  Thursday
  Friday
  Saturday
)

const noDelay time.Duration = 0
const timeout = 5 * time.Minute
func main(){
  fmt.Println(pi, e)
  fmt.Printf("%T %[1]v\n", noDelay)
  fmt.Printf("%T \n", timeout)
  fmt.Println(a, b, c, d, g, f)
  fmt.Println(Sunday, Monday, Tuesday)
}