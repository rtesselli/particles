package model

import (
	"testing"

	"github.com/rtesselli/particles/server/common"
)

func TestStart(t *testing.T) {
	positions := common.NewSafeMap[int, common.ParticleData]()
	environment := NewEnvironment(10, 10, &positions)
	environment.AddParticle(NewStaticParticle(10, 10, 1, 10))
	environment.AddParticle(NewStaticParticle(10, 10, 1, 20))
	environment.AddParticle(NewStaticParticle(10, 10, 1, 30))
	for i := 0; i < 30; i++ {
		environment.Tick()
	}
	if len(positions.GetMap()) != 0 {
		t.Errorf("Wrong size")
	}
}
