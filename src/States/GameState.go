package States

import (
  // "fmt"
  "math/rand"
  "Entity"
  "Stacy"
  "Model"
  "Gordon"

  "github.com/go-gl/glfw/v3.2/glfw"
  "github.com/go-gl/gl/v4.1-core/gl"
  // "github.com/go-gl/mathgl/mgl32"
)

type GameState struct {
  time, elapsed, previousTime float32
  angle float64
  modelUniform int32
  texture uint32
  model *Model.Model
  shader uint32
  gordon *Gordon.Gordon

  ticks uint32
  entities []*Entity.Entity
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

  game.w = width
  game.h = height

  game.manager = manager
  game.entities = make([]*Entity.Entity, 100)
  for i := 0; i < cap(game.entities); i++ {
    x := (rand.Float32() * 20) - 10
    y := (rand.Float32() * 20) - 10
    z := (rand.Float32() * 20) - 10

    x_speed := float32(0) // rand.Float32() + 1
    y_speed := float32(0) // rand.Float32() + 1
    z_speed := float32(0) // rand.Float32() + 1
    game.entities[i] = new(Entity.Entity).Init(x, y, z, x_speed, y_speed, z_speed, shader, game.model)
  }

  game.ticks = 0

  return game
}

func (game *GameState) Update(elapsed float64) {
  game.time = float32(glfw.GetTime())
  game.elapsed = game.time - game.previousTime
  game.previousTime = game.time

  game.gordon.Update(game.elapsed)

  for _, ent := range game.entities {
    ent.Update(game.elapsed)
  }

  game.ticks += 1
}

func (game *GameState) Draw() {
  gl.UseProgram(game.manager.shader_program)

  gl.BindVertexArray(game.manager.vao)

  gl.ActiveTexture(gl.TEXTURE0)
  gl.BindTexture(gl.TEXTURE_2D, game.manager.texture)
  for _, ent := range game.entities {
    ent.Draw(game.modelUniform)
  }
}

func (game *GameState) Stop() bool {
  return true
}
