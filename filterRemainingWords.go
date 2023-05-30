package main

import (
  "strings"
)

func filterRemainingWords(result string, guess string, remainingWords []string) []string {
  var filteredWords = []string{}

  for w := 0; w < len(remainingWords); w++ {
    keepWord := true
    for c := 0; c < 5; c++ {
      if string(result[c]) == "-" {
        if strings.Contains(remainingWords[w], string(guess[c])) {
          // - is a character not in the answer at all. Remove words that have that
          // character in any position
          keepWord = false
          break
        }
      } else if string(result[c]) == "+" {
        // + character must be in the word, but not the current position
        if string(guess[c]) == string(remainingWords[w][c]) || !(strings.Contains(remainingWords[w], string(guess[c]))) {
          keepWord = false
          break
        }
      } else if !(string(guess[c]) == string(remainingWords[w][c])) {
        // characters must match
        keepWord = false
        break
      }
    }

    if keepWord == true {
      filteredWords = append(filteredWords, remainingWords[w])
    }
  }
  return filteredWords
}
