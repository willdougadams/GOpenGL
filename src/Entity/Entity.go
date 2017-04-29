package Entity

import (
  "Model"
  "math/rand"
  "github.com/go-gl/mathgl/mgl32"
  // "fmt"
)

func max32(a, b float32) float32 {
  var m float32
  if a > b {
    m = a
  } else {
    m = b
  }
  return m
}

type Entity struct {
  x, y, z float32
  x_speed, y_speed, z_speed float32
  x_orient, y_orient, z_orient float32
  x_rotate_speed, y_rotate_speed, z_rotate_speed float32
  drag, rotational_drag float32
  model *Model.Model

  model_mat mgl32.Mat4
}

func (entity *Entity) Init(x float32,
                            y float32,
                            z float32,
                            x_speed float32,
                            y_speed float32,
                            z_speed float32,
                            shader uint32,
                            model *Model.Model) *Entity {
  entity.x = x
  entity.y = y
  entity.z = z
  entity.x_speed = x_speed
  entity.y_speed = y_speed
  entity.z_speed = z_speed

  entity.x_orient = x
  entity.y_orient = y
  entity.z_orient = z
  entity.x_rotate_speed = rand.Float32() - 0.5
  entity.y_rotate_speed = rand.Float32() - 0.5
  entity.z_rotate_speed = rand.Float32() - 0.5

  entity.drag = float32(0.005)
  entity.rotational_drag = float32(0.00005)

  entity.model = model

  entity.model_mat = mgl32.Ident4()

  return entity
}

func (entity *Entity) Update(elapsed float32) {
  entity.x += entity.x_speed * float32(elapsed)
  entity.y += entity.y_speed * float32(elapsed)
  entity.z += entity.z_speed * float32(elapsed)

  entity.x_orient += entity.x_rotate_speed * float32(elapsed)
  entity.y_orient += entity.y_rotate_speed * float32(elapsed)
  entity.z_orient += entity.z_rotate_speed * float32(elapsed)

  entity.x_speed = max32(entity.x_speed - entity.drag, 0.0)
  entity.y_speed = max32(entity.y_speed - entity.drag, 0.0)
  entity.z_speed = max32(entity.z_speed - entity.drag, 0.0)

  entity.x_rotate_speed = max32(entity.x_rotate_speed - entity.rotational_drag, 0.0)
  entity.y_rotate_speed = max32(entity.y_rotate_speed - entity.rotational_drag, 0.0)
  entity.z_rotate_speed = max32(entity.z_rotate_speed - entity.rotational_drag, 0.0)

  trans_mat := mgl32.Translate3D(entity.x, entity.y, entity.z)

  x_rotate_mat := mgl32.HomogRotate3D(entity.x_orient, mgl32.Vec3{1, 0, 0})
  y_rotate_mat := mgl32.HomogRotate3D(entity.y_orient, mgl32.Vec3{0, 1, 0})
  z_rotate_mat := mgl32.HomogRotate3D(entity.z_orient, mgl32.Vec3{0, 0, -1})

  rotate_mat := x_rotate_mat.Mul4(y_rotate_mat).Mul4(z_rotate_mat)

  entity.model_mat = trans_mat.Mul4(rotate_mat)
}

func (entity * Entity) Draw(model_uniform int32) {
  entity.model.Draw(model_uniform, entity.model_mat)
}

func (entity *Entity) Stop() {
  // Literally do nothing
}
