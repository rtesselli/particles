package main

import (
	"fmt"
	"time"

	"github.com/rtesselli/particles/server/common"
	"github.com/rtesselli/particles/server/model"
)

func main() {
	fmt.Println("Start main")
	positions := common.NewSafeMap[int, model.Particle]()
	environment := model.NewEnvironment(10, 10, &positions)
	environment.AddParticle(model.NewStaticParticle(10, 10, 0, 1, 400))
	environment.AddParticle(model.NewStaticParticle(10, 10, 1, 1, 200))
	environment.AddParticle(model.NewStaticParticle(10, 10, 2, 1, 100))
	environment.AddParticle(model.NewStaticParticle(10, 10, 3, 1, 300))
	for i := 0; i < 10; i++ {
		fmt.Printf("Peek %d\n", i)
		time.Sleep(500 * time.Millisecond)
		go func(p *common.SafeMap[int, model.Particle]) {
			for _, value := range p.GetMap() {
				fmt.Printf("PEEK: ID: %d AGE: %d\n", value.ID(), value.Age())
			}
		}(&positions)
	}
	fmt.Println("Done main")

	// positions := common.NewSafeMap[int, model.Particle]()
	// view := view.NewView(320, 240, &positions)
	// if err := ebiten.RunGame(view); err != nil {
	// 	panic(err)
	// }
}
