package model

import (
	"sync"

	"github.com/rtesselli/particles/server/common"
)

type Environment struct {
	outputPositions *common.SafeMap[int, common.ParticleData]
	particles       *common.SafeMap[int, LivingParticle]
	width, height   int
	waitingGroup    sync.WaitGroup
}

func NewEnvironment(width, height int) *Environment {
	particles := common.NewSafeMap[int, LivingParticle]()
	return &Environment{width: width, height: height, particles: &particles}
}

func (e *Environment) SetOutputPositions(outputPositions *common.SafeMap[int, common.ParticleData]) {
	e.outputPositions = outputPositions
}

func (e *Environment) AddParticle(particle Particle) {
	livingParticle := NewLivingParticle(e.particles.Size(), particle)
	e.particles.AddValue(livingParticle.id, livingParticle)
}

func (e *Environment) Width() int {
	return e.width
}

func (e *Environment) Height() int {
	return e.height
}

func (e *Environment) Tick() {
	for _, particle := range e.particles.GetMap() {
		e.waitingGroup.Add(1)
		go particle.Tick(&e.waitingGroup, e.outputPositions)
	}
	e.waitingGroup.Wait()
	toRemove := make([]int, 0, e.particles.Size())
	for _, particle := range e.particles.GetMap() {
		if !particle.Alive() {
			toRemove = append(toRemove, particle.id)
		}
	}
	for _, id := range toRemove {
		e.particles.RemoveValue(id)
	}
}
