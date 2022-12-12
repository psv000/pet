#version 410 core

out vec4 fragment;
in vec2 uv;

uniform vec4 color;

void main()
{
    float R = 1.0;
    float R2 = 0.7;
    float dist = sqrt(dot(uv,uv));
    if (dist >= R || dist <= R2) {
        discard;
    }
    float sm = smoothstep(R,R-0.01,dist);
    float sm2 = smoothstep(R2,R2+0.01,dist);
    float alpha = sm*sm2;
    fragment = vec4(color.rgb, alpha*color.a);
}