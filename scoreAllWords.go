package main

import (
  "sort"
  "strings"
)

func scoreAllWords(wordList []string, wordLength int) []wordScore {
  var positionCharacters = [][]string{}

  // Get slice of characters in wordList for all five positions
  for i := 0; i < wordLength; i++ {
    positionCharacters = append(positionCharacters, getPositionCharacters(wordList, i))
  }

  // Create a slice of word scores to store results in
  var allScores = []wordScore{}

  // Loop through words and generate score
  for wordIndex := 0; wordIndex < len(wordList); wordIndex++ {
    // Create score struct
    newWordScore := wordScore{
      word:           wordList[wordIndex],
      positionScores: []int{},
      totalScore:     0,
    }

    // Get a score for each position
    for i := 0; i < wordLength; i++ {
      var characterScore = scoreOneCharacter(string(wordList[wordIndex][i]), positionCharacters[i])
      if strings.Count(wordList[wordIndex], string(wordList[wordIndex][i])) == 1 {
        newWordScore.totalScore +=characterScore
      }
      newWordScore.positionScores = append(newWordScore.positionScores, characterScore)
    }
    allScores = append(allScores, newWordScore)
  }

  sort.SliceStable(allScores, func(i, j int) bool {
    return allScores[i].totalScore > allScores[j].totalScore
  })

  return allScores
}

