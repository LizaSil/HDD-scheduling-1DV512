package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

var InstructionFilePath string

func ReadInstructions() (instructions [][]int) {

  fmt.Println("Select instruction file from current directory")
  fmt.Scanln(&InstructionFilePath)
  InstructionFilePath = strings.Trim(InstructionFilePath, "'")

  //Check for txt extension
  if !strings.HasSuffix(InstructionFilePath, ".in") {
    InstructionFilePath += ".in"
  }

  file, err := os.Open(InstructionFilePath)
  if err != nil {
    fmt.Println("Error opening file")
    fmt.Println(InstructionFilePath)
    panic(err)
  }
  defer func(file *os.File) {
    err := file.Close()
    if err != nil {
      panic(err)
    }
  }(file)
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := strings.Split(scanner.Text(), ";")
    var lineInt []int
    for i, value := range line {
      if i == len(line)-1 {
        // comma separated numbers
        values := strings.Split(value, ",")
        for _, value := range values {
          valueInt, err := strconv.Atoi(value)
          if err != nil {
            panic(err)
          }
          lineInt = append(lineInt, valueInt)

        }
        break

      }

      valueInt, err := strconv.Atoi(value)
      if err != nil {
        panic(err)
      }
      lineInt = append(lineInt, valueInt)
    }
    instructions = append(instructions, lineInt)
  }
  return
}

func Output(results string) {
  // Add .out.txt extension
  InstructionFilePath = strings.TrimSuffix(InstructionFilePath, ".in")
  InstructionFilePath += ".out"

  file, err := os.Create(InstructionFilePath)
  if err != nil {
    panic(err)
  }

  defer func(file *os.File) {
    err := file.Close()
    if err != nil {
      panic(err)
    }
  }(file)

  if _, err = file.WriteString(results); err != nil {
    panic(err)
  }
}
