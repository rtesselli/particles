package model

import (
	"sync"
	"testing"

	"github.com/rtesselli/particles/server/common"
)

func TestParticleTick(t *testing.T) {
	living := NewLivingParticle(0, NewStaticParticle(10, 10, 1, 10))
	positions := common.NewSafeMap[int, common.ParticleData]()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go living.Tick(&wg, &positions)
	}
	wg.Wait()
	if len(positions.GetMap()) != 1 {
		t.Errorf("Wrong size")
	}
	if living.Alive() {
		t.Errorf("Not dead")
	}
}
