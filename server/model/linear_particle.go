package model

import (
	"math/rand"

	"github.com/rtesselli/particles/server/common"
)

type LinearParticle struct {
	position, velocity common.Point2D
	size, max_age, age int
}

func (p *LinearParticle) Position() common.Point2D {
	return p.position
}

func (p *LinearParticle) Size() int {
	return p.size
}

func (p *LinearParticle) MaxAge() int {
	return p.max_age
}

func (p *LinearParticle) Age() int {
	return p.age
}

func (p *LinearParticle) Update() {
	p.age++
	p.position = p.position.Add(p.velocity)
}

func NewLinearParticle(max_x, max_y, size, max_age int, velocity common.Point2D) *LinearParticle {
	initial_position := common.NewPoint2D(rand.Intn(max_x), rand.Intn(max_y))
	return &LinearParticle{position: initial_position, size: size, max_age: max_age, age: 0, velocity: velocity}
}
