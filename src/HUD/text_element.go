package HUD

import (
  "fmt"
  "strconv"
  "github.com/go-gl/mathgl/mgl32"
  "github.com/4ydx/gltext/v4.1"
)

type text_element struct {
  hud HUD
  location *mgl32.Vec2
  color *mgl32.Vec3
  float_formatting rune

  Text *v41.Text
  font *v41.Font
}

func (elem *text_element) Init(hud *HUD, font *v41.Font, float_formatting rune) *text_element {
	scaleMin, scaleMax := float32(1.0), float32(1.1)
  elem.float_formatting = float_formatting
	text := v41.NewText(font, scaleMin, scaleMax)
	str := "Frames per second: 0.00000"
	text.SetString(str)
	text.SetColor(mgl32.Vec3{0, 0, 0})
	text.SetPosition(mgl32.Vec2{0, 0})
	text.FadeOutPerFrame = 0.01

  elem.Text = text

  return elem
}

func (elem *text_element) Update(changes map[string]float32) {
  float_string := strconv.FormatFloat(float64(changes["new_fps"]), byte(elem.float_formatting), -1, 32)
  elem.Text.SetString(fmt.Sprintf("Frames per second: %s", float_string))

  /*
  x_pos := elem.Text.X2.X // -float32(elem.hud.w/2) + (elem.Text.Width()/2)
  y_pos := elem.Text.X2.Y
  elem.Text.SetPosition(mgl32.Vec2{x_pos, y_pos})
  */
}

func (elem *text_element) Draw() {
  elem.Text.Draw()
  elem.Text.Show()
}

func (elem *text_element) Show() {
  elem.Text.Show()
}

func (elem *text_element) SetPosition(loc mgl32.Vec2) {
  elem.Text.SetPosition(loc)
}

func (elem *text_element) Width() float32 {
  return elem.Text.Width()
}

func (elem *text_element) Height() float32 {
  return elem.Text.Height()
}
