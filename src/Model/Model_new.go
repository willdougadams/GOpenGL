package Model

import (
    "fmt"
    "os"
    "reflect"
    "github.com/go-gl/mathgl/mgl32"
    "github.com/go-gl/gl/v4.1-core/gl"
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

    gl.GenVertexArrays(1, &model.vao)
    gl.BindVertexArray(model.vao)

    gl.GenBuffers(1, &model.vbo)
    gl.BindBuffer(gl.ARRAY_BUFFER, model.vbo)

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
    gl.BindVertexArray(model.vao)
    gl.BindBuffer(gl.ARRAY_BUFFER, model.vbo)

    gl.UniformMatrix4fv(model_uniform, 1, false, &entity_model[0])
    gl.DrawArrays(gl.TRIANGLES, 0, int32( len(model.Faces)/3) )
}
