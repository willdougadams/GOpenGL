package Landscape

import (
	"fmt"
	"Model"

	"github.com/go-gl/mathgl/mgl32"
)

// const GROUND_LEVEL = 0.0
const GRAV_ACCEL = -9.8
const TERM_VEL = 1.75 * GRAV_ACCEL
const TERRAIN_SIZE = 100


type Landscape struct {
	model *Model.Model
	heightmap []float32
	scale_factor float32
	x_offset, z_offset int64
}

func (land *Landscape) Init(shader uint32) *Landscape {
	fmt.Printf("Init Landscape...\n")
	land.model = new(Model.Model).Init("res/mountain/mount.obj", shader)

	for i, _ := range land.model.Faces {
		if i % 4 == 1 {
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

	land.heightmap = make([]float32, (len(land.model.Faces)/4+1) * (len(land.model.Faces)/4+1))

	for i, _ := range land.model.Faces {
		if i % 4 == 1 {
			x := int64(land.model.Faces[i-1])
			z := int64(land.model.Faces[i+1])

			o_x := land.x_offset + x
			o_z := land.z_offset + z
			land.heightmap[(o_x*o_z)+o_z] = land.model.Faces[i]
		}
	}

	land.scale_factor = float32(1.0)

	return land
}

func (land *Landscape) Update() {
	// literally do nothing
}

func (land *Landscape) Draw(model_uniform int32) {
	scale_mat := mgl32.Scale3D(land.scale_factor, land.scale_factor, land.scale_factor)
	model_mat := mgl32.Ident4().Mul4(scale_mat)
	land.model.Draw(model_uniform, model_mat)
}

func (land *Landscape) GetHeight(x, z int) float32 {
	x_dist := x + int(land.x_offset)
	z_dist := z + int(land.z_offset)
	return land.heightmap[(x_dist * z_dist) + z_dist]
}
