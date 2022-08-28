package view

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rtesselli/particles/server/common"
	"github.com/rtesselli/particles/server/model"
)

type View struct {
	positions     *common.SafeMap[int, model.Particle]
	width, height int
}

func NewView(width, height int, positions *common.SafeMap[int, model.Particle]) *View {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Particles")
	return &View{width: width, height: height, positions: positions}
}

func (v *View) Update() error {
	return nil
}

func (v *View) Draw(screen *ebiten.Image) {

}

func (v *View) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return v.width, v.height
}

func (v *View) Run() {
	if err := ebiten.RunGame(v); err != nil {
		panic(err)
	}
}
