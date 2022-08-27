package model

import (
	"math/rand"
)

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

func (p StaticParticle) ID() int {
	return p.id
}

func NewStaticParticle(max_x, max_y, id, size, energy int) StaticParticle {
	initial_position := Point2D{rand.Intn(max_x), rand.Intn(max_y)}
	return StaticParticle{position: initial_position, id: id, size: size, energy: energy}
}
