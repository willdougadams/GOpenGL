package Entity

import (
	"Model"

	"github.com/go-gl/mathgl/mgl32"
)

type Entity struct {
	location, speed_vec                            mgl32.Vec3
	x_orient, y_orient, z_orient                   float32
	x_rotate_speed, y_rotate_speed, z_rotate_speed float32
	drag, rotational_drag                          float32
	model                                          *Model.Model

	model_mat mgl32.Mat4
	model_scale float32
}

func (entity *Entity) X() float32 { return entity.location.X() }
func (entity *Entity) Y() float32 { return entity.location.Y() }
func (entity *Entity) Z() float32 { return entity.location.Z() }

func (entity *Entity) XSpeed() float32 { return entity.speed_vec.X() }
func (entity *Entity) YSpeed() float32 { return entity.speed_vec.Y() }
func (entity *Entity) ZSpeed() float32 { return entity.speed_vec.Z() }

func (entity *Entity) SetX(new_x float32) { entity.location = mgl32.Vec3{new_x, entity.Y(), entity.Z()} }
func (entity *Entity) SetY(new_y float32) { entity.location = mgl32.Vec3{entity.X(), new_y, entity.Z()} }
func (entity *Entity) SetZ(new_z float32) { entity.location = mgl32.Vec3{entity.X(), entity.Y(), new_z} }

func (entity *Entity) SetXSpeed(x_spd float32) {
	entity.speed_vec = mgl32.Vec3{x_spd, entity.speed_vec.Y(), entity.speed_vec.Z()}
}
func (entity *Entity) SetYSpeed(y_spd float32) {
	entity.speed_vec = mgl32.Vec3{entity.speed_vec.X(), y_spd, entity.speed_vec.Z()}
}
func (entity *Entity) SetZSpeed(z_spd float32) {
	entity.speed_vec = mgl32.Vec3{entity.speed_vec.X(), entity.speed_vec.Y(), z_spd}
}

func (entity *Entity) GetLocation() mgl32.Vec3 {
	return entity.location
}

func (entity *Entity) SetLocation(loc mgl32.Vec3) {
	entity.location = loc
}

func (entity *Entity) Init(x float32,
	y float32,
	z float32,
	x_speed float32,
	y_speed float32,
	z_speed float32,
	shader uint32,
	model *Model.Model,
	model_scale float32) *Entity {
	entity.location = mgl32.Vec3{x, y, z}
	entity.speed_vec = mgl32.Vec3{x_speed, y_speed, z_speed}

	entity.x_orient = 0.0
	entity.y_orient = 0.0
	entity.z_orient = 0.0
	entity.x_rotate_speed = 0.0
	entity.y_rotate_speed = 0.0
	entity.z_rotate_speed = 0.0

	entity.drag = float32(0.05)
	entity.rotational_drag = float32(0.05)

	entity.model = model
	entity.model_mat = mgl32.Ident4()
	entity.model_scale = model_scale

	return entity
}

func (entity *Entity) Update(elapsed float32) {
	//	moved to Entity.physics call in GameState.Update
}

func (entity *Entity) Draw(model_uniform int32) {
	if entity.model != nil {
		entity.model.Draw(model_uniform, entity.model_mat)
	}
}

func (entity *Entity) Stop() {
	// Literally do nothing
}
