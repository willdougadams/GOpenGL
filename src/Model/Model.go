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
)

type Model struct {
	shader, vao, vbo uint32
	Faces, UVs, Normals []float32
	Max_radius float32
}

func (model *Model) Init(filename string, shader_program uint32) *Model {
	faces, uvs, norms, err := loadObjFile(filename)
	if err != nil {
		fmt.Printf("Failed to load Model: %v\n", err)
		os.Exit(1)
	}
	model.Faces = faces
	model.UVs = uvs
	model.Normals = norms

	model.Max_radius = 0.0
	for _, val := range faces {
		if val > model.Max_radius {
			model.Max_radius = val
		}
	}

	var buffer_data []float32
	buffer_data = append(buffer_data, model.Faces...)
	buffer_data = append(buffer_data, model.UVs...)
	buffer_data = append(buffer_data, model.Normals...)

	model.shader = shader_program
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

	tex_offset := int( uintptr(len(faces)) * reflect.TypeOf(faces).Elem().Size() )
	tex_attrib := uint32(gl.GetAttribLocation(model.shader, gl.Str("tex\x00")))
	gl.EnableVertexAttribArray(tex_attrib)
	gl.VertexAttribPointer(tex_attrib, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))

	norm_offset := int( uintptr(len(uvs)) * reflect.TypeOf(uvs).Elem().Size() ) + tex_offset
	norm_attrib := uint32(gl.GetAttribLocation(model.shader, gl.Str("norm\x00")))
	gl.EnableVertexAttribArray(norm_attrib)
	gl.VertexAttribPointer(norm_attrib, 3, gl.FLOAT, false, 0, gl.PtrOffset(norm_offset))

	/*
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	buffer_size := int( uintptr(len(faces)) * reflect.TypeOf(faces).Elem().Size() )
	gl.BufferData(gl.ARRAY_BUFFER, buffer_size, gl.Ptr(faces), gl.STATIC_DRAW)
	vert_attrib := uint32(gl.GetAttribLocation(model.shader, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vert_attrib)
	gl.VertexAttribPointer(vert_attrib, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	uvs_size := int( uintptr(len(uvs)) * reflect.TypeOf(uvs).Elem().Size() )
	gl.BufferData(gl.ARRAY_BUFFER, uvs_size, gl.Ptr(uvs), gl.STATIC_DRAW)
	tex_attrib := uint32(gl.GetAttribLocation(model.shader, gl.Str("UV\x00")))
	gl.EnableVertexAttribArray(tex_attrib)
	gl.VertexAttribPointer(tex_attrib, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	norms_size := int( uintptr(len(norms)) * reflect.TypeOf(norms).Elem().Size() )
	gl.BufferData(gl.ARRAY_BUFFER, norms_size, gl.Ptr(norms), gl.STATIC_DRAW)
	norm_attrib := uint32(gl.GetAttribLocation(model.shader, gl.Str("norm\x00")))
	gl.EnableVertexAttribArray(norm_attrib)
	gl.VertexAttribPointer(norm_attrib, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))
	*/

	return model
}

func (model *Model) Draw(model_uniform int32, entity_model mgl32.Mat4) {
	gl.UseProgram(model.shader)
	gl.BindVertexArray(model.vao)
	// gl.BindBuffer(gl.ARRAY_BUFFER, model.vbo)
	gl.UniformMatrix4fv(model_uniform, 1, false, &entity_model[0])
	gl.DrawArrays(gl.TRIANGLES, 0, int32( len(model.Faces)/3) )
}

func loadObjFile(file string) (face_floats, tex_floats, norm_floats []float32, err error) {
	file_handle, file_err := os.Open(file)
	if file_err != nil {
		err = errors.New("Cannot open file: %s\n")
		return
	}

	scanner := bufio.NewScanner(file_handle)
	var face_verts, norm_verts []*mgl32.Vec3
	var tex_verts []*mgl32.Vec2
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")

		if words[0] == "v" {
			var v *mgl32.Vec3
			x, x_err := strconv.ParseFloat(words[1], 32)
			y, y_err := strconv.ParseFloat(words[2], 32)
			z, z_err := strconv.ParseFloat(words[3], 32)
			if x_err != nil ||y_err != nil || z_err != nil {
				fmt.Printf("[!!] failed parsing")
				os.Exit(1)
			}
			v = new(mgl32.Vec3)
			v[0] = float32(x)
			v[1] = float32(y)
			v[2] = float32(z)
			face_verts = append(face_verts, v)
		} else if words[0] == "vt" {
			var vec *mgl32.Vec2
			u, u_err := strconv.ParseFloat(words[1], 32)
			v, v_err := strconv.ParseFloat(words[2], 32)
			if u_err != nil ||v_err != nil {
				fmt.Printf("[!!] failed parsing")
				os.Exit(1)
			}
			vec = new(mgl32.Vec2)
			vec[0] = float32(u)
			vec[1] = float32(v)
			tex_verts = append(tex_verts, vec)
		} else if words[0] == "vn" {
			var v *mgl32.Vec3
			x, x_err := strconv.ParseFloat(words[1], 32)
			y, y_err := strconv.ParseFloat(words[2], 32)
			z, z_err := strconv.ParseFloat(words[3], 32)
			if x_err != nil ||y_err != nil || z_err != nil {
				fmt.Printf("[!!] failed parsing")
				os.Exit(1)
			}
			v = new(mgl32.Vec3)
			v[0] = float32(x)
			v[1] = float32(y)
			v[2] = float32(z)
			norm_verts = append(norm_verts, v)
		} else if words[0] == "f" {
			var f, t, n []float32
			var err error

			f, t, n, err = parse_face(words, face_verts, norm_verts, tex_verts)

			/*
			for _, f1 := range f {
				fmt.Printf(fmt.Sprintf("f: %f\n", f1))
			}

			for _, t1 := range t {
				fmt.Printf(fmt.Sprintf("t: %f\n", t1))
			}

			for _, n1 := range n {
				fmt.Printf(fmt.Sprintf("n: %f\n", n1))
			}
			*/

			if err != nil {
				panic(err)
			}

			if len(f) > 0 {
				face_floats = append(face_floats, f...)
			}
			if len(t) > 0 {
				tex_floats = append(tex_floats, t...)
			}
			if len(n) > 0 {
				norm_floats = append(norm_floats, n...)
			}
		}
	}
	return
}
