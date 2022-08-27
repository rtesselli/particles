package model

import (
	"testing"
)

func TestStaticParticle(t *testing.T) {
	particle := NewStaticParticle(Point2D{0, 0}, 0, 1, 10)
	particle.Live()
}
