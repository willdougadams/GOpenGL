package Model

import (
  "reflect"

  "github.com/go-gl/mathgl/mgl32"
  "github.com/go-gl/gl/v4.1-core/gl"
)

var model_uniform int32

func buffer(model *Model) (vao uint32) {
  gl.UseProgram(model.shader)

  model_uniform = gl.GetUniformLocation(model.shader, gl.Str("model\x00"))

  var buffer_data []float32
	buffer_data = append(buffer_data, model.Faces...)
	buffer_data = append(buffer_data, model.UVs...)
	buffer_data = append(buffer_data, model.Normals...)

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

  var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	buffer_size := int( uintptr(len(buffer_data)) * reflect.TypeOf(buffer_data).Elem().Size() )
	gl.BufferData(gl.ARRAY_BUFFER, buffer_size, gl.Ptr(buffer_data), gl.STATIC_DRAW)

	vert_attrib := uint32(gl.GetAttribLocation(model.shader, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vert_attrib)
	gl.VertexAttribPointer(vert_attrib, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	tex_offset := int( uintptr(len(model.Faces)) * reflect.TypeOf(model.Faces).Elem().Size() )
	tex_attrib := uint32(gl.GetAttribLocation(model.shader, gl.Str("tex\x00")))
	gl.EnableVertexAttribArray(tex_attrib)
	gl.VertexAttribPointer(tex_attrib, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))

	norm_offset := int( uintptr(len(model.UVs)) * reflect.TypeOf(model.UVs).Elem().Size() ) + tex_offset
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

  return
}

func draw(model *Model, entity_model mgl32.Mat4) {
  gl.UseProgram(model.shader)
	gl.BindVertexArray(model.vao)
	// gl.BindBuffer(gl.ARRAY_BUFFER, model.vbo)
	gl.UniformMatrix4fv(model_uniform, 1, false, &entity_model[0])
	gl.DrawArrays(gl.TRIANGLES, 0, int32( len(model.Faces)/3) )
}
