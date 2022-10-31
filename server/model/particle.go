package model

import (
	"sync"

	"github.com/rtesselli/particles/server/common"
)

type Particle interface {
	Position() common.Point2D
	Size() int
	MaxAge() int
	Age() int
	Update()
	Alive() bool
}

type LivingParticle struct {
	Particle
	id int
}

func NewLivingParticle(id int, particle Particle) LivingParticle {
	return LivingParticle{particle, id}
}

func (l LivingParticle) Tick(waitingGroup *sync.WaitGroup, outputPositions *common.SafeMap[int, common.ParticleData]) {
	defer waitingGroup.Done()
	if l.Alive() {
		outputPositions.AddValue(l.id, *common.NewParticleData(l.Position(), l.Size(), l.Age(), l.MaxAge()))
		l.Update()
		return
	}
}
