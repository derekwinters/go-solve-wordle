package main

import (
  "fmt"
  "strconv"
  "math/rand"
  "time"
)

type wordScore struct {
  word           string
  positionScores []int
  totalScore     int
}

func main() {
  // Load dictionary
  fmt.Println("Loading dictionary...")
  remainingWords := readWordList("dictionary_wordle_valid.txt", 5)
  fmt.Println("Loaded " + strconv.Itoa(len(remainingWords)) + " words")

  // Pick a random word
  possibleAnswers := readWordList("dictionary_wordle_answers.txt", 5)
  rand.Seed(time.Now().Unix())
  answer := possibleAnswers[rand.Intn(len(possibleAnswers))]
  fmt.Println("Chose answer: " + answer)

  // Variables
  var wordScores []wordScore
  var nextGuess string
  var result string

  for i := 0; i < 6; i++ {
    // Start
    fmt.Println("--------------------------------------------------------------------------------")
    fmt.Println("Guess " + strconv.Itoa(i+1))
    fmt.Println("--------------------------------------------------------------------------------")
    // Calculate Word Scores
    fmt.Println("  Words Left: " + strconv.Itoa(len(remainingWords)))
    wordScores = scoreAllWords(remainingWords)

    // Get Best Word
    // implement sort, this is inefficient
    nextGuess = getBestWord(wordScores)
    fmt.Println("  Best Guess: " + nextGuess)

    // Make Guess, Get Result
    result = getGuessResult(nextGuess, answer)
    fmt.Println("  Result:     " + result)
    // Exit Loop if Correct
    if result == nextGuess {
      break
    }

    // Filter remainingWords based on new result information
    remainingWords = filterRemainingWords(result, nextGuess, remainingWords)
  }

  fmt.Println("--------------------------------------------------------------------------------")
  fmt.Println("Finished")
  fmt.Println("--------------------------------------------------------------------------------")
  if result == nextGuess {
    fmt.Println("  Result:     CORRECT!")
  } else {
    fmt.Println("  Result:     I LOST!")
  }
}
