package Model

import (
	"fmt"

	"Debugs"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/tbogdala/gombz"
	"github.com/tbogdala/assimp-go"
)

type Model struct {
	meshes []*gombz.Mesh
	shader, vao, vbo uint32
	Faces, UVs, Normals []float32
}

func (model *Model) Init(filename string, shader_program uint32) *Model {
	Debugs.Print(fmt.Sprintf("Initializing Model from %s", filename))

	meshes, mesh_err := assimp.ParseFile(filename)
	if mesh_err != nil {
		fmt.Printf(fmt.Sprintf("Failed to load assimp mesh: %s\n", mesh_err))
	} else {
		fmt.Printf(fmt.Sprintf("mesh:\n%v\n", meshes))
	}

	model.Faces = []float32{}
	model.UVs = []float32{}
	model.Normals = []float32{}
	for _, mesh := range meshes {
		uv_array := mesh.UVChannels[0]
		for _, f := range mesh.Faces {
			for _, j := range f {
				v := mesh.Vertices[j]
				model.Faces = append(model.Faces, v.X())
				model.Faces = append(model.Faces, v.Y())
				model.Faces = append(model.Faces, v.Z())

				if len(uv_array) > int(j) {
					u := uv_array[j]
					model.UVs = append(model.UVs, u.X())
					model.UVs = append(model.UVs, 1-u.Y())
				}

				if len(mesh.Normals) > 0 {
					t := mesh.Normals[j]
					model.Normals  = append(model.Normals, t.X())
					model.Normals  = append(model.Normals, t.Y())
					model.Normals  = append(model.Normals, t.Z())
				}
			}
		}
	}

	Debugs.Print(fmt.Sprintf("Loaded %d vert floats\n", len(model.Faces)))
	Debugs.Print(fmt.Sprintf("Loaded %d uv floats\n", len(model.UVs)))
	Debugs.Print(fmt.Sprintf("Loaded %d norm floats\n", len(model.Normals)))

	model.shader = shader_program
	model.vao = buffer(model)

	return model
}

func (model *Model) Draw(model_uniform int32, entity_model mgl32.Mat4) {
	draw(model, entity_model)
}
