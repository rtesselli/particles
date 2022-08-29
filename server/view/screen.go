package view

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rtesselli/particles/server/common"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/fogleman/gg"
)

type View struct {
	positions     *common.SafeMap[int, common.ParticleData]
	width, height int
	font          font.Face
}

func NewView(width, height int, positions *common.SafeMap[int, common.ParticleData]) *View {
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
	img := ebiten.NewImageFromImage(Circle())
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(img, op)
}

func (v *View) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return v.width, v.height
}

func (v *View) Run() {
	if err := ebiten.RunGame(v); err != nil {
		panic(err)
	}
}

func Circle() image.Image {
	dc := gg.NewContext(100, 100)
	dc.DrawCircle(50, 50, 25)
	dc.SetRGB(1, 0, 0)
	dc.Fill()
	return dc.Image()
}
