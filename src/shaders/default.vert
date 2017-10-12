#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 vert;
in vec3 norm;

// out vec3 model_vert;
out vec4 world_vert;
out vec4 frag_norm;
out vec4 light_location;

void main() {
  // model_vert = vert;
  frag_norm = projection * camera * model * vec4(norm, 0.0);
  world_vert = projection * camera * model * vec4(vert, 1.0);
  light_location = vec4(0.5, 1000.5, 0.95, 1.0);

  gl_Position = world_vert;
}
