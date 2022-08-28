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
	Energy() int
	ID() int
}

type LivingParticle struct {
	Particle
}

func (l LivingParticle) Live(output_positions *common.SafeMap[int, Particle]) {
	age := 0
	energy := l.Energy()
	for energy >= 0 {
		fmt.Printf("Particle: ID %d, current age %d\n", l.ID(), age)
		output_positions.AddValue(l.ID(), l.Particle)
		energy--
		age++
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Printf("ID %d dies\n", l.ID())
}
