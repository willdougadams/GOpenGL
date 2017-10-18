#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 vert;
in vec2 UV;
in vec3 norm;

// out vec3 model_vert;
out vec4 world_vert;
out vec4 frag_norm;
out vec4 light_location;
out vec2 uv;

void main() {
  // model_vert = vert;
  uv = UV;
  frag_norm = projection * model * vec4(norm, 0.0);
  world_vert = projection * camera * model * vec4(vert, 1.0);
  light_location = vec4(0.0, 1000.5, 0.0, 1.0);

  gl_Position = world_vert;
}
