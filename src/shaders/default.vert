#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 vert;
in vec2 UV;
in vec3 norm;

out vec4 world_vert;
out vec4 frag_norm;
out vec2 uv;
out mat4 mvp;

void main() {
  uv = UV;
  mvp = projection * camera * model;
  frag_norm = projection * camera * model * vec4(norm, 0.0);
  world_vert = projection * camera * model * vec4(vert, 1.0);

  gl_Position = world_vert;
}
