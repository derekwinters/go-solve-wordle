package main

import (
  "strings"
)

func getGuessResult(guess string, answer string, wordLength int) string {
  var result string
  for i := 0; i < wordLength; i++ {
    if guess[i] == answer[i] {
      result = result + string(guess[i])
    } else if strings.Contains(answer, string(guess[i])) {
      result = result + "+"
    } else {
      result = result + "-"
    }
  }
  return result
}
