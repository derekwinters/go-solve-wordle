package main

import (
)

func getPositionCharacters(wordList []string, position int) []string {
  var returnList []string

  for i := 0; i < len(wordList); i++ {
    returnList = append(returnList, string(wordList[i][position]))
  }

  return returnList
}
