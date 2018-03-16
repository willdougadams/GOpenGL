#version 330

uniform sampler2D tex;
uniform vec4 light_location;

in vec4 world_vert;     // vert position
in vec4 frag_norm;
in vec2 uv;

out vec4 outputColor;

void main() {
  vec4 light_color = vec4(1.0, 1.0, 1.0, 0.0);
  vec4 tex_color = texture(tex, uv);
  float diff = max(0.0, dot(normalize(light_location-world_vert), normalize(frag_norm)));

  outputColor = tex_color + diff * light_color;
}
