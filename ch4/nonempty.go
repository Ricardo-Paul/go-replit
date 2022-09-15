package main

import(
  "fmt"
)

func nonempty(strings []string) []string{
  // var st []string
  counter := 0
  i := 0
  for _, s := range strings {
    if s != "" {
      fmt.Println(s)
      counter++
      strings[i] = s
      i++
    }
  }
  fmt.Println(counter)
  return strings
}

func main(){
  var n = nonempty([]string{ "Hello", "","","", "world" })
  fmt.Println(n)
}