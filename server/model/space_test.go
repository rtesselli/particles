package model

import (
	"testing"

	"github.com/rtesselli/particles/server/common"
)

func TestPoint2D(t *testing.T) {
	point := common.NewPoint2D(1, 2)
	if point.X() != 1 && point.Y() != 2 {
		t.Errorf("Point init error")
	}
}
