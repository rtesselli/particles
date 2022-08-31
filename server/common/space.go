package common

type Point2D struct {
	x, y int
}

func NewPoint2D(x, y int) Point2D {
	return Point2D{x: x, y: y}
}

func (p Point2D) X() int {
	return p.x
}

func (p Point2D) Y() int {
	return p.y
}

func (p Point2D) Add(other Point2D) Point2D {
	return NewPoint2D(p.x+other.x, p.y+other.y)
}
