#version 330 core

// Input vertex data, different for all executions of this shader.
//layout(location = 0) in vec3 vert;
//layout(location = 1) in vec3 color;

in vec3 vert;
in vec3 color;

// Output data ; will be interpolated for each fragment.
out vec3 fragmentColor;

void main(){

	// Output position of the vertex, in clip space : MVP * position
	gl_Position = vec4(vert ,1);

	// The color of each vertex will be interpolated
	// to produce the color of each fragment
	fragmentColor = color;
}
