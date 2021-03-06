package Application

import (
	"States"
	"Debugs"

	_ "image/png"
	"log"
	"time"
	"runtime"
	"C"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const windowWidth = 1200
const windowHeight = 900

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

type Application struct {
	mr_manager *States.StateManager
	window *glfw.Window
}

func (app *Application) Init() *Application {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	Debugs.Print("Initializing GLFW...")
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	var err error
	app.window, err = glfw.CreateWindow(windowWidth, windowHeight, "GOpenGL", nil, nil)
	if err != nil {
		panic(err)
	}
	app.window.MakeContextCurrent()
	Debugs.Print("OS detected: " + runtime.GOOS + "\n")

	cursor_mode := glfw.CursorHidden    // default cursor mode
	if runtime.GOOS == "darwin" {
		cursor_mode = glfw.CursorDisabled // Mac lappyhas a touchpad
	}
	app.window.SetInputMode(glfw.CursorMode, cursor_mode)

	Debugs.Print("Initializing GL...")
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	Debugs.Print("OpenGL version " + version)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	Debugs.Print("Initializing Mr. Manager...")

	app.mr_manager = new(States.StateManager).Init(windowWidth, windowHeight, app.window)

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
