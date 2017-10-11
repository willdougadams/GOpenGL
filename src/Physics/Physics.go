package Physics

import (
	// "fmt"
	"Landscape"
	"Entity"
	// "Gordon"
)

func GroundCollision(land *Landscape.Landscape, ents []*Entity.Entity) {
	for _, ent := range ents {
		heightmap_height := land.GetHeight(int(ent.X()), int(ent.Z()))
		if ent.Y() <  heightmap_height {
			ent.SetY(heightmap_height)
			ent.SetYSpeed(float32(0))
		}
	}
}
