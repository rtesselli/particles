package model

import (
	"fmt"
	"time"

	"github.com/rtesselli/particles/server/common"
)

type Particle interface {
	Position() Point2D
	Velocity() Point2D
	Size() int
	MaxAge() int
	Age() int
	IncrementAge()
	ID() int
}

type LivingParticle struct {
	Particle
}

func (l LivingParticle) Live(output_positions *common.SafeMap[int, Particle]) {
	for l.Age() < l.MaxAge() {
		fmt.Printf("Particle: ID %d, current age %d\n", l.ID(), l.Age())
		output_positions.AddValue(l.ID(), l.Particle)
		l.IncrementAge()
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Printf("ID %d dies\n", l.ID())
	output_positions.RemoveValue(l.ID())
}
