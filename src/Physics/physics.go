package Physics

import (
	"Entity"
	"Landscape"

	// "github.com/go-gl/mathgl/mgl32"
)

const GRAV_ACCEL = -9.8
const TERM_VEL = 1.75 * GRAV_ACCEL
const TERRAIN_SIZE = 100

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

func gravity(ent *Entity.Entity, elapsed float32) {
	new_y_speed := ent.YSpeed() + GRAV_ACCEL*elapsed
	if new_y_speed < TERM_VEL {
		new_y_speed = TERM_VEL
	}
	ent.SetXSpeed(ent.XSpeed())
	ent.SetYSpeed(new_y_speed)
	ent.SetZSpeed(ent.XSpeed())
}

func inertia(ent *Entity.Entity, elapsed float32) {
	ent.SetLocation(ent.GetLocation().Add(ent.SpeedVec()))
	/*
	ent.x_orient += ent.x_rotate_speed * elapsed
	ent.y_orient += ent.y_rotate_speed * elapsed
	ent.z_orient += ent.z_rotate_speed * elapsed

	curr_xrspd := ent.x_rotate_speed
	if curr_xrspd != 0.0 {
		ent.x_rotate_speed = (abs32(curr_xrspd) / curr_xrspd) * max32((abs32(curr_xrspd)-(ent.drag*elapsed)), 0.0)
	} else {
		ent.x_rotate_speed = 0.0
	}

	curr_yrspd := ent.y_rotate_speed
	if curr_yrspd != 0.0 {
		ent.y_rotate_speed = (abs32(curr_yrspd) / curr_yrspd) * max32((abs32(curr_yrspd)-(ent.drag*elapsed)), 0.0)
	} else {
		ent.y_rotate_speed = 0.0
	}

	curr_zrspd := ent.z_rotate_speed
	if curr_zrspd != 0.0 {
		ent.z_rotate_speed = (abs32(curr_zrspd) / curr_zrspd) *
			max32((abs32(curr_zrspd)-(ent.drag*elapsed)), 0.0)
	} else {
		ent.z_rotate_speed = 0.0
	}

	trans_mat := mgl32.Translate3D(ent.location.X(), ent.location.Y(), ent.location.Z())

	x_rotate_mat := mgl32.HomogRotate3D(ent.x_orient, mgl32.Vec3{1, 0, 0})
	y_rotate_mat := mgl32.HomogRotate3D(ent.y_orient, mgl32.Vec3{0, 1, 0})
	z_rotate_mat := mgl32.HomogRotate3D(ent.z_orient, mgl32.Vec3{0, 0, -1})

	rotate_mat := x_rotate_mat.Mul4(y_rotate_mat).Mul4(z_rotate_mat)
	trans_rot_mat := trans_mat.Mul4(rotate_mat)

	scale_factor := ent.model_scale
	scale_mat := mgl32.Scale3D(scale_factor, scale_factor, scale_factor)
	ent.model_mat = scale_mat.Mul4(trans_rot_mat)
	*/
}

func Physics(land *Landscape.Landscape, ents []*Entity.Entity, elapsed float32) {
	for _, ent := range ents {
		inertia(ent, elapsed)

		heightmap_height := land.GetHeight(ent.X(), ent.Z())
		eye_level := heightmap_height + ent.Height()
		if ent.Y() <= eye_level {
			ent.SetY(eye_level)
			ent.SetYSpeed(0.0)
			ent.SetXSpeed(ent.XSpeed() * 0.5)
			ent.SetZSpeed(ent.ZSpeed() * 0.5)
		} else {
			gravity(ent, elapsed)
		}
	}
}
