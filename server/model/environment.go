package model

import "fmt"

type Environment struct {
	particles     []*LivingParticle
	width, height int
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

func (e *Environment) Start() {
	fmt.Println("Starting environment")
	for idx, particle := range e.particles {
		fmt.Printf("Starting particle with id: %d\n", idx)
		go particle.Live()
	}
	fmt.Println("Done")
}
