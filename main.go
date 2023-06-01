package main

import (
  "flag"
  "fmt"
  "regexp"
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
  flagDictionary := flag.String("dictionary", "", "Select this dictionary for both guesses and answers")
  flagGuessDict := flag.String("guess-dictionary", "wordle-valid", "Select guesses from wordle-valid, wordle-answers, scrabble")
  flagAnswerDict := flag.String("answer-dictionary", "wordle-answers", "Select answers from wordle-valid, wordle-answers, scrabble")
  flagWordLength := flag.Int("word-length", 5, "Select the length of the answer word")
  flag.Parse()

  mode := *flagMode
  answer := *flagAnswer
  dictionary := *flagDictionary
  guessDict := *flagGuessDict
  answerDict := *flagAnswerDict
  wordLength := *flagWordLength

  if dictionary != "" {
    guessDict = dictionary
    answerDict = dictionary
  }

    if mode != "demo" && mode != "play" {
      fmt.Println("  Invalid mode: " + mode )
      os.Exit(1)
    }

  fmt.Println("  Mode: " + mode)

  // Load dictionary
  if wordLength != 5 {
    if guessDict != "kids" && answerDict != "kids" {
      guessDict = "scrabble"
      answerDict = "scrabble"
    }
  }
  validWordList := readWordList("dictionaries/" + guessDict + ".txt", wordLength)
  remainingWords := validWordList
  fmt.Println("  Loaded " + strconv.Itoa(len(remainingWords)) + " words from dictionary")

  // Pick a random word
  possibleAnswers := readWordList("dictionaries/" + answerDict + ".txt", wordLength)
  rand.Seed(time.Now().Unix())

  switch mode {
    case "demo":
      if answer == "" {
        answer = possibleAnswers[rand.Intn(len(possibleAnswers))]
      }
      fmt.Println("  Selected word: " + answer)
    case "play":
      answer = possibleAnswers[rand.Intn(len(possibleAnswers))]
      repl := regexp.MustCompile(`\w`)
      fmt.Println("  Selected word: " + repl.ReplaceAllString(answer,"*"))
    case "solve":
      fmt.Println("  SOLVE feature not yet available.")
      os.Exit(0)
    default:
      fmt.Println("  Invalid mode: " + mode )
      os.Exit(1)
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
    fmt.Println("  Words Left:  " + strconv.Itoa(len(remainingWords)))
    wordScores = scoreAllWords(remainingWords, wordLength)

    if mode == "play" || mode == "solve" {
      // Get user guess
      nextGuess = getUserGuess(wordLength)

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
      fmt.Println("  Best Guess:  " + nextGuess)
    }

    if mode == "solve" {

    } else {
      // Make Guess, Get Result
      result = getGuessResult(nextGuess, answer, wordLength)
      fmt.Println("  Result:      " + result)
      // Exit Loop if Correct
      if result == nextGuess {
        break
      }
    }

    // Filter remainingWords based on new result information
    remainingWords = filterRemainingWords(result, nextGuess, remainingWords, wordLength)
  }

  fmt.Println("--------------------------------------------------------------------------------")
  fmt.Println("Finished")
  fmt.Println("--------------------------------------------------------------------------------")
  if result == nextGuess {
    fmt.Println("  Result:      CORRECT!")
  } else {
    fmt.Println("  Result:      LOST!")
    fmt.Println("  Answer:      " + answer)
  }
}
