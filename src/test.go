package main

import (
  // "fmt"
  "States"
)

func main() {
  man := new(States.StateManager).Init(600, 800)
  for true {
    man.Update(0.0)
    man.Draw()
  }
}
