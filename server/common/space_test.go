package common

import (
	"testing"
)

func TestPoint2D(t *testing.T) {
	point := NewPoint2D(1, 2)
	if point.X() != 1 && point.Y() != 2 {
		t.Errorf("Point init error")
	}
}
