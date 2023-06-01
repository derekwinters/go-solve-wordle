package main

import (
  "bufio"
  "fmt"
  "os"
)

func getUserGuess(wordLength int) string {
  var guess string = ""

  for len(guess) != wordLength {

    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("  Enter guess: ")
    scanner.Scan()
    guess = scanner.Text()
  }

  return guess
}
