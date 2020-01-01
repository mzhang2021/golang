package main

import "fmt"

func split(a, b []int) [][]int {
  ret := make([][]int, 2)
  for _, i := range a {
    if contains(b, i) {
      ret[0] = append(ret[0], i)
    } else {
      ret[1] = append(ret[1], i)
    }
  }
  for _, i := range b {
    if contains(a, i) && !contains(ret[0], i) {
      ret[0] = append(ret[0], i)
    } else if !contains(a, i) && !contains(ret[1], i) {
      ret[1] = append(ret[1], i)
    }
  }
  return ret
}

func contains(a []int, val int) bool {
  for _, i := range a {
    if i == val {
      return true
    }
  }
  return false
}

func main() {
  a := split([]int{1, 2, 3, 4}, []int{1, 3, 5})
  for _, i := range a[0] {
    fmt.Print(i, " ")
  }
  fmt.Println()
  for _, i := range a[1] {
    fmt.Print(i, " ")
  }
  fmt.Println("\n----------------")

  b := split([]int{1, 2, 3}, []int{})
  for _, i := range b[0] {
    fmt.Print(i, " ")
  }
  fmt.Println()
  for _, i := range b[1] {
    fmt.Print(i, " ")
  }
  fmt.Println("\n----------------")

  c := split([]int{}, []int{1, 2, 3})
  for _, i := range c[0] {
    fmt.Print(i, " ")
  }
  fmt.Println()
  for _, i := range c[1] {
    fmt.Print(i, " ")
  }
  fmt.Println("\n----------------")

  d := split([]int{1, 2, 3, 4}, []int{5, 6, 7})
  for _, i := range d[0] {
    fmt.Print(i, " ")
  }
  fmt.Println()
  for _, i := range d[1] {
    fmt.Print(i, " ")
  }
  fmt.Println("\n----------------")

  e := split([]int{}, []int{})
  for _, i := range e[0] {
    fmt.Print(i, " ")
  }
  fmt.Println()
  for _, i := range e[1] {
    fmt.Print(i, " ")
  }
  fmt.Println("\n----------------")
}