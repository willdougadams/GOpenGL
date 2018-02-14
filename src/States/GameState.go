package States

import (
	"math/rand"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"

	"Entity"
	"Stacy"
	"Model"
	"Gordon"
	"Landscape"
	"Shader"
	"Debugs"
)

type GameState struct {
	model_uniform int32
	texture, shader uint32
	model *Model.Model
	gordon *Gordon.Gordon

	ticks uint32
	entities []*Entity.Entity
	land *Landscape.Landscape
	manager *StateManager
	stacy *Stacy.Stacy
	w, h int
}

func (game *GameState) Init(manager *StateManager, width int, height int, window *glfw.Window) State {
	Debugs.Print("Initializing Game...\n")
	game.stacy = new(Stacy.Stacy).Init()
	var temp_err error
	game.shader, temp_err = Shader.NewProgram("src/shaders/default.vert", "src/shaders/default.frag")
	if temp_err != nil {
		panic(temp_err)
	}
	gl.UseProgram(game.shader)

	texture_uniform := gl.GetUniformLocation(game.shader, gl.Str("tex\x00"))
	gl.Uniform1i(texture_uniform, 0)
	gl.BindFragDataLocation(game.shader, 0, gl.Str("outputColor\x00"))
	game.texture, temp_err = Shader.NewTexture("res/music_box/music_box_d.png")
	if temp_err != nil {
		panic(temp_err)
	}

	game.gordon = new(Gordon.Gordon).Init(0.0, 0.0, 0.0, game.shader, width, height, window)
	game.model = new(Model.Model).Init("res/music_box/music_box.obj", game.shader)

	game.w = width
	game.h = height

	game.manager = manager
	game.entities = make([]*Entity.Entity, 0)
	for i := 0; i < 1; i++ {
		x := (rand.Float32() * 10) - 5
		y := (rand.Float32() * 10) - 5
		z := (rand.Float32() * 10) - 5
		x_speed := float32(0.0)
		y_speed := (rand.Float32() * 10)
		z_speed := float32(0.0)
		game.entities = append(game.entities, new(Entity.Entity).Init(x, y, z, x_speed, y_speed, z_speed, game.shader, game.model))
	}

	game.land = new(Landscape.Landscape).Init(game.shader)

	game.ticks = 0

	return game
}

func (game *GameState) Update(elapsed float32) {
	game.gordon.Update(elapsed)

	Entity.Physics(game.land, game.entities, elapsed)

	if game.manager.window.GetKey(glfw.KeyY) == glfw.Press {
		game.manager.ChangeState()
	}

	game.ticks += 1
}

func (game *GameState) Draw() {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, game.texture)
	gl.UseProgram(game.shader)

	for _, ent := range game.entities {
		ent.Draw(game.model_uniform)
	}

	game.land.Draw(game.model_uniform)
}

func (game *GameState) Stop() bool {
	return true
}
