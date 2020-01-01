package main

import "fmt"

func printCandidateInfo(s string) {
  if len(s) == 0 {
    fmt.Println("No votes were cast.\n----------------")
    return
  }

  var winner rune
  winnerCount := 0
  voteCount := make(map[rune]int)
  for _, c := range s {
    voteCount[c]++
    if voteCount[c] > winnerCount {
      winner = c
      winnerCount = voteCount[c]
    }
  }

  for k, v := range voteCount {
    fmt.Println(string(k), "received", v, "votes.")
  }
  fmt.Println("Winner:", string(winner))
  for k, v := range voteCount {
    fmt.Println(string(k), "was", winnerCount - v, "votes away from winning.")
  }
  fmt.Println("----------------")
}

func main() {
  printCandidateInfo("ABCDEE")

  printCandidateInfo("ADECAFECA");

  printCandidateInfo("");

  var s string
  for c := 'A'; c <= 'E'; c++ {
    for i := 0; i < 1000; i++ {
      s += string(c)
    }
  }
  s += "D"
  printCandidateInfo(s)
}