package Collision

import (
	"Entity"
	"math"
)

func dist3D(one, two *Entity.Entity) float64 {
	x_dist := float64(one.X() - two.X())
	y_dist := float64(one.Y() - two.Y())
	z_dist := float64(one.Z() - two.Z())

	return math.Sqrt(math.Pow(x_dist, 2) +
	math.Pow(y_dist, 2) +
	math.Pow(z_dist, 2))
}

func resolve(one, two *Entity.Entity) {
	diff_vec := one.
}

func ResolveEntities(entities []*Entity.Entity) {
	for index, ent := range entities {
		for i, other := range entities {
			if index == i {
				continue
			}
			if float32(dist3D(ent, other)) < ent.Collision_dist {
				// fmt.Printf("[!!] Collision detected: %d, %d\n", index, i)
				resolve(ent, other)
			}
		}
	}
}
