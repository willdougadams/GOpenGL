package Application

import (
  "States"

  "fmt"
	"image"
	"image/draw"
	_ "image/png"
	"log"
	"os"
	"runtime"
	"strings"
  "io/ioutil"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	// "github.com/go-gl/mathgl/mgl32"
)

const windowWidth = 800
const windowHeight = 600

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

type Application struct {
  mr_manager *States.StateManager

  temp_err error
  window *glfw.Window
  time, elapsed, previousTime float64
  angle float64
  program, vao, vbo uint32
  //model int32
  texture uint32
  textureUniform int32
  vertAttrib, normAttrib uint32
}

func (app *Application) Init() *Application {
  if err := glfw.Init(); err != nil {
    log.Fatalln("failed to initialize glfw:", err)
  }
  // defer glfw.Terminate()

  glfw.WindowHint(glfw.Resizable, glfw.False)
  glfw.WindowHint(glfw.ContextVersionMajor, 4)
  glfw.WindowHint(glfw.ContextVersionMinor, 1)
  glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
  glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
  app.window, app.temp_err = glfw.CreateWindow(windowWidth, windowHeight, "Cube", nil, nil)
  if app.temp_err != nil {
    panic(app.temp_err)
  }
  app.window.MakeContextCurrent()
  app.window.SetInputMode(glfw.CursorMode, glfw.CursorHidden)

  // Initialize Glow
  if err := gl.Init(); err != nil {
    panic(err)
  }

  version := gl.GoStr(gl.GetString(gl.VERSION))
  fmt.Println("OpenGL version", version)

  // Configure the vertex and fragment shaders
  app.program, app.temp_err = newProgram("src/shaders/default.vert", "src/shaders/default.frag")
  if app.temp_err != nil {
    panic(app.temp_err)
  }

  gl.UseProgram(app.program)

  model_uniform := gl.GetUniformLocation(app.program, gl.Str("model\x00"))
/*
  model := mgl32.Ident4()
  gl.UniformMatrix4fv(app.modelUniform, 1, false, &model[0])
*/

  app.textureUniform = gl.GetUniformLocation(app.program, gl.Str("tex\x00"))
  gl.Uniform1i(app.textureUniform, 0)

  gl.BindFragDataLocation(app.program, 0, gl.Str("outputColor\x00"))

  // Load the texture
  app.texture, app.temp_err = newTexture("/home/will/code/go/res/square.png")
  if app.temp_err != nil {
    log.Fatalln(app.temp_err)
  }

  // Configure the vertex data
  gl.GenVertexArrays(1, &app.vao)
  gl.BindVertexArray(app.vao)

  gl.GenBuffers(1, &app.vbo)
  gl.BindBuffer(gl.ARRAY_BUFFER, app.vbo)

  // Configure global settings
  gl.Enable(gl.DEPTH_TEST)
  gl.DepthFunc(gl.LESS)
  gl.ClearColor(1.0, 1.0, 1.0, 1.0)

  app.mr_manager = new(States.StateManager).Init(windowWidth,
                                                  windowHeight,
                                                  app.program,
                                                  app.vao,
                                                  app.texture,
                                                  model_uniform,
                                                  app.window)


  app.angle = 0.0
  app.previousTime = glfw.GetTime()

  return app
}

func (app *Application) Run() {
  for !app.window.ShouldClose() {
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

    // Update
    app.mr_manager.Update(app.elapsed)

    // Render
    app.mr_manager.Draw()

    // Maintenance
    app.window.SwapBuffers()
    glfw.PollEvents()
  }
}

func key_callback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
  if action == glfw.Press {
    switch key {
    case glfw.KeyEscape:
      window.SetShouldClose(true)
    }
  }
}

//////////////////////////
///// OTHER STUFF
//////////////////////////

func newProgram(vertexShaderFilename, fragmentShaderFilename string) (uint32, error) {
  vertexShaderBytes, v_err := ioutil.ReadFile(vertexShaderFilename)
  if v_err != nil {
    panic(v_err)
  }
  vertexShaderSource := string(vertexShaderBytes)
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

  fragmentShaderBytes, f_err := ioutil.ReadFile(fragmentShaderFilename)
  if f_err != nil {
    panic(f_err)
  }
  fragmentShaderSource := string(fragmentShaderBytes)

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func newTexture(file string) (uint32, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return 0, fmt.Errorf("texture %q not found on disk: %v", file, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return 0, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return 0, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture, nil
}
