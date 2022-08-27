package model

import (
	"fmt"
	"time"
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

func (l LivingParticle) Live() {
	age := 0
	energy := l.Energy()
	for energy >= 0 {
		fmt.Printf("StaticParticle: ID %d, current age %d\n", l.ID(), age)
		energy--
		age++
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Printf("ID %d dies\n", l.ID())
}
