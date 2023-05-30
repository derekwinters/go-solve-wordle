package main

import (
  "bufio"
  "fmt"
  "os"
)

func readWordList(path string) []string {
  readFile, err := os.Open(path)

  if err != nil {
    fmt.Println(err)
  }

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
  var fileLines []string

  for fileScanner.Scan() {
    fileLines = append(fileLines, fileScanner.Text())
  }

  readFile.Close()

  return fileLines
}
