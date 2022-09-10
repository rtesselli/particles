package model

import (
	"testing"
	"time"

	"github.com/rtesselli/particles/server/common"
)

func TestLive(t *testing.T) {
	living := NewLivingParticle(0, NewStaticParticle(10, 10, 1, 10))
	positions := common.NewSafeMap[int, common.ParticleData]()
	go living.Live(&positions)
	for i := 0; i < 10; i++ {
		living.tick <- true
		time.Sleep(time.Millisecond)
	}
	if len(positions.GetMap()) != 0 {
		t.Errorf("Wrong size")
	}
}
