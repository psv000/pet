#version 410

layout (location = 0) in vec3 position;
layout (location = 1) in vec3 normal;

uniform float thickness;
uniform vec2 resolution;
uniform mat4 model;
uniform mat4 view;
uniform mat4 project;

void main() {
    vec3 n = normal;
    n.xy /= resolution;
    vec3 delta = vec3(n * thickness);
    gl_Position = project * view * model * vec4(position + delta, 1);
}