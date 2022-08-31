package main

import (
	"math/rand"

	"github.com/rtesselli/particles/server/common"
	"github.com/rtesselli/particles/server/model"
	"github.com/rtesselli/particles/server/view"
)

func randInRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	positions := common.NewSafeMap[int, common.ParticleData]()
	const width = 1280
	const height = 960
	view := view.NewView(width, height, &positions)
	environment := model.NewEnvironment(width, height, &positions)
	for i := 0; i < 1000; i++ {
		environment.AddParticle(model.NewLinearParticle(width, height, 1, i*5, common.NewPoint2D(randInRange(-2, 2), randInRange(-2, 2))))
	}
	view.Run()
}
