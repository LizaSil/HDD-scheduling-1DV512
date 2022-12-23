package main

func FCFS(numberOfCylinders int, startCylinder int, cylindersToVisit []int) (numberOfMovements int, visitedCylinder []int) {
  var cylinder int
  visitedCylinder = append(visitedCylinder, startCylinder)

  // Calculate distance of current cylinder to start cylinder
  for i := 0; i < len(cylindersToVisit); i++ {
    cylinder = cylindersToVisit[i]
    numberOfMovements += func() int {
      if cylinder > startCylinder {
        return cylinder - startCylinder
      }
      return startCylinder - cylinder
    }()
    // Set the current cylinder as the start cylinder
    startCylinder = cylinder
    visitedCylinder = append(visitedCylinder, cylinder)
  }
  return
}

func SCAN(numberOfCylinders int, startCylinder int, cylindersToVisit []int, direction int) (numberOfMovements int, visitedCylinder []int) {
  var cylinder int
  start, end := []int{}, []int{}

  // Set the start and end cylinders
  if direction == 1 {
    end = append(end, numberOfCylinders-1)
  } else {
    start = append(start, 0)
  }

  for i := range cylindersToVisit {
    cylinder = cylindersToVisit[i]
    if cylinder < startCylinder {
      start = append(start, cylinder)
    } else if cylinder > startCylinder {
      end = append(end, cylinder)
    }
  }

  for i := 0; i < len(start); i++ {
    for j := i + 1; j < len(start); j++ {
      if start[i] > start[j] {
        start[i], start[j] = start[j], start[i]
      }
    }
  }

  for i := 0; i < len(end); i++ {
    for j := i + 1; j < len(end); j++ {
      if end[i] > end[j] {
        end[i], end[j] = end[j], end[i]
      }
    }
  }

  if direction == 1 {
    // Scan right
    return ScanEnd(cylinder, end, start, startCylinder)
  } else {
    // Scan left
    return ScanStart(cylinder, end, start, startCylinder)

  }

}

func CSCAN(numberOfCylinders int, startCylinder int, cylindersToVisit []int) (numberOfMovements int, visitedCylinder []int) {
  // Add the start cylinder to the visited cylinders
  visitedCylinder = append(visitedCylinder, startCylinder)

  var cylinder int
  start, end := []int{}, []int{}

  start = append(start, 0)
  end = append(end, numberOfCylinders-1)

  for i := 0; i < len(cylindersToVisit); i++ {
    cylinder = cylindersToVisit[i]
    if cylinder < startCylinder {
      start = append(start, cylinder)
    } else if cylinder > startCylinder {
      end = append(end, cylinder)
    }
  }

  for i := 0; i < len(start); i++ {
    for j := i + 1; j < len(start); j++ {
      if start[i] > start[j] {
        start[i], start[j] = start[j], start[i]
      }
    }
  }

  for i := 0; i < len(end); i++ {
    for j := i + 1; j < len(end); j++ {
      if end[i] > end[j] {
        end[i], end[j] = end[j], end[i]
      }
    }
  }

  // Circular SCAN begins scanning end first
  for i := 0; i < len(end); i++ {
    cylinder = end[i]
    // get the distance from the current cylinder to the next cylinder
    numberOfMovements += func() int {
      if cylinder > startCylinder {
        return cylinder - startCylinder
      }
      return startCylinder - cylinder
    }()

    startCylinder = cylinder
    visitedCylinder = append(visitedCylinder, cylinder)
  }
  numberOfCylinders = numberOfCylinders - 1

  // Reverse direction
  numberOfMovements += numberOfCylinders
  // Start from beginning
  startCylinder = 0

  // Scan start
  for i := 0; i < len(start); i++ {
    cylinder = start[i]
    numberOfMovements += func() int {
      if cylinder > startCylinder {
        return cylinder - startCylinder
      }
      return startCylinder - cylinder
    }()
    startCylinder = cylinder
    visitedCylinder = append(visitedCylinder, cylinder)
  }

  return
}
