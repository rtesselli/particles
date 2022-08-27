package model

type Particle interface {
	Position()
	Size()
	Velocity()
	Energy()
	Live()
}
