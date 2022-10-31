package controller

import (
	"github.com/rtesselli/particles/server/common"
	"github.com/rtesselli/particles/server/model"
	"github.com/rtesselli/particles/server/view"
)

type Controller struct {
	model *model.Environment
	view  *view.View
}

func NewController(model *model.Environment, view *view.View) Controller {
	positions := common.NewSafeMap[int, common.ParticleData]()
	model.SetOutputPositions(&positions)
	view.SetPositions(&positions)
	return Controller{model, view}
}

func (c *Controller) Tick() {
	c.model.Tick()
}

func (c *Controller) AddParticle(particle model.Particle) {
	c.model.AddParticle(particle)
}

func (c *Controller) Run() {
	c.view.Run()
}
