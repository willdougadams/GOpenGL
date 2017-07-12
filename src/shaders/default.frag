#version 330
uniform sampler2D tex;

in vec3 model_vert;     // vert position
in vec4 frag_norm;
in vec4 light_dir;

out vec4 outputColor;

void main() {
  vec4 light_color = vec4(1.0, 0.0, 1.0, 1.0);

  normalize(frag_norm);
  normalize(light_dir);

  vec4 tex_color = texture(tex, vec2(0.5, 0.5));

  float diff = max(0.0, dot(light_dir, frag_norm));
  vec4 diffuse = diff * vec4(tex_color);

  vec4 diffuse_color = diffuse;
  outputColor = diffuse_color;
}
