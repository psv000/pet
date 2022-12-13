#version 410 core

out vec4 fragment;

uniform vec4 color;

void main()
{
    fragment.rgba = color.rgba;
}
