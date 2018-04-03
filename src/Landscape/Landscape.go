package Landscape

import (
	"Model"
	"Gordon"

	"time"
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
  "github.com/go-gl/gl/v4.1-core/gl"
	"github.com/larspensjo/Go-simplex-noise/simplexnoise"
)

const GRAV_ACCEL = -9.8
const TERM_VEL = 1.75 * GRAV_ACCEL
const TERRAIN_SIZE = 100

type Landscape struct {
	width, distance int
	shader, vao, vert_buffer, color_buffer uint32
	verts, colors []float32
	matrix_id int32
	mvp mgl32.Mat4
	counter float32
	gord *Gordon.Gordon
}

func (land *Landscape) Init(gord *Gordon.Gordon) *Landscape {
	land.gord  = gord
	land.width, land.distance = 100, 100
	vertex_data := []float32{}
	color_data := []float32{}

	for col := -land.width; col < land.width; col++ {
		for row := -land.distance; row < land.distance; row++ {
			c := float32(col)
			r := float32(row)
			y := land.GetHeight(c, r)
			vertex_data = append(vertex_data, c,     y, r)
			y = land.GetHeight(c+1, r)
			vertex_data = append(vertex_data, c+1.0, y, r)
			y = land.GetHeight(c, r+1)
			vertex_data = append(vertex_data, c,     y, r+1.0)
			y = land.GetHeight(c+1, r+1)
			vertex_data = append(vertex_data, c+1.0, y, r+1.0)

			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)

			r, g, b := r1.Float32(), r1.Float32(), r1.Float32()
			color_data = append(color_data, r, g, b)
			r, g, b = r1.Float32(), r1.Float32(), r1.Float32()
			color_data = append(color_data, r, g, b)
			r, g, b = r1.Float32(), r1.Float32(), r1.Float32()
			color_data = append(color_data, r, g, b)
			r, g, b = r1.Float32(), r1.Float32(), r1.Float32()
			color_data = append(color_data, r, g, b)
		}
	}

	var err error
	land.shader, err = Model.NewProgram("src/shaders/terrain.vert", "src/shaders/terrain.frag")
	if err != nil {
		panic(err)
	}
	gl.UseProgram(land.shader)
	gl.GenVertexArrays(1, &land.vao)
	gl.BindVertexArray(land.vao)
	land.matrix_id = gl.GetUniformLocation(land.shader, gl.Str("MVP\x00"))

	land.mvp = mgl32.Ident4()

	gl.GenBuffers(1, &land.vert_buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, land.vert_buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertex_data) * 4, gl.Ptr(vertex_data), gl.STATIC_DRAW)

	gl.GenBuffers(1, &land.color_buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, land.color_buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(color_data) * 4, gl.Ptr(color_data), gl.STATIC_DRAW)

	land.verts = vertex_data
	land.colors = color_data

	return land
}

func (land *Landscape) Update(elapsed float32) {
	land.mvp = land.gord.GetViewProjection()
}

func (land *Landscape) Draw() {
	gl.UseProgram(land.shader)
	gl.BindVertexArray(land.vao)

	gl.UniformMatrix4fv(land.matrix_id, 1, false, &land.mvp[0])

	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, land.vert_buffer)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, land.color_buffer)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 0, nil)

	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, int32(len(land.verts)))

	gl.DisableVertexAttribArray(0)
	gl.DisableVertexAttribArray(1)
}

func (land *Landscape) GetHeight(x, z float32) float32 {
	return float32(simplexnoise.Noise2(float64(x), float64(z)))
}
