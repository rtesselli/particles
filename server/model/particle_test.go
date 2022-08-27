package model

import (
	"testing"
)

func TestLive(t *testing.T) {
	living := LivingParticle{NewStaticParticle(10, 10, 0, 1, 10)}
	living.Live()
}
