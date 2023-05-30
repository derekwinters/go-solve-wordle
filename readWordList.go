package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func readWordList(path string, length int) []string {
  readFile, err := os.Open(path)

  if err != nil {
    fmt.Println(err)
  }

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
  var fileLines []string

  for fileScanner.Scan() {
    if len(fileScanner.Text()) == length {
      fileLines = append(fileLines, strings.ToLower(fileScanner.Text()))
    }
  }

  readFile.Close()

  return fileLines
}
