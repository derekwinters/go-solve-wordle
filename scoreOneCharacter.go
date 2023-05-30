package main

import (
)

// Score a specific character in a word by using the freqency that occurs in
// the list
func scoreOneCharacter(char string, positionFrequency []string) int {
  count := 0
  for item := 0; item < len(positionFrequency); item++ {
    if positionFrequency[item] == char {
      count++
    }
  }

  return count
}
