package States

import (
  "fmt"
)

type MenuState struct {
  options []string
  manager *StateManager
}

/*
func (menu *MenuState) Init(manager *StateManager, width int, height int, shader uint32, modelUniform int32) State {
  menu.manager = manager
  menu.options = append(menu.options, "Resume")
  menu.options = append(menu.options, "Quit")
  return menu
}
*/

func (menu *MenuState) Update(elapsed float32) {
  // Literally do nothing
}

func (menu *MenuState) Draw() {
  fmt.Printf("=== Menu ===\n")
  for i, v := range menu.options {
    fmt.Printf("Option %d: %s\n", i, v)
  }
}

func (menu *MenuState) Stop() bool {
  // Literally do nothing
  return true
}
