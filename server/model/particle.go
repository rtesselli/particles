package model

import (
	"time"

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
	particle Particle
	id       int
}

func NewLivingParticle(id int, particle Particle) LivingParticle {
	return LivingParticle{particle: particle, id: id}
}

func (l LivingParticle) Live(output_positions *common.SafeMap[int, common.ParticleData]) {
	for l.particle.Age() < l.particle.MaxAge() {
		output_positions.AddValue(l.id, common.NewParticleData(l.particle.Position(), l.particle.Size(), l.particle.Age(), l.particle.MaxAge()))
		l.particle.Update()
		time.Sleep(10 * time.Millisecond)
	}
	output_positions.RemoveValue(l.id)
}
