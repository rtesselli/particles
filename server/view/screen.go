package view

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rtesselli/particles/server/common"
	"github.com/rtesselli/particles/server/model"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type View struct {
	positions     *common.SafeMap[int, model.Particle]
	width, height int
	font          font.Face
}

func NewView(width, height int, positions *common.SafeMap[int, model.Particle]) *View {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Particles")
	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	font, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    10,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	return &View{width: width, height: height, positions: positions, font: font}
}

func (v *View) Update() error {
	return nil
}

func (v *View) Draw(screen *ebiten.Image) {
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS())
	text.Draw(screen, msg, v.font, 20, 40, color.White)
}

func (v *View) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return v.width, v.height
}

func (v *View) Run() {
	if err := ebiten.RunGame(v); err != nil {
		panic(err)
	}
}
