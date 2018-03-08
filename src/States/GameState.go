package States

import (
	"math/rand"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"

	"Entity"
	"Model"
	"Gordon"
	"Landscape"
	"Debugs"
	"HUD"
)

type GameState struct {
	model_uniform int32
	texture, shader uint32
	models []*Model.Model
	gordon *Gordon.Gordon
	hud *HUD.HUD

	ticks uint32
	entities []*Entity.Entity
	land *Landscape.Landscape
	manager *StateManager
	w, h int
}

func (game *GameState) Init(manager *StateManager, width int, height int, window *glfw.Window) State {
	Debugs.Print("Initializing Game...\n")
	var temp_err error
	game.shader, temp_err = Model.NewProgram("src/shaders/default.vert", "src/shaders/default.frag")
	if temp_err != nil {
		panic(temp_err)
	}
	gl.UseProgram(game.shader)
	gl.BindFragDataLocation(game.shader, 0, gl.Str("outputColor\x00"))

	light_loc_uni := gl.GetUniformLocation(game.shader, gl.Str("light_location\x00"))
	gl.Uniform4f(light_loc_uni, 1.0, 100.0, 10.0, 1.0)

	game.gordon = new(Gordon.Gordon).Init(0.0, 0.0, 0.0, game.shader, width, height, window)
	game.land = new(Landscape.Landscape).Init(game.shader)

	game.models = make([]*Model.Model, 0)
	game.models = append(game.models, new(Model.Model).Init("res/music_box/music_box.obj", "res/music_box/music_box_d.png", game.shader))
	game.models = append(game.models, new(Model.Model).Init("res/wolf/Wolf_dae.dae", "res/wolf/textures/Wolf_Body.png", game.shader))

	game.hud = new(HUD.HUD).Init()

	game.w = width
	game.h = height

	game.manager = manager
	game.entities = make([]*Entity.Entity, 0)
	for i := 0; i < 2; i++ {
		x := (rand.Float32() * 10) - 5
		y := (rand.Float32() * 10) - 5
		z := (rand.Float32() * 10) - 5
		x_speed := float32(0.0)
		y_speed := (rand.Float32() * 10)
		z_speed := float32(0.0)
		game.entities = append(game.entities, new(Entity.Entity).Init(x, y, z, x_speed, y_speed, z_speed, game.shader, game.models[i%len(game.models)]))
	}

	game.ticks = 0

	return game
}

func (game *GameState) Update(elapsed float32) {
	gl.UseProgram(game.shader)

	game.gordon.Update(elapsed)

	Entity.Physics(game.land, game.entities, elapsed)

	if game.manager.window.GetKey(glfw.KeyY) == glfw.Press {
		game.manager.ChangeState()
	}

	game.ticks += 1
}

func (game *GameState) Draw() {
	gl.UseProgram(game.shader)

	game.land.Draw(game.model_uniform)

	for _, ent := range game.entities {
		ent.Draw(game.model_uniform)
	}

	game.hud.Draw()
}

func (game *GameState) Stop() bool {
	return true
}
