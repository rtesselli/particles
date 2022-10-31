package main

import (
	"math/rand"

	"github.com/rtesselli/particles/server/common"
	"github.com/rtesselli/particles/server/controller"
	"github.com/rtesselli/particles/server/model"
	"github.com/rtesselli/particles/server/view"
)

const width = 1280
const height = 960

func randInRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	view := view.NewView(width, height)
	environment := model.NewEnvironment(width, height)
	controller := controller.NewController(environment, view)
	for i := 0; i < 1000; i++ {
		controller.AddParticle(model.NewLinearParticle(width, height, 1, i*5, common.NewPoint2D(randInRange(-2, 2), randInRange(-2, 2))))
	}
	controller.Run()
}
