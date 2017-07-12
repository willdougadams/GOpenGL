package Model

import (
  "os"
  "strconv"
  "fmt"
  "errors"
  "bufio"
  "strings"

/*
  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/mathgl/mgl32"
  "reflect"
*/
)

type vec3 struct {
  X, Y, Z, W float32
}

func (v *vec3) Init(x, y, z, w string) *vec3 {
  x_val, xerr := strconv.ParseFloat(x, 32)
  y_val, yerr := strconv.ParseFloat(y, 32)
  z_val, zerr := strconv.ParseFloat(z, 32)
  w_val, werr := strconv.ParseFloat(w, 32)

  if xerr != nil {
    fmt.Printf("Failed to parse x float\n")
    os.Exit(1)
  }
  if yerr != nil {
    fmt.Printf("Failed to parse y float\n")
    os.Exit(1)
  }
  if zerr != nil {
    fmt.Printf("Failed to parse z float\n")
    os.Exit(1)
  }
  if werr != nil {
    fmt.Printf("Failed to parse w float\n")
    os.Exit(1)
  }

  v.X = float32(x_val)
  v.Y = float32(y_val)
  v.Z = float32(z_val)
  v.W = float32(w_val)

  return v
}

/*
type Model struct {
    shader uint32
  Faces []float32
  Normals []float32
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

    model.shader = shader_program

  buffer_size := int( uintptr(len(buffer_data)) * reflect.TypeOf(buffer_data).Elem().Size() )
  gl.BufferData(gl.ARRAY_BUFFER, buffer_size, gl.Ptr(buffer_data), gl.STATIC_DRAW)

  vert_attrib := uint32(gl.GetAttribLocation(shader_program, gl.Str("vert\x00")))
  gl.EnableVertexAttribArray(vert_attrib)
  gl.VertexAttribPointer(vert_attrib, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

  norm_attrib := uint32(gl.GetAttribLocation(shader_program, gl.Str("norm\x00")))
  gl.EnableVertexAttribArray(norm_attrib)
  gl.VertexAttribPointer(norm_attrib, 3, gl.FLOAT, false, 0, gl.PtrOffset(buffer_size/2))

  return model
}

func (model *Model) Draw(model_uniform int32, entity_model mgl32.Mat4) {
  gl.UniformMatrix4fv(model_uniform, 1, false, &entity_model[0])

  gl.DrawArrays(gl.TRIANGLES, 0, int32( len(model.Faces)/3) )
}
*/

func loadObjFile(file string) (face_floats []float32,
                                norm_floats []float32,
                                err error) {
  file_handle, file_err := os.Open(file)
  if file_err != nil {
    err = errors.New("Cannot open file: %s\n")
    return
  }

  scanner := bufio.NewScanner(file_handle)
  var face_verts, norm_verts []*vec3
  for scanner.Scan() {
    line := scanner.Text()
    words := strings.Split(line, " ")

    if words[0] == "v" {
      v := new(vec3).Init(words[1], words[2], words[3], "1")
      face_verts = append(face_verts, v)
    } else if words[0] == "vn" {
      v := new(vec3).Init(words[1], words[2], words[3], "1")
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

        f = append(f, face_vert.X)
        f = append(f, face_vert.Y)
        f = append(f, face_vert.Z)

        n = append(n, norm_vert.X)
        n = append(n, norm_vert.Y)
        n = append(n, norm_vert.Z)
      }
      face_floats = append(face_floats, f...)
      norm_floats = append(norm_floats, n...)
    }
  }

  return
}
