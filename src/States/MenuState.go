package States

import (
  "fmt"
  "os"
  // "strconv"

  "Shader"

  "github.com/nullboundary/glfont"
  "github.com/go-gl/glfw/v3.2/glfw"
  "github.com/go-gl/gl/v4.1-core/gl"
)

type MenuState struct {
  options []string
  manager *StateManager
  font *glfont.Font

  shader, vao, vbo uint32
}


func (menu *MenuState) Init(manager *StateManager,
                            width int,
                            height int,
                            windot *glfw.Window) *MenuState {
  menu.manager = manager
  menu.options = append(menu.options, "Resume")
  menu.options = append(menu.options, "Quit")
  var err error
  fmt.Printf("Loading font...\n")
  menu.font, err = glfont.LoadFont("res/fonts/Roboto-Light.ttf", int32(52), width, height)
  if err != nil {
      fmt.Printf("Failed to load font, exiting...\n")
      os.Exit(1)
  }

  var buffer_data []float32
  x := float32(50)
  y := float32(50)
  buffer_data = append(buffer_data, x)
  buffer_data = append(buffer_data, y)
  buffer_data = append(buffer_data, -1)

  buffer_data = append(buffer_data, x+10)
  buffer_data = append(buffer_data, y)
  buffer_data = append(buffer_data, -1)

  buffer_data = append(buffer_data, x)
  buffer_data = append(buffer_data, y+10)
  buffer_data = append(buffer_data, -1)

  buffer_data = append(buffer_data, x+10)
  buffer_data = append(buffer_data, x+10)
  buffer_data = append(buffer_data, -1)

  gl.GenVertexArrays(1, &menu.vao)
  gl.BindVertexArray(menu.vao)
  gl.GenBuffers(1, &menu.vbo)
  gl.BindBuffer(gl.ARRAY_BUFFER, menu.vbo)
  gl.BufferData(gl.ARRAY_BUFFER, 4, gl.Ptr(buffer_data), gl.STATIC_DRAW)

  var temp_err error
  menu.shader, temp_err = Shader.NewProgram("src/shaders/menu.vert", "src/shaders/menu.frag")
  if temp_err != nil {
    panic(temp_err)
  }

  vert_attrib := uint32(gl.GetAttribLocation(menu.shader, gl.Str("vert\x00")))
  gl.EnableVertexAttribArray(vert_attrib)
  gl.VertexAttribPointer(vert_attrib, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

  return menu
}

func (menu *MenuState) Update(elapsed float32) {

  // Literally do nothing
}

func (menu *MenuState) Draw() {
    gl.BindVertexArray(menu.vao)
    gl.BindBuffer(gl.ARRAY_BUFFER, menu.vbo)

    // gl.UniformMatrix4fv(model_uniform, 1, false, &entity_model[0])
    gl.DrawArrays(gl.TRIANGLES, 0, 4)
    //  fmt.Printf("=== Menu ===\n")
    for i, v := range menu.options {
    // fmt.Printf("Option %d: %s\n", i, v)
    menu.font.SetColor(1.0, 0.0, 0.0, 1.0) //r,g,b,a font color
	menu.font.Printf(100, 100*float32(i+1), 1.0, fmt.Sprintf("%s", v)) //x,y,scale,string,printf args
  }
}

func (menu *MenuState) Stop() bool {
  // Literally do nothing
  return true
}
