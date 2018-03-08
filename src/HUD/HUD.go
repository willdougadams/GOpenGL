package HUD

import (
  "os"
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/4ydx/gltext"
	"github.com/4ydx/gltext/v4.1"
	"golang.org/x/image/math/fixed"
)

type HUD struct {
  text *v41.Text
}

func (hud *HUD) Init() *HUD {
  config, err := gltext.LoadTruetypeFontConfig("fontconfigs", "roboto")
  var font *v41.Font

	if err == nil {
		font, err = v41.NewFont(config)
		if err != nil {
			panic(err)
		}
		fmt.Println("Font loaded from disk...")
	} else {
		fd, err := os.Open("res/fonts/Roboto-Light.ttf")
		if err != nil {
			panic(err)
		}
		defer fd.Close()

		runeRanges := make(gltext.RuneRanges, 0)
		runeRanges = append(runeRanges, gltext.RuneRange{Low: 0x0020, High: 0x0080})

		scale := fixed.Int26_6(24)
		runesPerRow := fixed.Int26_6(128)
		config, err = gltext.NewTruetypeFontConfig(fd, scale, runeRanges, runesPerRow)
		if err != nil {
			panic(err)
		}
		config.Name = "roboto"

		err = config.Save("fontconfigs")
		if err != nil {
			panic(err)
		}
		font, err = v41.NewFont(config)
		if err != nil {
			panic(err)
		}
	}

	font.ResizeWindow(float32(1200), float32(900))

	scaleMin, scaleMax := float32(1.0), float32(1.1)
	hud.text = v41.NewText(font, scaleMin, scaleMax)
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	hud.text.SetString(str)
	hud.text.SetColor(mgl32.Vec3{0, 0, 0})
	hud.text.SetPosition(mgl32.Vec2{0, 0})
	hud.text.FadeOutPerFrame = 0.01

  return hud
}

func (hud *HUD) Update(elapsed float32) {

}

func (hud *HUD) Draw() {
  hud.text.Draw()
  hud.text.Show()
}
