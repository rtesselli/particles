package model

import (
	"github.com/rtesselli/particles/server/common"
)

type Environment struct {
	output_positions *common.SafeMap[int, common.ParticleData]
	width, height    int
	particle_count   int
}

func NewEnvironment(width, height int, output_positions *common.SafeMap[int, common.ParticleData]) Environment {
	return Environment{width: width, height: height, output_positions: output_positions}
}

func (e *Environment) AddParticle(particle Particle) {
	go NewLivingParticle(e.particle_count, particle).Live(e.output_positions)
	e.particle_count++
}

func (e *Environment) Width() int {
	return e.width
}

func (e *Environment) Height() int {
	return e.height
}
