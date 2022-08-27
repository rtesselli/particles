package model

import (
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	environment := Environment{}
	environment.AddParticle(NewStaticParticle(10, 10, 1, 1, 10))
	environment.AddParticle(NewStaticParticle(10, 10, 2, 1, 20))
	environment.AddParticle(NewStaticParticle(10, 10, 3, 1, 30))
	environment.Start()
	time.Sleep(10 * time.Second) // TODO find better way to wait for goros
}
