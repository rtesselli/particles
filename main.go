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
	environment := model.NewEnvironment(10, 10)
	environment.AddParticle(model.NewStaticParticle(10, 10, 0, 1, 400))
	environment.AddParticle(model.NewStaticParticle(10, 10, 1, 1, 200))
	environment.AddParticle(model.NewStaticParticle(10, 10, 2, 1, 100))
	environment.AddParticle(model.NewStaticParticle(10, 10, 3, 1, 300))
	environment.Start(&positions)
	for i := 0; i < 10; i++ {
		fmt.Printf("Peek %d", i)
		time.Sleep(500 * time.Millisecond)
		go func(p *common.SafeMap[int, model.Particle]) {
			fmt.Print(p.GetMap())
		}(&positions)
	}
	fmt.Println("Done main")
}
