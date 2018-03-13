package HUD

import (
  "github.com/go-gl/mathgl/mgl32"
)

type element interface {
  SetPosition(loc mgl32.Vec2)
  Update(changes map[string]float32)
  Draw()
  Show()
  Width() float32
  Height() float32
}
