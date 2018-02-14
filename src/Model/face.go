package Model

import (
	"strings"
	"strconv"
	"fmt"
	"errors"

	"github.com/go-gl/mathgl/mgl32"
)

func parse_face(words []string, face_verts, norm_verts []*mgl32.Vec3, tex_verts []*mgl32.Vec2) (f, t, n []float32, err error) {
	if strings.Contains(words[1], "//") { // just verts and norms
		for _, val := range words[1:] {
			split_face_norm := strings.Split(val, "//")
			face, norm := split_face_norm[0], split_face_norm[1]

			face_index, face_e := strconv.ParseUint(face, 10, 32)
			norm_index, norm_e := strconv.ParseUint(norm, 10, 32)

			if face_e != nil || norm_e != nil {
				err_str := fmt.Sprintf("Error parsing int: %v\n", val)
				err = errors.New(err_str)
				return
			}

			face_vert := face_verts[face_index-1]
			norm_vert := norm_verts[norm_index-1]

			f = append(f, face_vert.X())
			f = append(f, face_vert.Y())
			f = append(f, face_vert.Z())

			n = append(n, norm_vert.X())
			n = append(n, norm_vert.Y())
			n = append(n, norm_vert.Z())
		}
	} else if strings.Contains(words[1], "/") { // verts and textures, maybe norms
		if len(strings.Split(words[1], "/")) == 1 { // no norms
			for _, val := range words[1:] {
				split_face_norm := strings.Split(val, "/")
				face, tex := split_face_norm[0], split_face_norm[1]

				face_index, face_e := strconv.ParseUint(face, 10, 32)
				tex_index, tex_e := strconv.ParseUint(tex, 10, 32)

				if face_e != nil || tex_e != nil {
					err_str := fmt.Sprintf("Error parsing int: %v\n", val)
					err = errors.New(err_str)
					return
				}

				face_vert := face_verts[face_index-1]
				tex_vert := tex_verts[tex_index-1]

				f = append(f, face_vert.X())
				f = append(f, face_vert.Y())
				f = append(f, face_vert.Z())

				t = append(n, tex_vert.X())
				t = append(n, tex_vert.Y())
			}
		} else {			// with norms
			for _, val := range words[1:] {
				split_face_norm := strings.Split(val, "/")
				face, tex, norm := split_face_norm[0], split_face_norm[1], split_face_norm[2]

				face_index, face_e := strconv.ParseUint(face, 10, 32)
				tex_index, tex_e := strconv.ParseUint(tex, 10, 32)
				norm_index, norm_e := strconv.ParseUint(norm, 10, 32)

				if face_e != nil || norm_e != nil || tex_e != nil {
					err_str := fmt.Sprintf("Error parsing int: %v\n", val)
					err = errors.New(err_str)
					return
				}

				face_vert := face_verts[face_index-1]
				tex_vert := tex_verts[tex_index-1]
				norm_vert := norm_verts[norm_index-1]

				f = append(f, face_vert.X())
				f = append(f, face_vert.Y())
				f = append(f, face_vert.Z())

				t = append(t, tex_vert.X())
				t = append(t, tex_vert.Y())

				n = append(n, norm_vert.X())
				n = append(n, norm_vert.Y())
				n = append(n, norm_vert.Z())
			}
		}
	} else {  // just verts
		for _, val := range words[1:] {
			face := val
			face_index, face_e := strconv.ParseUint(face, 10, 32)

			if face_e != nil {
				err_str := fmt.Sprintf("Error parsing int: %v\n", val)
				err = errors.New(err_str)
				return
			}

			face_vert := face_verts[face_index-1]

			f = append(f, face_vert.X())
			f = append(f, face_vert.Y())
			f = append(f, face_vert.Z())
		}
	}

	return
}
