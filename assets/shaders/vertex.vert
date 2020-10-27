#version 460 core

in vec3 vertex;
out vec3 color;

uniform float yOff;
uniform float blueValue;

void main() {
    gl_Position = vec4(vertex.x, vertex.y + yOff, vertex.z, 1);
    color = vec3(1, 0, blueValue);
}