package model

import (
	"github.com/rtesselli/particles/server/common"
)

type Environment struct {
	output_positions *common.SafeMap[int, common.ParticleData]
	particles        []LivingParticle
	width, height    int
}

func NewEnvironment(width, height int, output_positions *common.SafeMap[int, common.ParticleData]) Environment {
	return Environment{width: width, height: height, output_positions: output_positions}
}

func (e *Environment) AddParticle(particle Particle) {
	living_particle := NewLivingParticle(len(e.particles), particle)
	e.particles = append(e.particles, living_particle)
	go living_particle.Live(e.output_positions)
}

func (e *Environment) Width() int {
	return e.width
}

func (e *Environment) Height() int {
	return e.height
}

func (e *Environment) Tick() {
	for _, particle := range e.particles {
		particle.tick <- true
	}
}
