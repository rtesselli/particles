package main

import (
	"fmt"

	"github.com/rtesselli/particles/server/common"
	"github.com/rtesselli/particles/server/model"
)

func main() {
	fmt.Println("Start main")
	safe_map := common.NewSafeMap[int, model.Particle]()
	safe_map.AddValue(1, model.NewStaticParticle(10, 10, 10, 10, 10))
	safe_map.AddValue(2, model.NewStaticParticle(20, 20, 20, 20, 20))
	fmt.Println("Done main")
}
