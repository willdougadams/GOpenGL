package States

import (
  "fmt"
  "math/rand"

  "github.com/go-gl/glfw/v3.2/glfw"

  "Entity"
  "Stacy"
  "Model"
  "Gordon"
  "Landscape"
)

type GameState struct {
  angle float64
  modelUniform int32
  texture uint32
  model *Model.Model
  shader uint32
  gordon *Gordon.Gordon

  ticks uint32
  entities []*Entity.Entity
  land *Landscape.Landscape
  manager *StateManager
  stacy *Stacy.Stacy
  w, h int
}

func (game *GameState) Init(manager *StateManager,
                            width int,
                            height int,
                            shader uint32,
                            modelUniform int32,
                            window *glfw.Window) State {
    game.stacy = new(Stacy.Stacy).Init()
    game.shader = shader

    game.gordon = new(Gordon.Gordon).Init(0.0, 0.0, 0.0, shader, width, height, window)
    game.modelUniform = modelUniform
    game.model = new(Model.Model).Init("res/bunny.obj", shader)
    // land := new(Model.Model).Init("res/mountain/mount.obj", shader)

    game.w = width
    game.h = height

    game.manager = manager
    game.land = new(Landscape.Landscape).Init(game.shader)
    game.entities = make([]*Entity.Entity, 0)


    x := float32(0.0)
    y := float32(20.0)
    z := float32(15.0)

    x_speed := (rand.Float32() * 1) - 0.5
    y_speed := (rand.Float32() * 10)
    z_speed := (rand.Float32() * 1) - 0.5

    game.entities = append(game.entities, new(Entity.Entity).Init(x, y, z, x_speed, y_speed, z_speed, shader, game.model))
    //game.entities = append(game.entities, new(Entity.Entity).Init(0.0, 0.0, 0.0, 0.0, 0.0, 0.0, shader, land))

    game.ticks = 0

    fmt.Printf("Starting game...")
    return game
}

func (game *GameState) Update(elapsed float32) {
    game.gordon.Update(elapsed)

    for _, ent := range game.entities {
        ent.Update(elapsed)
    }

    game.ticks += 1
}

func (game *GameState) Draw() {
    // gl.UseProgram(game.manager.shader_program)
    // gl.BindVertexArray(game.manager.vao)
    // gl.ActiveTexture(gl.TEXTURE0)
    // gl.BindTexture(gl.TEXTURE_2D, game.manager.texture)
    for _, ent := range game.entities {
        ent.Draw(game.modelUniform)
    }

    game.land.Draw(game.modelUniform)
}

func (game *GameState) Stop() bool {
    return true
}
