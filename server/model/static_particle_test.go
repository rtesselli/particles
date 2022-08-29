package model

import (
	"testing"
)

func TestStaticParticle(t *testing.T) {
	particle := NewStaticParticle(10, 10, 0, 1, 10)
	if particle.position.X() < 0 || particle.position.X() >= 10 {
		t.Errorf("Wrong x")
	}
	if particle.position.Y() < 0 || particle.position.Y() >= 10 {
		t.Errorf("Wrong y")
	}
	if particle.id != 0 {
		t.Errorf("Wrong ID")
	}
	if particle.size != 1 {
		t.Errorf("Wrong size")
	}
	if particle.max_age != 10 {
		t.Errorf("Wrong max age")
	}
}
