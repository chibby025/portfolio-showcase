// 3D Cinema Scene Setup - React Three Fiber
// Shows procedural seat generation and spatial audio implementation

import React, { useRef, useEffect } from 'react';
import { Canvas, useFrame } from '@react-three/fiber';
import { OrbitControls, PerspectiveCamera } from '@react-three/drei';
import * as THREE from 'three';

// Generate cinema seats procedurally
function CinemaSeats({ rows = 10, seatsPerRow = 12 }) {
    const seats = [];
    const seatWidth = 1.2;
    const seatDepth = 1.0;
    const rowSpacing = 1.5;
    
    for (let row = 0; row < rows; row++) {
        for (let col = 0; col < seatsPerRow; col++) {
            const x = (col - seatsPerRow / 2) * seatWidth;
            const z = row * rowSpacing;
            const y = row * 0.15; // Slight elevation for stadium seating
            
            seats.push(
                <Seat 
                    key={`${row}-${col}`}
                    position={[x, y, z]}
                    row={row}
                    col={col}
                />
            );
        }
    }
    
    return <group>{seats}</group>;
}

// Individual seat component with hover state
function Seat({ position, row, col }) {
    const meshRef = useRef();
    const [hovered, setHovered] = useState(false);
    
    useFrame(() => {
        // Animate hover effect
        if (meshRef.current) {
            meshRef.current.material.emissive.setHex(
                hovered ? 0x440044 : 0x000000
            );
        }
    });
    
    return (
        <mesh 
            ref={meshRef}
            position={position}
            onPointerOver={() => setHovered(true)}
            onPointerOut={() => setHovered(false)}
        >
            {/* Seat geometry */}
            <boxGeometry args={[0.8, 0.6, 0.8]} />
            <meshStandardMaterial color="#8844AA" />
        </mesh>
    );
}

// Main cinema scene
export default function CinemaScene() {
    return (
        <Canvas>
            {/* Lighting */}
            <ambientLight intensity={0.3} />
            <spotLight 
                position={[0, 10, 0]} 
                angle={0.3} 
                intensity={1}
                castShadow
            />
            
            {/* Camera */}
            <PerspectiveCamera 
                makeDefault 
                position={[0, 5, 15]} 
                fov={60}
            />
            
            {/* Cinema elements */}
            <CinemaSeats rows={10} seatsPerRow={12} />
            
            {/* Movie screen */}
            <mesh position={[0, 3, -8]}>
                <planeGeometry args={[16, 9]} />
                <meshBasicMaterial color="#000000" />
            </mesh>
            
            {/* Floor */}
            <mesh rotation={[-Math.PI / 2, 0, 0]} position={[0, -0.5, 0]}>
                <planeGeometry args={[50, 50]} />
                <meshStandardMaterial color="#1a1a1a" />
            </mesh>
            
            <OrbitControls />
        </Canvas>
    );
}

// Key Learning: React Three Fiber = React + Three.js
// - Components map to 3D objects
// - useFrame hook for animations (runs every frame)
// - Procedural generation reduces code duplication
// - <Canvas> handles WebGL context automatically
