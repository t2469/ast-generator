import {StrictMode, useRef} from 'react';
import {Canvas, useFrame} from '@react-three/fiber';
import {createRoot} from 'react-dom/client'
import './shader.css'
import * as THREE from 'three';

const startTime = Date.now();

const Render = () => {
    const ref = useRef<THREE.Mesh>(null);

    useFrame(({clock}) => {
        if (ref.current && ref.current.material && (ref.current.material as THREE.ShaderMaterial).uniforms.uTime) {
            (ref.current.material as THREE.ShaderMaterial).uniforms.uTime.value =
                clock.getElapsedTime() - ((startTime / 1000) % 86000);
        }
    });
    return (
        <mesh ref={ref}>
            <planeGeometry args={[2, 2, 64, 64]}/>
            <shaderMaterial
                vertexShader={`
                    uniform float uTime;
                    varying vec2 vUv;
                    void main () {
                        vUv = vec2(uv.x+0.15,uv.y+0.2);
                        gl_Position = projectionMatrix * viewMatrix * modelMatrix * vec4(position, 1.0);
                    }
                    `}
                fragmentShader={`
                    precision highp float;
                    varying vec2 vUv;
                    uniform float uTime;
                    #define PI 3.1415
                    #define S(x,y,z) smoothstep(x,y,z)
                    #define B(x,y,z,b) S(x, x+b, z)*S(y+b, y, z)
                    #define sat(x) clamp(x,0.,1.)

                    #define MOD3 vec3(.1031,.11369,.13787)

                    #define MOONPOS vec2(1.3, .8)

                    //----------------------------------------------------------------------------------------
                    //  1 out, 1 in...
                    float hash11(float p) {
                        // From Dave Hoskins
                        vec3 p3  = fract(vec3(p) * MOD3);
                        p3 += dot(p3, p3.yzx + 19.19);
                        return fract((p3.x + p3.y) * p3.z);
                    }

                    //  1 out, 2 in...
                    float hash12(vec2 p) {
                        vec3 p3  = fract(vec3(p.xyx) * MOD3);
                        p3 += dot(p3, p3.yzx + 19.19);
                        return fract((p3.x + p3.y) * p3.z);
                    }

                    float remap(float a, float b, float c, float d, float t) {
                        return ((t-a) / (b-a)) * (d-c) + c;
                    }

                    float within(float a, float b, float t) {
                        return (t-a) / (b-a); 
                    }

                    float skewbox(vec2 uv, vec3 top, vec3 bottom, float blur) {
                        float y = within(top.z, bottom.z, uv.y);
                        float left = mix(top.x, bottom.x, y);
                        float right = mix(top.y, bottom.y, y);
                        
                        float horizontal = B(left, right, uv.x, blur);
                        float vertical = B(bottom.z, top.z, uv.y, blur);
                        return horizontal*vertical;
                    }

                    vec4 pine(vec2 uv, vec2 p, float s, float focus) {
                        uv.x -= .5;
                        float c = skewbox(uv, vec3(.0, .0, 1.), vec3(-.14, .14, .65), focus);
                        c += skewbox(uv, vec3(-.10, .10, .65), vec3(-.18, .18, .43), focus);
                        c += skewbox(uv, vec3(-.13, .13, .43), vec3(-.22, .22, .2), focus);
                        c += skewbox(uv, vec3(-.04, .04, .2), vec3(-.04, .04, -.1), focus);
                        
                        vec4 col = vec4(1.,1.,1.,0.);
                        col.a = c;
                    
                        float shadow = skewbox(uv.yx, vec3(.6, .65, .13), vec3(.65, .65, -.1), focus);
                        shadow += skewbox(uv.yx, vec3(.43, .43, .13), vec3(.36, .43, -.2), focus);
                        shadow += skewbox(uv.yx, vec3(.15, .2, .08), vec3(.17, .2, -.08), focus);
                        
                        col.rgb = mix(col.rgb, col.rgb*.8, shadow);
                        
                        return col;
                    }

                    float getheight(float x) {
                        return sin(x) + sin(x*2.234+.123)*.5 + sin(x*4.45+2.2345)*.25;
                    }

                    vec4 landscape(vec2 uv, float d, float p, float f, float a, float y, float seed, float focus) {
                        uv *= d;
                        float x = uv.x*PI*f+p;
                        float c = getheight(x)*a+y;
                        
                        float b = floor(x*5.)/5.+.1;
                        float h =  getheight(b)*a+y;
                        
                        float e = fwidth(uv.y);
                        
                        vec4 col = vec4(S(c+e, c-e, uv.y));
                        //col.rgb *= mix(0.9, 1., abs(uv.y-c)*20.);
                        
                        x *= 5.;
                        float id = floor(x);
                        float n = hash11(id+seed);
                        
                        x = fract(x);
                        
                        y = (uv.y - h)*mix(5., 3., n)*3.5;
                        float treeHeight = (.07/d) * mix(1.3, .5, n);
                        y = within(h, h+treeHeight, uv.y);
                        x += (n-.5)*.6;
                        vec4 pineCol = pine(vec2(x, y/d), vec2(0.), 1., focus+d*.1);
                        //col += pineCol;
                        col.rgb = mix(col.rgb, pineCol.rgb, pineCol.a);
                        col.a = max(col.a, pineCol.a);
                        
                        return sat(col);
                    }

                    vec4 gradient(vec2 uv) {
                        
                        float c = 1.-length(MOONPOS-uv)/1.4;
                        
                        vec4 col = vec4(c);
                        
                        return col;
                    }

                    float circ(vec2 uv, vec2 pos, float radius, float blur) {
                        float dist = length(uv-pos);
                        return S(radius+blur, radius-blur, dist);
                    }

                    vec4 moon(vec2 uv) {
                        float c = circ(uv, MOONPOS, .07, .001);
                        
                        c *= 1.-circ(uv, MOONPOS+vec2(.03), .07, .001)*.95;
                        c = sat(c);
                        
                        vec4 col = vec4(c);
                        col.rgb *=.8;
                        
                        return col;
                    }

                    vec4 moonglow(vec2 uv, float foreground) {
                        
                        float c = circ(uv, MOONPOS, .1, .2);
                        
                        vec4 col = vec4(c);
                        col.rgb *=.2;
                        
                        return col;
                    }

                    float stars(vec2 uv, float t) {
                        t*=3.;
                        
                        float n1 = hash12(uv*10000.);
                        float n2 = hash12(uv*11234.);
                        float alpha1 = pow(n1, 20.);
                        float alpha2 = pow(n2, 20.);
                        
                        float twinkle = sin((uv.x-t+cos(uv.y*20.+t))*10.);
                        twinkle *= cos((uv.y*.234-t*3.24+sin(uv.x*12.3+t*.243))*7.34);
                        twinkle = (twinkle + 1.)/2.;
                        return alpha1 * alpha2 * twinkle;
                    }

                    void main()
                    {
                        float t = uTime*.05;
                        
                        vec2 bgUV = vUv;
                        
                        vec4 col = gradient(bgUV)*.8;
                        col += moon(bgUV);
                        col += stars(vUv, t);
                        
                        float dist = .10;
                        float height = -.01;
                        float amplitude = .02;
                        
                        dist = 1.;
                        height = .55;
                        
                        vec4 trees = vec4(0.);
                        for(float i=0.; i<7.; i++) {    
                            vec4 layer = landscape(vUv, dist, t+i, 3., amplitude, height, i, .01);
                            layer.rgb *= mix(vec3(.1, .1, .2), vec3(.3)+gradient(vUv).x, 1.-i/10.);
                            trees = mix(trees, layer, layer.a);
                            
                            dist -= .1;
                            height -= .06;
                        }
                        col = mix(col, trees, trees.a);
                        
                        col += moonglow(bgUV, 1.);
                        col = sat(col);
                        
                        //vec4 foreground = landscape(vUv, .02, t, 3., .0, -0.04, 1., .1);
                        //foreground.rgb *= vec3(.1, .1, .2)*.5;
                        
                        //col = mix(col, foreground, foreground.a);
                        
                        gl_FragColor = vec4(col.rgb,1.);
                    }
                    `}
                uniforms={{uTime: {value: 10.}}}
            />
        </mesh>
    )
}

const Shader = () => {
    return (
        <Canvas
            camera={{
                fov: 10,
                near: 0.2,
                far: 1000,
                //position: [0, 0, calcDistFromFov(60)],
            }}
            id="canvas"
        >
            <Render/>
        </Canvas>
    )
}

export default Shader

createRoot(document.getElementById('shader')!).render(
    <StrictMode>
        <Shader/>
    </StrictMode>,
)
  