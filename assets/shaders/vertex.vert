#version 460 core

in vec3 vertex;
out vec4 color;

void main() {
    gl_Position = vec4(vertex, 1);
    color = vec4(1, 0, 0, 1);
}