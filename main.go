package main

import (
  "fmt"
)

func main() {
  instructions := ReadInstructions()

  numberOfCylinder := instructions[0][0]
  startCylinder := instructions[1][0]
  destination := instructions[2][0]
  cylindersToVisit := instructions[3]

  FCFSmovements, FCFSvisit := FCFS(numberOfCylinder, startCylinder, cylindersToVisit)
  SCANmovements, SCANvisit := SCAN(numberOfCylinder, startCylinder, cylindersToVisit, destination)
  CSCANmovements, CSCANvisit := CSCAN(numberOfCylinder, startCylinder, cylindersToVisit)

  results :=
    "FCFS\nNumber of movements: " + fmt.Sprint(FCFSmovements) +
      "\nVisited cylinders: " + fmt.Sprint(FCFSvisit) +
      "\nSCAN\nNumber of movements: " + fmt.Sprint(SCANmovements) +
      "\nVisited cylinders: " + fmt.Sprint(SCANvisit) +
      "\nC-SCAN\nNumber of movements: " + fmt.Sprint(CSCANmovements) +
      "\nVisited cylinders: " + fmt.Sprint(CSCANvisit) + "\n"

  fmt.Println(results)
  Output(results)

}
