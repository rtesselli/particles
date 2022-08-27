package model

import "fmt"

type StaticParticle struct {
	position         Point2D
	id, size, energy int
}

func (p StaticParticle) Position() Point2D {
	return p.position
}

func (p StaticParticle) Velocity() Point2D {
	return Point2D{0, 0}
}

func (p StaticParticle) Size() int {
	return p.size
}

func (p StaticParticle) Energy() int {
	return p.energy
}

func NewStaticParticle(initial_position Point2D, id, size, energy int) StaticParticle {
	return StaticParticle{position: initial_position, id: id, size: size, energy: energy}
}

func (p StaticParticle) Live() {
	age := 0
	for p.energy >= 0 {
		fmt.Printf("StaticParticle: ID %d, current age %d\n", p.id, age)
		p.energy--
		age++
	}
	fmt.Printf("ID %d dies\n", p.id)
}
