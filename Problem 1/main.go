package main

import "fmt"

func findSubstring(a, b string) []int {
  var ret []int
  if len(b) == 0 {
    return ret
  }
  for i := 0; i <= len(a) - len(b); i++ {
    substr := a[i:i+len(b)]
    if substr == b {
      ret = append(ret, i)
    }
  }
  return ret
}

func main() {
  a := findSubstring("googlygoggles", "googly")
  for _, i := range a {
    fmt.Print(i, " ")
  }
  fmt.Println()

  b := findSubstring("googlygoggles", "oo")
  for _, i := range b {
    fmt.Print(i, " ")
  }
  fmt.Println()

  c := findSubstring("googlygoggles", "g")
  for _, i := range c {
    fmt.Print(i, " ")
  }
  fmt.Println()

  d := findSubstring("mmmmm", "mm")
  for _, i := range d {
    fmt.Print(i, " ")
  }
  fmt.Println()

  e := findSubstring("asdf", "")
  for _, i := range e {
    fmt.Print(i, " ")
  }
  fmt.Println()

  f := findSubstring("", "")
  for _, i := range f {
    fmt.Print(i, " ")
  }
  fmt.Println()

  g := findSubstring("q", "qwer")
  for _, i := range g {
    fmt.Print(i, " ")
  }
  fmt.Println()

  h := findSubstring("qwer", "qwer")
  for _, i := range h {
    fmt.Print(i, " ")
  }
  fmt.Println()
}