// TOON

void main()
{
  vec3 light_dir = vec3(0.5, 0.5, 0.95);
  float intensity = max(0.0, dot(light_dir, fNormal));
  float tooniness = 0.75;

  vec4 toon_color;

  if (intensity > pow(0.90, tooniness)) {
    toon_color = vec4(vec3(1.0), 1.0);
  } else if (intensity > pow(0.7, tooniness)) {
    toon_color = vec4(vec3(0.6), 1.0);
  } else if (intensity > pow(0.5, tooniness)) {
    toon_color = vec4(vec3(0.4), 1.0);
  } else if (intensity > pow(0.3, tooniness)) {
    toon_color = vec4(vec3(0.2), 1.0);
  } else if (intensity > pow(0.1, tooniness)) {
    toon_color = vec4(vec3(0.2), 1.0);
  } else {
    toon_color = vec4(vec3(0.0), 1.0);
  }

  vec4 n_color = 1.0 * vec4(fNormal, 1.0);
  toon_color = toon_color * 1.0;
  // toon_color = vec4(1.0);
  gl_FragColor = n_color * toon_color;
}


// CEL
/*
precision highp float;
uniform float time;
uniform vec2 resolution;
varying vec3 fPosition;
varying vec3 fNormal;

void main()
{
  vec4 color = vec4(fNormal, 1.0);
  float normal_angle = dot(fNormal, vec3(0, 0, -1));
  if (normal_angle <= -0.45) {
  gl_FragColor = color;
  }
  else {
    gl_FragColor = vec4(0.25, 0.25, 0.25, 1.0) * color;
  }
}
*/
