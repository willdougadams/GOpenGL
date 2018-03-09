package HUD

import (
  "os"
	"fmt"

	// "github.com/go-gl/mathgl/mgl32"
	"github.com/4ydx/gltext"
	"github.com/4ydx/gltext/v4.1"
	"golang.org/x/image/math/fixed"
)

type HUD struct {
  font *v41.Font
  text *v41.Text
  text_elem element
}

func (hud *HUD) Init() *HUD {
  hud.font = config_font("luximr")
	hud.text_elem = new(text_element).Init(hud.font)

  return hud
}

func (hud *HUD) Update(elapsed float32) {

}

func (hud *HUD) Draw() {
  hud.text_elem.Draw()
}

func config_font(font_name string) (font *v41.Font) {
  config, err := gltext.LoadTruetypeFontConfig("fontconfigs", "luximr")

	if err == nil {
		font, err = v41.NewFont(config)
		if err != nil {
			panic(err)
		}
		fmt.Println("Font loaded from disk...")
	} else {
		fd, err := os.Open(fmt.Sprintf("res/fonts/%s.ttf", font_name))
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
		config.Name = font_name

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

  return font
}
