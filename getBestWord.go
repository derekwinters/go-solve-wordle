package main

import (
)

func getBestWord(wordScores []wordScore) string {
  var bestWord = ""
  var wordScore = 0

  for i := 0; i < len(wordScores); i++ {
    if wordScores[i].totalScore > wordScore {
      bestWord = wordScores[i].word
      wordScore = wordScores[i].totalScore
    }
  }
  return bestWord
}
