package Model

import (
	"strings"
	"strconv"
	"fmt"
	"errors"
	"os"
	"bufio"

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
		if strings.Count(words[1], "/") == 1 { // no norms
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
				t = append(n, 1-tex_vert.Y())
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
				t = append(t, 1-tex_vert.Y())

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

func loadObjFile(file string) (face_floats, tex_floats, norm_floats []float32, err error) {
	file_handle, file_err := os.Open(file)
	if file_err != nil {
		err = errors.New(fmt.Sprintf("Cannot open file: %s\n", file))
		return
	}

	scanner := bufio.NewScanner(file_handle)
	var face_verts, norm_verts []*mgl32.Vec3
	var tex_verts []*mgl32.Vec2
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		words := strings.Fields(line)

		if words[0] == "v" {
			var v *mgl32.Vec3
			x, x_err := strconv.ParseFloat(words[1], 32)
			y, y_err := strconv.ParseFloat(words[2], 32)
			z, z_err := strconv.ParseFloat(words[3], 32)
			if x_err != nil ||y_err != nil || z_err != nil {
				fmt.Printf("[!!] failed parsing")
				os.Exit(1)
			}
			v = new(mgl32.Vec3)
			v[0] = float32(x)
			v[1] = float32(y)
			v[2] = float32(z)
			face_verts = append(face_verts, v)
		} else if words[0] == "vt" {
			var vec *mgl32.Vec2
			u, u_err := strconv.ParseFloat(words[1], 32)
			v, v_err := strconv.ParseFloat(words[2], 32)
			if u_err != nil ||v_err != nil {
				fmt.Printf("[!!] failed parsing")
				os.Exit(1)
			}
			vec = new(mgl32.Vec2)
			vec[0] = float32(u)
			vec[1] = float32(v)
			tex_verts = append(tex_verts, vec)
		} else if words[0] == "vn" {
			var v *mgl32.Vec3
			x, x_err := strconv.ParseFloat(words[1], 32)
			y, y_err := strconv.ParseFloat(words[2], 32)
			z, z_err := strconv.ParseFloat(words[3], 32)
			if x_err != nil ||y_err != nil || z_err != nil {
				fmt.Printf("[!!] failed parsing")
				os.Exit(1)
			}
			v = new(mgl32.Vec3)
			v[0] = float32(x)
			v[1] = float32(y)
			v[2] = float32(z)
			norm_verts = append(norm_verts, v)
		}
	}

	face_file_handle, face_file_err := os.Open(file)
	if face_file_err != nil {
		err = errors.New(fmt.Sprintf("Cannot open file: %s\n", file))
		return
	}
	face_scanner := bufio.NewScanner(face_file_handle)
	for face_scanner.Scan() {
		line := face_scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		words := strings.Fields(line)

		if words[0] == "f" {
			var f, t, n []float32
			var err error

			f, t, n, err = parse_face(words, face_verts, norm_verts, tex_verts)

			if err != nil {
				panic(err)
			}

			if len(f) > 0 {
				face_floats = append(face_floats, f...)
			}
			if len(t) > 0 {
				tex_floats = append(tex_floats, t...)
			}
			if len(n) > 0 {
				norm_floats = append(norm_floats, n...)
			}
		}
	}
	return
}
