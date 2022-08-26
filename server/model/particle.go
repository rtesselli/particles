package model

type particle interface {
	position()
}

type StaticParticle struct {
	_position Point2D
}

func (p StaticParticle) position() Point2D {
	return p._position
}
