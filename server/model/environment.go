package model

import (
	"fmt"

	"github.com/rtesselli/particles/server/common"
)

type Environment struct {
	particles     []*LivingParticle
	width, height int
}

func NewEnvironment(width, height int) Environment {
	return Environment{width: width, height: height}
}

func (e *Environment) AddParticle(particle Particle) {
	e.particles = append(e.particles, &LivingParticle{particle})
}

func (e *Environment) Width() int {
	return e.width
}

func (e *Environment) Height() int {
	return e.height
}

func (e *Environment) Start(output_positions *common.SafeMap[int, Particle]) {
	fmt.Println("Starting environment")
	for idx, particle := range e.particles {
		fmt.Printf("Starting particle with id: %d\n", idx)
		go particle.Live(output_positions)
	}
	fmt.Println("Done")
}
