package model

import (
	"testing"
	"time"

	"github.com/rtesselli/particles/server/common"
)

func TestStart(t *testing.T) {
	environment := NewEnvironment(10, 10)
	environment.AddParticle(NewStaticParticle(10, 10, 1, 1, 10))
	environment.AddParticle(NewStaticParticle(10, 10, 2, 1, 20))
	environment.AddParticle(NewStaticParticle(10, 10, 3, 1, 30))
	positions := common.NewSafeMap[int, Particle]()
	environment.Start(&positions)
	time.Sleep(2 * time.Second) // TODO find better way to wait for goros
	if len(positions.GetMap()) != 3 {
		t.Errorf("Wrong size")
	}
}
