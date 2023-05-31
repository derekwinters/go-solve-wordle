package main

import (
  "bufio"
  "fmt"
  "os"
)

func getUserGuess() string {
  var guess string = ""

  for len(guess) != 5 {

    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("  Enter guess: ")
    scanner.Scan()
    guess = scanner.Text()
  }

  return guess
}
