package main

import (
  "bufio"
  "fmt"
  "log"
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

  if len(fileLines) == 0 {
    log.Fatal("ERROR: No words found in dictionary file: " + path)
  }

  return fileLines
}
