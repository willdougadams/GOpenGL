package Application
/*
#include <stdlib.h>
*/

import (
	"States"
	"Shader"
	"Debug"

	//"fmt"
	//"os"
	_ "image/png"
	"log"
	"time"
	"runtime"
	"C"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
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
	program, vao, vbo uint32
	texture uint32
	textureUniform int32
	vertAttrib, normAttrib uint32
}

func (app *Application) Init() *Application {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	Debug.Print("Initializing GLFW...")
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

	Debug.Print("Initializing GL...")
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	Debug.Print("OpenGL version " + version)

	Debug.Print("Compiling and linking shaders...")
	// Configure the vertex and fragment shaders
	app.program, app.temp_err = Shader.NewProgram("src/shaders/default.vert", "src/shaders/default.frag")
	if app.temp_err != nil {
		panic(app.temp_err)
	}
	gl.UseProgram(app.program)

	Debug.Print("Accessing Uniforms...")
	model_uniform := gl.GetUniformLocation(app.program, gl.Str("model\x00"))
	/*
	model := mgl32.Ident4()
	gl.UniformMatrix4fv(app.modelUniform, 1, false, &model[0])
	*/

	app.textureUniform = gl.GetUniformLocation(app.program, gl.Str("tex\x00"))
	gl.Uniform1i(app.textureUniform, 0)

	gl.BindFragDataLocation(app.program, 0, gl.Str("outputColor\x00"))

	Debug.Print("Loading texture...")
	// Load the texture
	app.texture, app.temp_err = Shader.NewTexture("/home/will/code/gopengl/res/square.png")
	if app.temp_err != nil {
		log.Fatalln(app.temp_err)
	}

	Debug.Print("Getting VAO & VBO...")
	// Configure the vertex data
	gl.GenVertexArrays(1, &app.vao)
	gl.BindVertexArray(app.vao)

	gl.GenBuffers(1, &app.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, app.vbo)

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	Debug.Print("Initializing Mr. Manager...")
	app.mr_manager = new(States.StateManager).Init(windowWidth,
		windowHeight,
		app.program,
		app.vao,
		app.texture,
		model_uniform,
		app.window)

	return app
}

func (app *Application) Run() {
	var time_delta, start_time, end_time int64
	var elapsed_seconds float64

	defer glfw.Terminate()

	end_time = time.Now().UnixNano()

	for !app.window.ShouldClose() {
		start_time = time.Now().UnixNano()
		time_delta = start_time - end_time
		elapsed_seconds = float64(time_delta) / (float64(time.Second)/float64(time.Nanosecond))

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		app.mr_manager.Update(float32(elapsed_seconds))
		app.mr_manager.Draw()

		app.window.SwapBuffers()
		glfw.PollEvents()

		end_time = start_time
	}
}
