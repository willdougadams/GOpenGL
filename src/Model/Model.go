package Model

import (
	"os"
	"fmt"
	
	"Debugs"

	"github.com/go-gl/mathgl/mgl32"
)

type Model struct {
	shader, vao, vbo uint32
	Faces, UVs, Normals []float32
}

func (model *Model) Init(filename string, shader_program uint32) *Model {
	Debugs.Print(fmt.Sprintf("Initializing Model from %s", filename))
	faces, uvs, norms, err := loadObjFile(filename)
	if err != nil {
		fmt.Printf("Failed to load Model: %v\n", err)
		os.Exit(1)
	}
	model.Faces = faces
	model.UVs = uvs
	model.Normals = norms
	model.shader = shader_program

	model.vao = buffer(model)

	return model
}

func (model *Model) Draw(model_uniform int32, entity_model mgl32.Mat4) {
	draw(model, entity_model)
}
