package States

import (
	"os"
	"fmt"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/4ydx/gltext"
	"github.com/4ydx/gltext/v4.1"
	"golang.org/x/image/math/fixed"
)

type MenuState struct {
	options []string
	manager *StateManager
	ticks int

	font *v41.Font
	text *v41.Text
}


func (menu *MenuState) Init(manager *StateManager, width int, height int, window *glfw.Window) State {
	menu.manager = manager
	menu.options = append(menu.options, "Resume")
	menu.options = append(menu.options, "Quit")
	menu.ticks = 0

	//=========================
	config, err := gltext.LoadTruetypeFontConfig("fontconfigs", "font_1_honokamin")
	if err == nil {
		menu.font, err = v41.NewFont(config)
		if err != nil {
			panic(err)
		}
		fmt.Println("Font loaded from disk...")
	} else {
		fd, err := os.Open("res/fonts/font_1_honokamin.ttf")
		if err != nil {
			panic(err)
		}
		defer fd.Close()

		// Japanese character ranges
		// http://www.rikai.com/library/kanjitables/kanji_codes.unicode.shtml
		runeRanges := make(gltext.RuneRanges, 0)
		runeRanges = append(runeRanges, gltext.RuneRange{Low: 0x3000, High: 0x3030})
		runeRanges = append(runeRanges, gltext.RuneRange{Low: 0x3040, High: 0x309f})
		runeRanges = append(runeRanges, gltext.RuneRange{Low: 0x30a0, High: 0x30ff})
		runeRanges = append(runeRanges, gltext.RuneRange{Low: 0x4e00, High: 0x9faf})
		runeRanges = append(runeRanges, gltext.RuneRange{Low: 0xff00, High: 0xffef})

		scale := fixed.Int26_6(24)
		runesPerRow := fixed.Int26_6(128)
		config, err = gltext.NewTruetypeFontConfig(fd, scale, runeRanges, runesPerRow)
		if err != nil {
			panic(err)
		}
		config.Name = "font_1_honokamin"

		err = config.Save("fontconfigs")
		if err != nil {
			panic(err)
		}
		menu.font, err = v41.NewFont(config)
		if err != nil {
			panic(err)
		}
	}

	menu.font.ResizeWindow(float32(width), float32(height))

	scaleMin, scaleMax := float32(1.0), float32(1.1)
	menu.text = v41.NewText(menu.font, scaleMin, scaleMax)
	str := "梅干しが大好き。ウメボシガダイスキ。"
	menu.text.SetString(str)
	menu.text.SetColor(mgl32.Vec3{0, 0, 0})
	menu.text.SetPosition(mgl32.Vec2{0, 0})
	menu.text.FadeOutPerFrame = 0.01
	//=========================

	return menu
}

func (menu *MenuState) Update(elapsed float32) {
	if menu.manager.window.GetKey(glfw.KeyT) == glfw.Press {
		menu.manager.ChangeState()
	}
	menu.ticks += 1
}

func (menu *MenuState) Draw() {
	menu.text.Draw()
	menu.text.Show()
}

func (menu *MenuState) Stop() bool {
	// Literally do nothing
	return true
}
