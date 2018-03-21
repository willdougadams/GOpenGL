package Landscape

import (
	"Debugs"
	"Model"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/larspensjo/Go-simplex-noise/simplexnoise"
)

const GRAV_ACCEL = -9.8
const TERM_VEL = 1.75 * GRAV_ACCEL
const TERRAIN_SIZE = 100

type Landscape struct {
	model              *Model.Model
	heightmap          []float32
	scale_factor       float32
	x_offset, z_offset int64
}

func (land *Landscape) Init(shader uint32) *Landscape {
	Debugs.Print("Init Landscape...\n")
	land.model = new(Model.Model).Init("res/chêne/tree 1.obj", "res/chêne/textures/grass.png", shader, 1.0)

	for i, _ := range land.model.Faces[1 : len(land.model.Faces)-1] {
		if i%4 == 1 {
			x := int64(land.model.Faces[i-1])
			z := int64(land.model.Faces[i+1])

			if x < land.x_offset {
				land.x_offset = x
			}
			if z < land.z_offset {
				land.z_offset = z
			}
		}
	}

	if land.x_offset < 0 {
		land.x_offset = -land.x_offset
	}
	if land.z_offset < 0 {
		land.z_offset = -land.z_offset
	}

	land.scale_factor = float32(1.0)

	return land
}

func (land *Landscape) Update() {}

func (land *Landscape) Draw(model_uniform int32) {
	return
	scale_mat := mgl32.Scale3D(land.scale_factor, land.scale_factor, land.scale_factor)
	model_mat := mgl32.Ident4().Mul4(scale_mat)
	land.model.Draw(model_uniform, model_mat)
}

func (land *Landscape) GetHeight(x, z float32) float32 {
	return float32(simplexnoise.Noise2(float64(x), float64(z)))
}
