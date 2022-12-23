package main

// Scans disk list from end and reverses direction
func ScanEnd(cylinder int, end []int, start []int, startCylinder int) (numberOfMovements int, visitedCylinder []int) {
  moves := 0
  for i := 0; i < len(end); i++ {
    cylinder = end[i]
    moves = func() int {
      if cylinder > startCylinder {
        return cylinder - startCylinder
      }
      return startCylinder - cylinder
    }()
    numberOfMovements += moves
    startCylinder = cylinder
    visitedCylinder = append(visitedCylinder, cylinder)
  }

  for i := len(start) - 1; i >= 0; i-- {
    cylinder = start[i]
    moves = func() int {
      if cylinder > startCylinder {
        return cylinder - startCylinder
      }
      return startCylinder - cylinder
    }()
    numberOfMovements += moves
    startCylinder = cylinder
    visitedCylinder = append(visitedCylinder, cylinder)
  }
  return numberOfMovements, visitedCylinder
}

// Scans disk list from start and reverses direction
func ScanStart(cylinder int, end []int, start []int, startCylinder int) (numberOfMovements int, visitedCylinder []int) {
  moves := 0
  for i := len(start) - 1; i >= 0; i-- {
    cylinder = start[i]
    moves = func() int {
      if cylinder > startCylinder {
        return cylinder - startCylinder
      }
      return startCylinder - cylinder
    }()
    numberOfMovements += moves
    startCylinder = cylinder
    visitedCylinder = append(visitedCylinder, cylinder)
  }

  for i := 0; i < len(end); i++ {
    cylinder = end[i]
    moves = func() int {
      if cylinder > startCylinder {
        return cylinder - startCylinder
      }
      return startCylinder - cylinder
    }()
    numberOfMovements += moves
    startCylinder = cylinder
    visitedCylinder = append(visitedCylinder, cylinder)
  }
  return numberOfMovements, visitedCylinder

}
