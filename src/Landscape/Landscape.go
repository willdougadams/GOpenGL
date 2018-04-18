package Landscape

import (
	//"Model"
	"Gordon"

	//"fmt"
	//"time"
	//"math/rand"

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

func generate_color(x, y, z float32) (float32, float32, float32) {
	red := (float32(simplexnoise.Noise2(float64(x), float64(z)))+1)*0.1
	green := (float32(simplexnoise.Noise2(float64(x), float64(z)))+1)*0.5
	blue := (float32(simplexnoise.Noise2(float64(x), float64(z)))+1)*0.1
	return red, green, blue
}

func (land *Landscape) Init(gord *Gordon.Gordon) *Landscape {
	land.gord  = gord
	land.width, land.distance = 1000, 1000
	vertex_data := []float32{}
	color_data := []float32{}

	for col := -land.width; col < land.width; col++ {
		for row := -land.distance; row < land.distance; row++ {
			c := float32(col)
			r := float32(row)

			x := c
			z := r+1
			y := land.GetHeight(x, z)
			vertex_data = append(vertex_data, x, y, z)
			red, g, b := generate_color(x, y, z)
			color_data = append(color_data, red, g, b)

			x = c
			z = r
			y = land.GetHeight(x, z)
			vertex_data = append(vertex_data, x, y, z)
			red, g, b = generate_color(x, y, z)
			color_data = append(color_data, red, g, b)

			x = c+1
			z = r
			y = land.GetHeight(x, z)
			vertex_data = append(vertex_data, x, y, z)
			red, g, b = generate_color(x, y, z)
			color_data = append(color_data, red, g, b)

			x = c+1
			z = r+1
			y = land.GetHeight(x, z)
			vertex_data = append(vertex_data, x, y, z)
			red, g, b = generate_color(x, y, z)
			color_data = append(color_data, red, g, b)

			x = c
			z = r+1
			y = land.GetHeight(x, z)
			vertex_data = append(vertex_data, x, y, z)
			red, g, b = generate_color(x, y, z)
			color_data = append(color_data, red, g, b)

			x = c+1
			z = r
			y = land.GetHeight(c+1, r)
			vertex_data = append(vertex_data, c+1.0, y, r)
			red, g, b = generate_color(x, y, z)
			color_data = append(color_data, red, g, b)
		}
	}

	var err error
	land.shader, err = load_terrain_shader()
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
	gl.BufferData(gl.ARRAY_BUFFER, len(vertex_data) * 3, gl.Ptr(vertex_data), gl.STATIC_DRAW)

	gl.GenBuffers(1, &land.color_buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, land.color_buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(color_data) * 3, gl.Ptr(color_data), gl.STATIC_DRAW)

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

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(land.verts)/3))

	gl.DisableVertexAttribArray(0)
	gl.DisableVertexAttribArray(1)
}

func (land *Landscape) GetHeight(x, z float32) float32 {
	a := simplexnoise.Noise2(float64(x/10), float64(z/10))
	b := simplexnoise.Noise2(float64(x/20), float64(z/20))
	c := simplexnoise.Noise2(float64(x/30), float64(z/30)) * 2
	d := simplexnoise.Noise2(float64(x/40), float64(z/40)) * 2
	e := simplexnoise.Noise2(float64(x/100), float64(z/100)) * 10
	return float32(a + b + c + d + e)
}
