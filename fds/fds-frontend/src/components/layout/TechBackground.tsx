import React from 'react';
import { Box, styled, keyframes } from '@mui/material';

// Animation keyframes
const float = keyframes`
  0% { transform: translateY(0px); }
  50% { transform: translateY(-20px); }
  100% { transform: translateY(0px); }
`;

const pulse = keyframes`
  0% { opacity: 0.4; }
  50% { opacity: 0.8; }
  100% { opacity: 0.4; }
`;

const slideIn = keyframes`
  0% { transform: translateX(-100px) rotate(0deg); opacity: 0; }
  100% { transform: translateX(0px) rotate(360deg); opacity: 0.6; }
`;

const circuitFlow = keyframes`
  0% { stroke-dashoffset: 100; }
  100% { stroke-dashoffset: 0; }
`;

const hexagonSpin = keyframes`
  0% { transform: rotate(0deg) scale(1); }
  50% { transform: rotate(180deg) scale(1.1); }
  100% { transform: rotate(360deg) scale(1); }
`;

// Styled components
const BackgroundContainer = styled(Box)({
  position: 'fixed',
  top: 0,
  left: 0,
  width: '100vw',
  height: '100vh',
  overflow: 'hidden',
  zIndex: -1,
  pointerEvents: 'none',
});

const GeometricShape = styled(Box)<{ size?: number; delay?: number; x?: number; y?: number }>(
  ({ size = 60, delay = 0, x = 50, y = 50 }) => ({
    position: 'absolute',
    left: `${x}%`,
    top: `${y}%`,
    width: `${size}px`,
    height: `${size}px`,
    borderRadius: '50%',
    background: 'rgba(100, 116, 139, 0.2)',
    border: '1px solid rgba(100, 116, 139, 0.3)',
    animation: `${float} 6s ease-in-out infinite`,
    animationDelay: `${delay}s`,
  })
);

const HexagonShape = styled(Box)<{ size?: number; delay?: number; x?: number; y?: number }>(
  ({ size = 80, delay = 0, x = 30, y = 70 }) => ({
    position: 'absolute',
    left: `${x}%`,
    top: `${y}%`,
    width: `${size}px`,
    height: `${size}px`,
    background: 'rgba(34, 197, 94, 0.1)',
    border: '3px solid rgba(34, 197, 94, 0.4)',
    clipPath: 'polygon(50% 0%, 100% 25%, 100% 75%, 50% 100%, 0% 75%, 0% 25%)',
    animation: `${hexagonSpin} 8s linear infinite`,
    animationDelay: `${delay}s`,
  })
);

const TechGrid = styled(Box)({
  position: 'absolute',
  top: 0,
  left: 0,
  width: '100%',
  height: '100%',
  backgroundImage: `
    radial-gradient(circle at 1px 1px, rgba(100, 116, 139, 0.15) 1px, transparent 0)
  `,
  backgroundSize: '60px 60px',
  opacity: 0.6,
});

const CircuitPattern = styled('svg')({
  position: 'absolute',
  top: '10%',
  right: '5%',
  width: '300px',
  height: '200px',
  opacity: 0.3,
});

const ParticleContainer = styled(Box)({
  position: 'absolute',
  top: 0,
  left: 0,
  width: '100%',
  height: '100%',
});

const Particle = styled(Box)<{ delay?: number; x?: number; y?: number; duration?: number }>(
  ({ delay = 0, x = 0, y = 0, duration = 4 }) => ({
    position: 'absolute',
    left: `${x}%`,
    top: `${y}%`,
    width: '4px',
    height: '4px',
    borderRadius: '50%',
    background: 'rgba(59, 130, 246, 0.8)',
    boxShadow: '0 0 10px rgba(59, 130, 246, 0.5)',
    animation: `${pulse} ${duration}s ease-in-out infinite`,
    animationDelay: `${delay}s`,
  })
);

const GlowOrb = styled(Box)<{ size?: number; x?: number; y?: number; delay?: number }>(
  ({ size = 150, x = 20, y = 80, delay = 0 }) => ({
    position: 'absolute',
    left: `${x}%`,
    top: `${y}%`,
    width: `${size}px`,
    height: `${size}px`,
    borderRadius: '50%',
    background: 'radial-gradient(circle, rgba(59, 130, 246, 0.1), transparent 70%)',
    filter: 'blur(20px)',
    animation: `${pulse} 4s ease-in-out infinite`,
    animationDelay: `${delay}s`,
  })
);

export const TechBackground: React.FC = () => {
  // Generate random positions for particles
  const particles = Array.from({ length: 15 }, (_, i) => ({
    id: i,
    x: Math.random() * 100,
    y: Math.random() * 100,
    delay: Math.random() * 3,
    duration: 3 + Math.random() * 4,
  }));

  return (
    <BackgroundContainer>
      {/* Grid Pattern */}
      <TechGrid />
      
      {/* Geometric Shapes */}
      <GeometricShape size={80} delay={0} x={15} y={20} />
      <GeometricShape size={60} delay={1.5} x={85} y={15} />
      <GeometricShape size={100} delay={3} x={70} y={60} />
      
      {/* Hexagonal Shapes */}
      <HexagonShape size={60} delay={0.5} x={25} y={80} />
      <HexagonShape size={80} delay={2.5} x={90} y={45} />
      <HexagonShape size={70} delay={1} x={45} y={25} />
      
      {/* Glowing Orbs */}
      <GlowOrb size={200} x={80} y={70} delay={0} />
      <GlowOrb size={150} x={10} y={30} delay={2} />
      
      {/* Circuit Pattern */}
      <CircuitPattern viewBox="0 0 300 200">
        <defs>
          <linearGradient id="circuitGradient" x1="0%" y1="0%" x2="100%" y2="0%">
            <stop offset="0%" stopColor="rgba(100, 116, 139, 0.6)" />
            <stop offset="50%" stopColor="rgba(59, 130, 246, 0.6)" />
            <stop offset="100%" stopColor="rgba(34, 197, 94, 0.6)" />
          </linearGradient>
        </defs>
        
        {/* Circuit Lines */}
        <path
          d="M10,50 L50,50 L70,30 L120,30 L140,50 L200,50 L220,70 L280,70"
          stroke="url(#circuitGradient)"
          strokeWidth="2"
          fill="none"
          strokeDasharray="10,5"
          style={{
            animation: `${circuitFlow} 3s linear infinite`,
          }}
        />
        <path
          d="M10,100 L40,100 L60,120 L100,120 L120,100 L180,100 L200,80 L250,80"
          stroke="url(#circuitGradient)"
          strokeWidth="2"
          fill="none"
          strokeDasharray="8,4"
          style={{
            animation: `${circuitFlow} 4s linear infinite`,
            animationDelay: '1s',
          }}
        />
        
        {/* Circuit Nodes */}
        <circle cx="50" cy="50" r="4" fill="rgba(59, 130, 246, 0.8)" />
        <circle cx="120" cy="30" r="3" fill="rgba(34, 197, 94, 0.8)" />
        <circle cx="200" cy="50" r="4" fill="rgba(168, 85, 247, 0.8)" />
        <circle cx="100" cy="120" r="3" fill="rgba(239, 68, 68, 0.8)" />
        <circle cx="180" cy="100" r="4" fill="rgba(245, 158, 11, 0.8)" />
      </CircuitPattern>
      
      {/* Animated Particles */}
      <ParticleContainer>
        {particles.map((particle) => (
          <Particle
            key={particle.id}
            x={particle.x}
            y={particle.y}
            delay={particle.delay}
            duration={particle.duration}
          />
        ))}
      </ParticleContainer>
    </BackgroundContainer>
  );
}; 