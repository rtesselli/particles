package model

type Particle interface {
	position()
}

type StaticParticle struct {
	position Point2D
}

func (p StaticParticle) Position() Point2D {
	return p.position
}
