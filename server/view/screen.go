package view

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rtesselli/particles/server/common"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type View struct {
	positions     *common.SafeMap[int, common.ParticleData]
	width, height int
	font          font.Face
	redCircle     *ebiten.Image
}

func NewView(width, height int) *View {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Particles")
	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	font, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    10,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	return &View{width: width, height: height, font: font, redCircle: CircleImage(2, 1, 0, 0, 1)}
}

func (v *View) SetPositions(positions *common.SafeMap[int, common.ParticleData]) {
	v.positions = positions
}

func (v *View) Update() error {
	return nil
}

func (v *View) Draw(screen *ebiten.Image) {
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS())
	text.Draw(screen, msg, v.font, 20, 40, color.White)
	positions := v.positions.GetMap()
	for _, position := range positions {
		v.DrawCircle(screen, position.Position)
	}
}

func (v *View) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return v.width, v.height
}

func (v *View) DrawCircle(screen *ebiten.Image, position common.Point2D) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(position.X()), float64(position.Y()))
	screen.DrawImage(v.redCircle, op)
}

func (v *View) Run() {
	if err := ebiten.RunGame(v); err != nil {
		panic(err)
	}
}
