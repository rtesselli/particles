package main

import (
	"github.com/rtesselli/particles/server/common"
	"github.com/rtesselli/particles/server/model"
	"github.com/rtesselli/particles/server/view"
)

func main() {
	positions := common.NewSafeMap[int, common.ParticleData]()
	const width = 1280
	const height = 960
	view := view.NewView(width, height, &positions)
	environment := model.NewEnvironment(width, height, &positions)
	for i := 0; i < 1000; i++ {
		environment.AddParticle(model.NewStaticParticle(width, height, 1, i*10))
	}
	view.Run()
}
