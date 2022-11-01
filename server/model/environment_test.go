package model

import (
	"fmt"
	"testing"

	"github.com/rtesselli/particles/server/common"
)

func TestEnvironmentTick(t *testing.T) {
	positions := common.NewSafeMap[int, common.ParticleData]()
	environment := NewEnvironment(10, 10)
	environment.SetOutputPositions(&positions)
	environment.AddParticle(NewStaticParticle(environment.width, environment.height, 1, 10))
	environment.AddParticle(NewStaticParticle(environment.width, environment.height, 1, 20))
	environment.AddParticle(NewStaticParticle(environment.width, environment.height, 1, 30))
	for i := 0; i < 30; i++ {
		fmt.Printf("Tick number %d\n", i)
		environment.Tick()
	}
	if v := positions.Size(); v != 0 {
		t.Errorf("Wrong size %d", v)
	}
}
