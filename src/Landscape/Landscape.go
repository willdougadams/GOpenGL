package Landscape

import (
	"Debugs"
	"Model"

	"fmt"
	"reflect"

	// "github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/larspensjo/Go-simplex-noise/simplexnoise"
)

const GRAV_ACCEL = -9.8
const TERM_VEL = 1.75 * GRAV_ACCEL
const TERRAIN_SIZE = 100

type Landscape struct {
	vao, shader uint32
	scale_factor float32
	width, depth int
}

func (land *Landscape) Init() *Landscape {
	Debugs.Print("Init Landscape...\n")
	verts := []float32{}
	colors := []float32{}

	land.width = 100
	land.depth = 100

	for row := -land.depth; row < land.depth; row++ {
		for col := -land.width; col < land.width; col++ {
			fmt.Printf(fmt.Sprintf("%d, %d\n", row, col))
			r := float32(row)
			c := float32(col)
			verts = append(verts, r, 0.0, c)
			colors = append(colors, 1.0, 0.0, 0.0)

			verts = append(verts, r, 0.0, c+1.0)
			colors = append(colors, 0.0, 1.0, 0.0)

			verts = append(verts, r+1.0, 0.0, c)
			colors = append(colors, 0.0, 0.0, 1.0)

			verts = append(verts, r+1.0, 0.0, c+1.0)
			colors = append(colors, 1.0, 0.0, 1.0)
		}
	}

	var shader_err error
	land.shader, shader_err = Model.NewProgram("src/shaders/terrain.vert", "src/shaders/terrain.frag")
	if shader_err != nil {
		Debugs.Print("[!!] Failed to load lanscape shader, exiting...")
	}
	gl.UseProgram(land.shader)
	gl.BindFragDataLocation(land.shader, 0, gl.Str("color\x00"))

	gl.GenVertexArrays(1, &land.vao)
	gl.BindVertexArray(land.vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	buffer_size := int( uintptr(len(verts)) * reflect.TypeOf(verts).Elem().Size() )
	gl.BufferData(gl.ARRAY_BUFFER, buffer_size, gl.Ptr(verts), gl.STATIC_DRAW)
	vert_attrib := uint32(gl.GetAttribLocation(land.shader, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vert_attrib)
	gl.VertexAttribPointer(vert_attrib, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	buffer_size = int( uintptr(len(colors)) * reflect.TypeOf(colors).Elem().Size() )
	gl.BufferData(gl.ARRAY_BUFFER, buffer_size, gl.Ptr(colors), gl.STATIC_DRAW)
	color_attrib := uint32(gl.GetAttribLocation(land.shader, gl.Str("color\x00")))
	gl.EnableVertexAttribArray(color_attrib)
	gl.VertexAttribPointer(color_attrib, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	return land
}

func (land *Landscape) Update() {}

func (land *Landscape) Draw() {
	gl.UseProgram(land.shader)
	gl.BindVertexArray(land.vao)
	//gl.UniformMatrix4fv(model_uniform, 1, false, &entity_model[0])
	//gl.BindTexture(gl.TEXTURE_2D, model.texture)
	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, int32( land.width * land.depth * 12) )
}

func (land *Landscape) GetHeight(x, z float32) float32 {
	return float32(simplexnoise.Noise2(float64(x), float64(z)))
}
