package States

import (
	// "Gordon"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type State interface {
	Init(manager *StateManager, width int, height int, shader uint32, modelUniform int32, window *glfw.Window) State
	Update(elapsed float32)
	Draw()
	Stop() bool
}
