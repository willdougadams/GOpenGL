package Landscape

import (
	"Model"

	"github.com/go-gl/mathgl/mgl32"
  "github.com/go-gl/gl/v4.1-core/gl"
	"github.com/larspensjo/Go-simplex-noise/simplexnoise"
)

const GRAV_ACCEL = -9.8
const TERM_VEL = 1.75 * GRAV_ACCEL
const TERRAIN_SIZE = 100

type Landscape struct {
	shader, vao, vert_buffer, color_buffer uint32
	matrix_id int32
	model_mat mgl32.Mat4
}

func (land *Landscape) Init() *Landscape {
	var err error
	land.shader, err = Model.NewProgram("src/shaders/terrain.vert", "src/shaders/terrain.frag")
	if err != nil {
		panic(err)
	}
	gl.UseProgram(land.shader)
	gl.GenVertexArrays(1, &land.vao)
	gl.BindVertexArray(land.vao)
	land.matrix_id = gl.GetUniformLocation(land.shader, gl.Str("MVP\x00"))
	land.model_mat = mgl32.Ident4()

	g_vertex_buffer_data := []float32{
		-1.0,-1.0,-1.0,
		-1.0,-1.0, 1.0,
		-1.0, 1.0, 1.0,
		 1.0, 1.0,-1.0,
		-1.0,-1.0,-1.0,
		-1.0, 1.0,-1.0,
		 1.0,-1.0, 1.0,
		-1.0,-1.0,-1.0,
		 1.0,-1.0,-1.0,
		 1.0, 1.0,-1.0,
		 1.0,-1.0,-1.0,
		-1.0,-1.0,-1.0,
		-1.0,-1.0,-1.0,
		-1.0, 1.0, 1.0,
		-1.0, 1.0,-1.0,
		 1.0,-1.0, 1.0,
		-1.0,-1.0, 1.0,
		-1.0,-1.0,-1.0,
		-1.0, 1.0, 1.0,
		-1.0,-1.0, 1.0,
		 1.0,-1.0, 1.0,
		 1.0, 1.0, 1.0,
		 1.0,-1.0,-1.0,
		 1.0, 1.0,-1.0,
		 1.0,-1.0,-1.0,
		 1.0, 1.0, 1.0,
		 1.0,-1.0, 1.0,
		 1.0, 1.0, 1.0,
		 1.0, 1.0,-1.0,
		-1.0, 1.0,-1.0,
		 1.0, 1.0, 1.0,
		-1.0, 1.0,-1.0,
		-1.0, 1.0, 1.0,
		 1.0, 1.0, 1.0,
		-1.0, 1.0, 1.0,
		 1.0,-1.0, 1.0,
	}

	g_color_buffer_data := []float32{
		0.583,  0.771,  0.014,
		0.609,  0.115,  0.436,
		0.327,  0.483,  0.844,
		0.822,  0.569,  0.201,
		0.435,  0.602,  0.223,
		0.310,  0.747,  0.185,
		0.597,  0.770,  0.761,
		0.559,  0.436,  0.730,
		0.359,  0.583,  0.152,
		0.483,  0.596,  0.789,
		0.559,  0.861,  0.639,
		0.195,  0.548,  0.859,
		0.014,  0.184,  0.576,
		0.771,  0.328,  0.970,
		0.406,  0.615,  0.116,
		0.676,  0.977,  0.133,
		0.971,  0.572,  0.833,
		0.140,  0.616,  0.489,
		0.997,  0.513,  0.064,
		0.945,  0.719,  0.592,
		0.543,  0.021,  0.978,
		0.279,  0.317,  0.505,
		0.167,  0.620,  0.077,
		0.347,  0.857,  0.137,
		0.055,  0.953,  0.042,
		0.714,  0.505,  0.345,
		0.783,  0.290,  0.734,
		0.722,  0.645,  0.174,
		0.302,  0.455,  0.848,
		0.225,  0.587,  0.040,
		0.517,  0.713,  0.338,
		0.053,  0.959,  0.120,
		0.393,  0.621,  0.362,
		0.673,  0.211,  0.457,
		0.820,  0.883,  0.371,
		0.982,  0.099,  0.879,
	}

	gl.GenBuffers(1, &land.vert_buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, land.vert_buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(g_vertex_buffer_data), gl.Ptr(g_vertex_buffer_data), gl.STATIC_DRAW)

	gl.GenBuffers(1, &land.color_buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, land.color_buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(g_color_buffer_data), gl.Ptr(g_color_buffer_data), gl.STATIC_DRAW)

	return land
}

func (land *Landscape) Update() {}

func (land *Landscape) Draw() {
	gl.UseProgram(land.shader)
	gl.BindVertexArray(land.vao)

	gl.UniformMatrix4fv(land.matrix_id, 1, false, &land.model_mat[0])

	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, land.vert_buffer)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, land.color_buffer)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 0, nil)

	gl.DrawArrays(gl.TRIANGLES, 0, 12*3)

	gl.DisableVertexAttribArray(0)
	gl.DisableVertexAttribArray(1)
}

func (land *Landscape) GetHeight(x, z float32) float32 {
	return float32(simplexnoise.Noise2(float64(x), float64(z)))
}
