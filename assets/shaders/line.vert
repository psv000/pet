#version 410 core


layout (location = 0) in vec3 position;
layout (location = 1) in vec2 normal;

uniform float thickness;
uniform mat4 projection;

void main() {
    //push the point along its normal by half thickness
    vec2 p = position.xy + vec2(normal * thickness/2.0);
    gl_Position = projection * vec4(p, 0.0, 1.0);
}