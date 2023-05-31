package main

import (
  "flag"
  "fmt"
  "strconv"
  "math/rand"
  "time"
  "os"
)

type wordScore struct {
  word           string
  positionScores []int
  totalScore     int
}

func main() {
  fmt.Println("--------------------------------------------------------------------------------")
  fmt.Println("Setup")
  fmt.Println("--------------------------------------------------------------------------------")
  flagMode := flag.String("mode", "demo", "Select the mode: demo, play, solve")
  flagAnswer := flag.String("answer", "", "Select a specific answer for demo mode")
  flagGuessDict := flag.String("guess-dictionary", "wordle-valid", "Select guesses from wordle-valid, wordle-answers, scrabble")
  flagAnswerDict := flag.String("answer-dictionary", "wordle-answers", "Select answers from wordle-valid, wordle-answers, scrabble")
  flag.Parse()

  mode := *flagMode
  answer := *flagAnswer
  guessDict := *flagGuessDict
  answerDict := *flagAnswerDict

  fmt.Println("  Mode: " + mode)

  // Load dictionary
  validWordList := readWordList("dictionaries/" + guessDict + ".txt", 5)
  remainingWords := validWordList
  fmt.Println("  Loaded " + strconv.Itoa(len(remainingWords)) + " words from dictionary")

  // Pick a random word
  possibleAnswers := readWordList("dictionaries/" + answerDict + ".txt", 5)
  rand.Seed(time.Now().Unix())

  if mode == "demo" {
    if  answer == "" {
      answer = possibleAnswers[rand.Intn(len(possibleAnswers))]
    }
    fmt.Println("  Selected word: " + answer)
  } else if mode == "play" {
    answer = possibleAnswers[rand.Intn(len(possibleAnswers))]
    fmt.Println("  Selected word: *****")
  } else if mode == "solve" {
    fmt.Println("  SOLVE feature not yet available.")
    os.Exit(0)
  } else {
    fmt.Println("  Invalid mode: " + mode )
    os.Exit(1)
  }

  if (mode == "demo" && answer == "") || mode == "play" {
    answer = possibleAnswers[rand.Intn(len(possibleAnswers))]
  }
  if mode != "play" {
    fmt.Println("  Chose answer: " + answer)
  }

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

    if mode == "play" || mode == "solve" {
      // Get user guess
      nextGuess = getUserGuess()

      // Show guess score
      for _, v := range wordScores {
        if v.word == nextGuess {
          fmt.Println("  Guess Score: " + strconv.Itoa(v.totalScore))
          break
        }
      }
    } else if mode == "demo" {
      // Get Best Word
      // implement sort, this is inefficient
      nextGuess = getBestWord(wordScores)
      fmt.Println("  Best Guess: " + nextGuess)
    }

    if mode == "solve" {

    } else {
      // Make Guess, Get Result
      result = getGuessResult(nextGuess, answer)
      fmt.Println("  Result:     " + result)
      // Exit Loop if Correct
      if result == nextGuess {
        break
      }
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
