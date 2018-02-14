package Entity

import (
	"Model"

	"github.com/go-gl/mathgl/mgl32"
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

func abs32(a float32) float32 {
	var ab float32
	if a < 0 {
		ab = 0 - a
	} else {
		ab = a
	}
	return ab
}

type Entity struct {
	location, speed_vec mgl32.Vec3
	x_orient, y_orient, z_orient float32
	x_rotate_speed, y_rotate_speed, z_rotate_speed float32
	drag, rotational_drag float32
	model *Model.Model

	model_mat mgl32.Mat4
}

func (entity *Entity) X() float32 {return entity.location.X()}
func (entity *Entity) Y() float32 {return entity.location.Y()}
func (entity *Entity) Z() float32 {return entity.location.Z()}

func (entity *Entity) XSpeed() float32 {return entity.speed_vec.X()}
func (entity *Entity) YSpeed() float32 {return entity.speed_vec.Y()}
func (entity *Entity) ZSpeed() float32 {return entity.speed_vec.Z()}


func (entity *Entity) SetX(new_x float32) {entity.location = mgl32.Vec3{new_x, entity.Y(), entity.Z()}}
func (entity *Entity) SetY(new_y float32) {entity.location = mgl32.Vec3{entity.X(), new_y, entity.Z()}}
func (entity *Entity) SetZ(new_z float32) {entity.location = mgl32.Vec3{entity.X(), entity.Y(), new_z}}

func (entity *Entity) SetXSpeed(x_spd float32) {entity.speed_vec = mgl32.Vec3{x_spd, entity.speed_vec.Y(), entity.speed_vec.Z()}}
func (entity *Entity) SetYSpeed(y_spd float32) {entity.speed_vec = mgl32.Vec3{entity.speed_vec.X(), y_spd, entity.speed_vec.Z()}}
func (entity *Entity) SetZSpeed(z_spd float32) {entity.speed_vec = mgl32.Vec3{entity.speed_vec.X(), entity.speed_vec.Y(), z_spd}}

func (entity *Entity) Init(x float32,
	y float32,
	z float32,
	x_speed float32,
	y_speed float32,
	z_speed float32,
	shader uint32,
	model *Model.Model) *Entity {
	entity.location = mgl32.Vec3{x, y, z}
	entity.speed_vec = mgl32.Vec3{x_speed, y_speed, z_speed}

	entity.x_orient = 0.0
	entity.y_orient = 0.0
	entity.z_orient = 0.0
	entity.x_rotate_speed = 0.0 // (rand.Float32() * 2) - 1
	entity.y_rotate_speed = 0.0 // (rand.Float32() * 2) - 1
	entity.z_rotate_speed = 0.0 // (rand.Float32() * 2) - 1

	entity.drag = float32(0.05)
	entity.rotational_drag = float32(0.05)

	entity.model = model
	entity.model_mat = mgl32.Ident4()

	return entity
}

func (entity *Entity) Update(elapsed float32) {
	new_y_speed := entity.speed_vec.Y() + GRAV_ACCEL * elapsed
	if new_y_speed < TERM_VEL {
		new_y_speed = TERM_VEL
	}
	entity.speed_vec = mgl32.Vec3{entity.speed_vec.X(), new_y_speed, entity.speed_vec.Z()}

	entity.location = entity.location.Add(entity.speed_vec)

	entity.x_orient += entity.x_rotate_speed * elapsed
	entity.y_orient += entity.y_rotate_speed * elapsed
	entity.z_orient += entity.z_rotate_speed * elapsed

	curr_xrspd := entity.x_rotate_speed
	if curr_xrspd != 0.0 {
		entity.x_rotate_speed = (abs32(curr_xrspd)/curr_xrspd) *
		max32((abs32(curr_xrspd) - (entity.drag * elapsed)), 0.0)
	} else {
		entity.x_rotate_speed = 0.0
	}

	curr_yrspd := entity.y_rotate_speed
	if curr_yrspd != 0.0 {
		entity.y_rotate_speed = (abs32(curr_yrspd)/curr_yrspd) *
		max32((abs32(curr_yrspd) - (entity.drag * elapsed)), 0.0)
	} else {
		entity.y_rotate_speed = 0.0
	}

	curr_zrspd := entity.z_rotate_speed
	if curr_zrspd != 0.0 {
		entity.z_rotate_speed = (abs32(curr_zrspd)/curr_zrspd) *
		max32((abs32(curr_zrspd) - (entity.drag * elapsed)), 0.0)
	} else {
		entity.z_rotate_speed = 0.0
	}

	trans_mat := mgl32.Translate3D(entity.location.X(), entity.location.Y(), entity.location.Z())

	x_rotate_mat := mgl32.HomogRotate3D(entity.x_orient, mgl32.Vec3{1, 0, 0})
	y_rotate_mat := mgl32.HomogRotate3D(entity.y_orient, mgl32.Vec3{0, 1, 0})
	z_rotate_mat := mgl32.HomogRotate3D(entity.z_orient, mgl32.Vec3{0, 0, -1})

	rotate_mat := x_rotate_mat.Mul4(y_rotate_mat).Mul4(z_rotate_mat)
	trans_rot_mat := trans_mat.Mul4(rotate_mat)

	scale_factor := float32(10.0)
	scale_mat := mgl32.Scale3D(scale_factor, scale_factor, scale_factor)
	entity.model_mat = scale_mat.Mul4(trans_rot_mat)
}

func (entity * Entity) Draw(model_uniform int32) {
	entity.model.Draw(model_uniform, entity.model_mat)
}

func (entity *Entity) Stop() {
	// Literally do nothing
}
