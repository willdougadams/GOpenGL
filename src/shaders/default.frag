#version 330
uniform sampler2D tex;

in vec4 world_vert;     // vert position
in vec4 frag_norm;
in vec4 light_location;
in vec2 uv;

out vec4 outputColor;

void main() {
  vec4 light_color = vec4(0.88, 0.61, 0.21, 0.0);

  normalize(frag_norm);

  vec4 tex_color = texture(tex, uv);

  float diff = max(0.0, dot(normalize(light_location-world_vert), frag_norm));
  vec4 diffuse = (0.25 * diff * light_color) + (0.75 * tex_color);

  outputColor = diffuse;
}
