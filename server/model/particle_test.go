package model

import (
	"testing"

	"github.com/rtesselli/particles/server/common"
)

func TestLive(t *testing.T) {
	living := NewLivingParticle(0, NewStaticParticle(10, 10, 1, 10))
	positions := common.NewSafeMap[int, common.ParticleData]()
	living.Live(&positions)
	if len(positions.GetMap()) != 0 {
		t.Errorf("Wrong size")
	}
}
