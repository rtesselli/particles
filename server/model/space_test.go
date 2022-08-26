package model

import (
	"testing"
)

func TestPoint2D(t *testing.T) {
	point := Point2D{x: 1, y: 2}
	if point.x != 1 && point.y != 2 {
		t.Errorf("Point init error")
	}
}
