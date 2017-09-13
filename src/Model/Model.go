package Model

import (
	"os"
	"strconv"
	"fmt"
	"errors"
	"bufio"
	"strings"
	"reflect"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/gl/v4.1-core/gl"

	"Shader"
)

type Model struct {
	shader, vao, vbo uint32
	Faces, Normals []float32
	Max_radius float32
}

func (model *Model) Init(filename string, shader_program uint32) *Model {
	faces, norms, err := loadObjFile(filename)
	if err != nil {
		fmt.Printf("Failed to load Model: %v\n", err)
		os.Exit(1)
	}
	model.Faces = faces
	model.Normals = norms

	var buffer_data []float32
	buffer_data = append(buffer_data, model.Faces...)
	buffer_data = append(buffer_data, model.Normals...)
	model.Max_radius = 0.0
	for _, val := range faces {
		if val > model.Max_radius {
			model.Max_radius = val
		}
	}

	var temp_err error
	model.shader, temp_err = Shader.NewProgram("src/shaders/default.vert", "src/shaders/default.frag")
	if temp_err != nil {
		panic(temp_err)
	}
	gl.UseProgram(model.shader)

	gl.GenVertexArrays(1, &model.vao)
	gl.BindVertexArray(model.vao)

	gl.GenBuffers(1, &model.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, model.vbo)

	buffer_size := int( uintptr(len(buffer_data)) * reflect.TypeOf(buffer_data).Elem().Size() )
	gl.BufferData(gl.ARRAY_BUFFER, buffer_size, gl.Ptr(buffer_data), gl.STATIC_DRAW)

	vert_attrib := uint32(gl.GetAttribLocation(model.shader, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vert_attrib)
	gl.VertexAttribPointer(vert_attrib, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	norm_attrib := uint32(gl.GetAttribLocation(model.shader, gl.Str("norm\x00")))
	gl.EnableVertexAttribArray(norm_attrib)
	gl.VertexAttribPointer(norm_attrib, 3, gl.FLOAT, false, 0, gl.PtrOffset(buffer_size/2))

	return model
}

func (model *Model) Draw(model_uniform int32, entity_model mgl32.Mat4) {
	gl.UseProgram(model.shader)

	gl.BindVertexArray(model.vao)
	// gl.BindBuffer(gl.ARRAY_BUFFER, model.vbo)

	gl.UniformMatrix4fv(model_uniform, 1, false, &entity_model[0])
	gl.DrawArrays(gl.TRIANGLES, 0, int32( len(model.Faces)/3) )
}

func loadObjFile(file string) (face_floats []float32,
	norm_floats []float32,
	err error) {
	file_handle, file_err := os.Open(file)
	if file_err != nil {
		err = errors.New("Cannot open file: %s\n")
		return
	}

	scanner := bufio.NewScanner(file_handle)
	var face_verts, norm_verts []*mgl32.Vec3
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")

		if words[0] == "v" {
			var v *mgl32.Vec3
			x, x_err := strconv.ParseFloat(words[1], 32)
			y, y_err := strconv.ParseFloat(words[2], 32)
			z, z_err := strconv.ParseFloat(words[3], 32)
			if x_err != nil ||y_err != nil || z_err != nil {
				fmt.Printf("failed parsing !!!!!!!!!!!!!!!!!!!!!!!!!!!")
				os.Exit(1)
			}
			v = new(mgl32.Vec3)
			v[0] = float32(x)
			v[1] = float32(y)
			v[2] = float32(z)
			face_verts = append(face_verts, v)
		} else if words[0] == "vn" {
			var v *mgl32.Vec3
			x, x_err := strconv.ParseFloat(words[1], 32)
			y, y_err := strconv.ParseFloat(words[2], 32)
			z, z_err := strconv.ParseFloat(words[3], 32)
			if x_err != nil ||y_err != nil || z_err != nil {
				fmt.Printf("failed parsing !!!!!!!!!!!!!!!!!!!!!!!!!!!")
				os.Exit(1)
			}
			v = new(mgl32.Vec3)
			v[0] = float32(x)
			v[1] = float32(y)
			v[2] = float32(z)
			norm_verts = append(norm_verts, v)
		} else if words[0] == "f" {
			var f []float32
			var n []float32

			for _, val := range words[1:] {
				split_face_norm := strings.Split(val, "//")
				face, norm := split_face_norm[0], split_face_norm[1]

				face_index, face_e := strconv.ParseUint(face, 10, 32)
				norm_index, norm_e := strconv.ParseUint(norm, 10, 32)

				if face_e != nil || norm_e != nil {
					err_str := fmt.Sprintf("Error parsing int: %v\n", val)
					err = errors.New(err_str)
					return
				}

				face_vert := face_verts[face_index-1]
				norm_vert := norm_verts[norm_index-1]

				f = append(f, face_vert.X())
				f = append(f, face_vert.Y())
				f = append(f, face_vert.Z())

				n = append(n, norm_vert.X())
				n = append(n, norm_vert.Y())
				n = append(n, norm_vert.Z())
			}
			face_floats = append(face_floats, f...)
			norm_floats = append(norm_floats, n...)
		}
	}
	return
}
