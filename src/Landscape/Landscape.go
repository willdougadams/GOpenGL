package Landscape

import (
    "fmt"
    "Model"
    "Entity"
)

const GROUND_LEVEL = 0.0
const GRAV_ACCEL = -9.8
const TERM_VEL = 1.75 * GRAV_ACCEL
const TERRAIN_SIZE = 100


type Landscape struct {
    model *Model.Model
    entity *Entity.Entity
    heightmap []float32
}

func (land *Landscape) Init(shader uint32) *Landscape {
    fmt.Printf("Init Landscape...\n")
    land.model = new(Model.Model).Init("res/mountain/mount.obj", shader)
    land.entity = new(Entity.Entity).Init(0.0, 0.0, 0.0, 0.0, 0.0, 0.0, shader, land.model)

    var x_offset, z_offset int64
    for i, _ := range land.model.Faces {
        if i % 4 == 1 {
            x := int64(land.model.Faces[i-1])
            z := int64(land.model.Faces[i+1])

            if x < x_offset {
                x_offset = x
            }
            if z < z_offset {
                z_offset = z
            }
        }
    }

    if x_offset < 0 {
        x_offset = -x_offset
    }
    if z_offset < 0 {
        z_offset = -z_offset
    }

    land.heightmap = make([]float32, (len(land.model.Faces)/4+1) * (len(land.model.Faces)/4+1))

    for i, _ := range land.model.Faces {
        if i % 4 == 1 {
            x := int64(land.model.Faces[i-1])
            z := int64(land.model.Faces[i+1])

            o_x := x_offset + x
            o_z := z_offset + z
            land.heightmap[(o_x*o_z)+o_z] = land.model.Faces[i]
        }
    }

    return land
}

func (land *Landscape) Update() {
    // literally do nothing
}

func (land *Landscape) Draw(model_uniform int32) {
    land.entity.Draw(model_uniform)
}
