package main

import (
	"github.com/rtesselli/particles/server/common"
	"github.com/rtesselli/particles/server/view"
)

func main() {
	// fmt.Println("Start main")
	// positions := common.NewSafeMap[int, common.ParticleData]()
	// environment := model.NewEnvironment(10, 10, &positions)
	// environment.AddParticle(model.NewStaticParticle(10, 10, 0, 1, 400))
	// environment.AddParticle(model.NewStaticParticle(10, 10, 1, 1, 200))
	// environment.AddParticle(model.NewStaticParticle(10, 10, 2, 1, 100))
	// environment.AddParticle(model.NewStaticParticle(10, 10, 3, 1, 300))
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("Peek %d\n", i)
	// 	time.Sleep(500 * time.Millisecond)
	// 	go func(p *common.SafeMap[int, common.ParticleData]) {
	// 		for id, value := range p.GetMap() {
	// 			fmt.Printf("PEEK: ID: %d AGE: %d\n", id, value.Age)
	// 		}
	// 	}(&positions)
	// }
	// fmt.Println("Done main")

	positions := common.NewSafeMap[int, common.ParticleData]()
	view := view.NewView(320, 240, &positions)
	view.Run()
}
