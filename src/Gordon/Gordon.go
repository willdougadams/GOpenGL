package Gordon

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/glfw/v3.2/glfw"
	"fmt"
	"math"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Gordon struct {
	location, orientation mgl32.Vec3
	move_speed float32

	camera, projection mgl32.Mat4
	camera_uniform, projection_uniform int32
	shader uint32

	window *glfw.Window
	window_h, window_w int

	horizontal_angle, vertical_angle float64
	mouse_sensitivity float64

	gravity float64
}

func (gord *Gordon) Init(x, y, z float32,
						shader uint32,
						width int,
						height int,
						window *glfw.Window) *Gordon {
	gord.shader = shader
	gord.location = mgl32.Vec3{x, y, z}
	gord.orientation = mgl32.Vec3{0.0, 10.0, 0.0}
	gord.mouse_sensitivity = 0.001
	gord.move_speed = 10.0
	gord.window = window
	gord.window_w = width
	gord.window_h = height

	gord.projection = mgl32.Perspective(mgl32.DegToRad(65.0), float32(gord.window_w)/float32(gord.window_h), 0.1, 1000.0)
	gord.projection_uniform = gl.GetUniformLocation(shader, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(gord.projection_uniform, 1, false, &gord.projection[0])

	gord.camera = mgl32.LookAtV(gord.location,
		gord.orientation,
		mgl32.Vec3{0, 1, 0})

	gord.camera_uniform = gl.GetUniformLocation(gord.shader, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(int32(gord.camera_uniform), 1, false, &gord.camera[0])

	return gord
}

func (gord *Gordon) Update(elapsed float32) {
	if gord.window.GetKey(glfw.KeyEscape) == glfw.Press {
		fmt.Printf("\nESC pressed, exiting...\n")
		gord.window.SetShouldClose(true)
	}

	xpos, ypos := gord.window.GetCursorPos()
	gord.window.SetCursorPos(float64(gord.window_w/2.0), float64(gord.window_h/2.0))

	gord.horizontal_angle += gord.mouse_sensitivity * (float64(gord.window_w)/2.0 - xpos)
	gord.vertical_angle += gord.mouse_sensitivity * (float64(gord.window_h)/2.0 - ypos)

	gord.orientation = mgl32.Vec3{float32(math.Cos(gord.vertical_angle) * math.Sin(gord.horizontal_angle)), float32(math.Sin(gord.vertical_angle)), float32(math.Cos(gord.vertical_angle) * math.Cos(gord.horizontal_angle))}

	right := mgl32.Vec3{float32(math.Sin(gord.horizontal_angle - 3.14/2.0)), 0.0, float32(math.Cos(gord.horizontal_angle - 3.14/2.0))}
	up := right.Cross(gord.orientation)
	move_dist := elapsed * gord.move_speed


	if gord.window.GetKey(glfw.KeyW) == glfw.Press {
		gord.location = gord.location.Add(gord.orientation.Mul(move_dist))
	}
	if gord.window.GetKey(glfw.KeyS) == glfw.Press {
		gord.location = gord.location.Add(gord.orientation.Mul(-move_dist))
	}
	if gord.window.GetKey(glfw.KeyD) == glfw.Press {
		gord.location = gord.location.Add(right.Mul(move_dist))
	}
	if gord.window.GetKey(glfw.KeyA) == glfw.Press {
		gord.location = gord.location.Add(right.Mul(-move_dist))
	}

	gord.camera = mgl32.LookAtV(gord.location, gord.location.Add(gord.orientation), up)
	gl.UniformMatrix4fv(int32(gord.camera_uniform), 1, false, &gord.camera[0])
	gl.UniformMatrix4fv(gord.projection_uniform, 1, false, &gord.projection[0])
}
