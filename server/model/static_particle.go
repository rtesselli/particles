package model

import (
	"math/rand"

	"github.com/rtesselli/particles/server/common"
)

type StaticParticle struct {
	position           common.Point2D
	size, max_age, age int
}

func (p *StaticParticle) Position() common.Point2D {
	return p.position
}

func (p *StaticParticle) Size() int {
	return p.size
}

func (p *StaticParticle) MaxAge() int {
	return p.max_age
}

func (p *StaticParticle) Age() int {
	return p.age
}

func (p *StaticParticle) Update() {
	p.age++
}

func NewStaticParticle(max_x, max_y, size, max_age int) *StaticParticle {
	initial_position := common.NewPoint2D(rand.Intn(max_x), rand.Intn(max_y))
	return &StaticParticle{position: initial_position, size: size, max_age: max_age, age: 0}
}
