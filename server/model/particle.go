package model

import (
	"github.com/rtesselli/particles/server/common"
)

type Particle interface {
	Position() common.Point2D
	Size() int
	MaxAge() int
	Age() int
	Update()
}

type LivingParticle struct {
	Particle
	tick chan bool
	id   int
}

func NewLivingParticle(id int, particle Particle) LivingParticle {
	return LivingParticle{particle, make(chan bool), id}
}

func (l LivingParticle) Live(output_positions *common.SafeMap[int, common.ParticleData]) {
	for l.Age() < l.MaxAge() {
		output_positions.AddValue(l.id, common.NewParticleData(l.Position(), l.Size(), l.Age(), l.MaxAge()))
		l.Update()
		<-l.tick
	}
	output_positions.RemoveValue(l.id)
}
