package States

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type StateManager struct {
	prev_states []State
	curr_state State
	w, h int
	shader_program, vao, texture uint32

	window *glfw.Window
}

func (mngr *StateManager) Init(width int,
	height int,
	shader uint32,
	vao uint32,
	texture uint32,
	modelUniform int32,
	window *glfw.Window) *StateManager {
	mngr.texture = texture
	mngr.shader_program = shader
	mngr.vao = vao
	mngr.prev_states = make([]State, 0)
	mngr.window = window
	mngr.curr_state = new(GameState).Init(mngr, width, height, modelUniform, window)
	mngr.prev_states = append(mngr.prev_states, new(MenuState).Init(mngr, width, height, modelUniform, window))

	return mngr
}

func (mngr *StateManager) Update(elapsed float32) {
	mngr.curr_state.Update(elapsed)
}

func (mngr *StateManager) Draw() {
	mngr.curr_state.Draw()
}

func (mngr *StateManager) ChangeState() {
	mngr.prev_states = append(mngr.prev_states, mngr.curr_state)
	mngr.curr_state = mngr.prev_states[0]
	mngr.prev_states = append(mngr.prev_states[:0], mngr.prev_states[1:]...)
}

func (mngr *StateManager) ReturnToLastState() bool {
	if len(mngr.prev_states) == 0 {
		return false
	}
	length := len(mngr.prev_states)
	new_state := mngr.prev_states[length - 1]
	mngr.curr_state = new_state
	mngr.prev_states = mngr.prev_states[:length - 1]
	return true
}
