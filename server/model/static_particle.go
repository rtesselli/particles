package model

type StaticParticle struct {
	position, velocity Point2D
	size               int
}

func (p StaticParticle) Position() Point2D {
	return p.position
}

func (p StaticParticle) Velocity() Point2D {
	return p.velocity
}

func (p StaticParticle) Size() int {
	return p.size
}
