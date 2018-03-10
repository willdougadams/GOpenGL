package HUD

import (
  "fmt"
  "github.com/go-gl/mathgl/mgl32"
  "github.com/4ydx/gltext/v4.1"
)

type element interface {
  Update(changes map[string]string)
  Draw()
}

type text_element struct {
  location *mgl32.Vec2
  color *mgl32.Vec3

  text *v41.Text
  font *v41.Font
}

func (elem *text_element) Init(font *v41.Font) *text_element {
	scaleMin, scaleMax := float32(1.0), float32(1.1)
	text := v41.NewText(font, scaleMin, scaleMax)
	str := "Frames per second: 0.0"
	text.SetString(str)
	text.SetColor(mgl32.Vec3{0, 0, 0})
	text.SetPosition(mgl32.Vec2{0, 0})
	text.FadeOutPerFrame = 0.01

  elem.text = text

  return elem
}

func (elem *text_element) Update(changes map[string]string) {
  elem.text.SetString(fmt.Sprintf("Frames per second: %s", changes["new_fps"]))
}

func (elem *text_element) Draw() {
  elem.text.Draw()
  elem.text.Show()
}
