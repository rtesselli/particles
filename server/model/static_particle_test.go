package model

import (
	"testing"
)

func TestStaticParticle(t *testing.T) {
	particle := NewStaticParticle(10, 10, 0, 1, 10)
	if particle.position.x < 0 || particle.position.x >= 10 {
		t.Errorf("Wrong x")
	}
	if particle.position.y < 0 || particle.position.y >= 10 {
		t.Errorf("Wrong y")
	}
	if particle.id != 0 {
		t.Errorf("Wrong ID")
	}
	if particle.size != 1 {
		t.Errorf("Wrong size")
	}
	if particle.energy != 10 {
		t.Errorf("Wrong energy")
	}
}
