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
	Particle
	id int
}

func NewLivingParticle(id int, particle Particle) LivingParticle {
	return LivingParticle{particle, id}
}

func (l LivingParticle) Live(output_positions *common.SafeMap[int, common.ParticleData]) {
	for l.Age() < l.MaxAge() {
		output_positions.AddValue(l.id, common.NewParticleData(l.Position(), l.Size(), l.Age(), l.MaxAge()))
		l.Update()
		time.Sleep(10 * time.Millisecond)
	}
	output_positions.RemoveValue(l.id)
}
