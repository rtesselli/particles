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

func (c *Controller) Listen() {
	for {
		if v := <-c.view.Tick; v {
			c.Tick()
			c.view.Tick <- true
		}
	}
}

func (c *Controller) Run() {
	go c.Listen()
	c.view.Run()
}
