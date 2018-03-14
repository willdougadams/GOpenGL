package HUD

import (
  "os"
	"fmt"
  "Debugs"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/4ydx/gltext"
	"github.com/4ydx/gltext/v4.1"
	"golang.org/x/image/math/fixed"
)

const ELEMENT_SPACER = 10

type HUD struct {
  font *v41.Font
  text *v41.Text
  elements []element

  w, h int
  counter float32
}

func (hud *HUD) Init(w, h int) *HUD {
  hud.font = config_font("luximr", w, h)
  formats := []rune{'b', 'e', 'E', 'f'}
  hud.elements = make([]element, 0)
  hud.w = w
  hud.h = h
  hud.counter = 0.0

  for i, format := range formats {
    t := new(text_element).Init(hud, hud.font, format)
    x_pos := -float32(w/2) + (t.Width()/2)
    y_pos := (float32(h/2) - (t.Height()/2)) - (float32(i) * (t.Height() + ELEMENT_SPACER))
    t.SetPosition(mgl32.Vec2{x_pos, y_pos})
    hud.elements = append(hud.elements, t)
  }

  return hud
}

func (hud *HUD) Update(elapsed float32, update_map map[string]float32) {
  if hud.counter < 0.75 {
    hud.counter += elapsed
    return
  } else {
    hud.counter = 0.0
  }

  for i, el := range hud.elements {
    el.Update(update_map)

    x_pos := -float32(hud.w/2) + (el.Width()/2)
    y_pos := (float32(hud.h/2) - (el.Height()/2)) - (float32(i) * (el.Height() + ELEMENT_SPACER))
    el.SetPosition(mgl32.Vec2{x_pos, y_pos})
  }
}

func (hud *HUD) Draw() {
  for _, el := range hud.elements {
    el.Draw()
  }

  for _, el := range hud.elements {
    el.Show()
  }
}

func config_font(font_name string, w, h int) (font *v41.Font) {
  config, err := gltext.LoadTruetypeFontConfig("fontconfigs", font_name)

	if err == nil {
		font, err = v41.NewFont(config)
		if err != nil {
			panic(err)
		}
		Debugs.Print(fmt.Sprintf("Font %s loaded from disk...", font_name))
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

	font.ResizeWindow(float32(w), float32(h))

  return font
}
