package model

import (
	"time"

	"github.com/rtesselli/particles/server/common"
)

type Particle interface {
	Position() common.Point2D
	Velocity() common.Point2D
	Size() int
	MaxAge() int
	Age() int
	IncrementAge()
	ID() int
}

type LivingParticle struct {
	Particle
}

func (l LivingParticle) Live(output_positions *common.SafeMap[int, common.ParticleData]) {
	for l.Age() < l.MaxAge() {
		output_positions.AddValue(l.ID(), common.NewParticleData(l.Position(), l.Size(), l.Age(), l.MaxAge()))
		l.IncrementAge()
		time.Sleep(10 * time.Millisecond)
	}
	output_positions.RemoveValue(l.ID())
}
