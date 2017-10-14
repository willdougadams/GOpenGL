package main

import (
	"strings"
	"fmt"
	"strconv"
)

func main() {
	just_vert := "f v1 v2 v3"
	vert_tex := "f v1/vt1 v2/vt2 v3/vt3"
	vert_tex_norm := "f v1/vt1/vn1 v2/vt2/vn2 v3/vt3/vn3"
	vert_norm := "f v1//vn1 v2//vn2 v3//vn3"

	fmt.Printf(just_vert + "\n")
	fmt.Printf(strconv.FormatBool(strings.Contains(just_vert, "/")) + "\n")
	fmt.Printf(strconv.FormatBool(strings.Contains(just_vert, "//")) + "\n")
	fmt.Printf(strings.Join(strings.Split(just_vert, "/"), "") + "\n")
	fmt.Printf(strings.Join(strings.Split(just_vert, "//"), "") + "\n\n\n")

	fmt.Printf(vert_tex + "\n")
	fmt.Printf(strconv.FormatBool(strings.Contains(vert_tex, "/")) + "\n")
	fmt.Printf(strconv.FormatBool(strings.Contains(vert_tex, "//")) + "\n")
	fmt.Printf(strings.Join(strings.Split(vert_tex, "/"), "") + "\n")
	fmt.Printf(strings.Join(strings.Split(vert_tex, "//"), "") + "\n\n\n")

	fmt.Printf(vert_tex_norm + "\n")
	fmt.Printf(strconv.FormatBool(strings.Contains(vert_tex_norm, "/")) + "\n")
	fmt.Printf(strconv.FormatBool(strings.Contains(vert_tex_norm, "//")) + "\n")
	fmt.Printf(strings.Join(strings.Split(vert_tex_norm, "/"), "") + "\n")
	fmt.Printf(strings.Join(strings.Split(vert_tex_norm, "//"), "") + "\n\n\n")

	fmt.Printf(vert_norm + "\n")
	fmt.Printf(strconv.FormatBool(strings.Contains(vert_norm, "/")) + "\n")
	fmt.Printf(strconv.FormatBool(strings.Contains(vert_norm, "//")) + "\n")
	fmt.Printf(strings.Join(strings.Split(vert_norm, "/"), "") + "\n")
	fmt.Printf(strings.Join(strings.Split(vert_norm, "//"), "") + "\n\n\n")
}
