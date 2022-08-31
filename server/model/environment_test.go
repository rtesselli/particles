package model

import (
	"testing"
	"time"

	"github.com/rtesselli/particles/server/common"
)

func TestStart(t *testing.T) {
	positions := common.NewSafeMap[int, common.ParticleData]()
	environment := NewEnvironment(10, 10, &positions)
	environment.AddParticle(NewStaticParticle(10, 10, 1, 10))
	environment.AddParticle(NewStaticParticle(10, 10, 1, 20))
	environment.AddParticle(NewStaticParticle(10, 10, 1, 30))
	time.Sleep(2 * time.Second) // TODO find better way to wait for goros
	if len(positions.GetMap()) != 0 {
		t.Errorf("Wrong size")
	}
}
