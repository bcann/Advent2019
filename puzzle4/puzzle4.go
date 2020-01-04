package main

import (
    "fmt"
    "strconv"
    // "bufio"
)

func HasTwins(number string) bool {
  prev := '0'
  duplicate_count := 0
  for _, x := range number {
    if x == prev {
      duplicate_count++
    } else if (x > prev) && (duplicate_count == 1) {
      return true
    } else if x > prev {
      duplicate_count = 0
    }
    prev = x
  }
  return duplicate_count == 1
}

func main() {
  lower_bound := 357253
  upper_bound := 892942
  possible_passwords := 0
  for i := lower_bound; i < upper_bound; i++ {
    pass := true
    num_str := strconv.Itoa(i) // gives me an array to work on
    var prev rune = '0'
    for _, x := range num_str {
      if x >= prev {
        prev = x
      } else {
        pass = false
        break
      }
    }
    if pass && HasTwins(num_str) {
      fmt.Println(i)
      possible_passwords++
    }
  }
  fmt.Println(possible_passwords)
}
