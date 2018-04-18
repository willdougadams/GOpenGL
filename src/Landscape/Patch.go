package Landscape

import (
  "io/ioutil"
  "math"
  "fmt"
  "strings"
  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/mathgl/mgl32"
)

type Patch struct {
  location mgl32.Vec2
  world_pos, local_scaling, local_translation mgl32.Vec3
  lod, gap float32
  vbo uint32

  is_leaf bool
  verts [16]mgl32.Vec2
}

func (patch *Patch) Init(location mgl32.Vec2, lod float32, index mgl32.Vec2, shader uint32)  *Patch {
  //patch.shader = load_terrain_shader()
  gl.GenBuffers(1, &patch.vbo)
  gl.BindBuffer(gl.ARRAY_BUFFER, patch.vbo)

  patch.location = location
  patch.lod = lod
  patch.is_leaf = true
  patch.gap = float32(1.0 / (8 * math.Pow(float64(patch.lod), float64(2))) ) // 1f/(TerrainQuadtree.getRootNodes() * (float)(Math.pow(2, lod)));

  patch.local_scaling = mgl32.Vec3{patch.gap, 0.0, patch.gap}
  patch.local_translation = mgl32.Vec3{patch.location.X(), 0.0, patch.location.Y()}

  i := 0

	patch.verts[i] = mgl32.Vec2{0,0}
  i += 1
	patch.verts[i] = mgl32.Vec2{0.333,0}
  i += 1
	patch.verts[i] = mgl32.Vec2{0.666,0}
  i += 1
	patch.verts[i] = mgl32.Vec2{1,0}
  i += 1

	patch.verts[i] = mgl32.Vec2{0,0.333}
  i += 1
	patch.verts[i] = mgl32.Vec2{0.333,0.333}
  i += 1
	patch.verts[i] = mgl32.Vec2{0.666,0.333}
  i += 1
	patch.verts[i] = mgl32.Vec2{1,0.333}
  i += 1

	patch.verts[i] = mgl32.Vec2{0,0.666}
  i += 1
	patch.verts[i] = mgl32.Vec2{0.333,0.666}
  i += 1
	patch.verts[i] = mgl32.Vec2{0.666,0.666}
  i += 1
	patch.verts[i] = mgl32.Vec2{1, 0.666}
  i += 1

	patch.verts[i] = mgl32.Vec2{0,1}
  i += 1
	patch.verts[i] = mgl32.Vec2{0.333,1}
  i += 1
	patch.verts[i] = mgl32.Vec2{0.666,1}
  i += 1
	patch.verts[i] = mgl32.Vec2{1,1}

  buffer_size := len(patch.verts) // int( uintptr(len(model.Faces)) * reflect.TypeOf(model.Faces).Elem().Size() )
	gl.BufferData(gl.ARRAY_BUFFER, buffer_size, gl.Ptr(patch.verts), gl.STATIC_DRAW)
	vert_attrib := uint32(gl.GetAttribLocation(shader, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vert_attrib)
	gl.VertexAttribPointer(vert_attrib, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))

  return patch
}

func (patch *Patch) Update(elapsed float32) {
  // literally do nothing
}

func (patch *Patch) render() {
  // literally do nothing
}

func load_terrain_shader() (uint32, error) {
  // vert
	vertexShaderBytes, v_err := ioutil.ReadFile("src/shaders/terrain/Terrain_VS.glsl")
	if v_err != nil {
		panic(v_err)
	}
	vertexShaderBytes = append(vertexShaderBytes, '\x00')

	vertexShader, err := compileShader(string(vertexShaderBytes), gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

  // Tess control
  tessControlShaderBytes, v_err := ioutil.ReadFile("src/shaders/terrain/Terrain_TC.glsl")
  if v_err != nil {
    panic(v_err)
  }
  tessControlShaderBytes = append(tessControlShaderBytes, '\x00')

  tessControlShader, err := compileShader(string(tessControlShaderBytes), gl.TESS_CONTROL_SHADER)
  if err != nil {
    return 0, err
  }

  // Tess eval
  tessEvalShaderBytes, v_err := ioutil.ReadFile("src/shaders/terrain/Terrain_TE.glsl")
  if v_err != nil {
    panic(v_err)
  }
  tessEvalShaderBytes = append(tessEvalShaderBytes, '\x00')

  tessEvalShader, err := compileShader(string(tessEvalShaderBytes), gl.TESS_EVALUATION_SHADER)
  if err != nil {
    return 0, err
  }

  //Geometry
  GeometryShaderBytes, v_err := ioutil.ReadFile("src/shaders/terrain/Terrain_GS.glsl")
  if v_err != nil {
    panic(v_err)
  }
  GeometryShaderBytes = append(GeometryShaderBytes, '\x00')

  geometryShader, err := compileShader(string(GeometryShaderBytes), gl.GEOMETRY_SHADER)
  if err != nil {
    return 0, err
  }

  // frag
	fragmentShaderBytes, f_err := ioutil.ReadFile("src/shaders/terrain/Terrain_FS.glsl")
	if f_err != nil {
		panic(f_err)
	}

	fragmentShaderBytes = append(fragmentShaderBytes, '\x00')
	fragmentShaderSource := string(fragmentShaderBytes)

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

  shader := gl.CreateProgram()
	gl.AttachShader(shader, vertexShader)
  gl.AttachShader(shader, tessControlShader)
  gl.AttachShader(shader, tessEvalShader)
  gl.AttachShader(shader, geometryShader)
	gl.AttachShader(shader, fragmentShader)
	gl.LinkProgram(shader)

	var status int32
	gl.GetProgramiv(shader, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
  gl.DeleteShader(tessControlShader)
  gl.DeleteShader(tessEvalShader)
  gl.DeleteShader(geometryShader)
	gl.DeleteShader(fragmentShader)

	return shader, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
