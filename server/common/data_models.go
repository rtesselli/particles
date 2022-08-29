package common

type ParticleData struct {
	Position Point2D
	Size     int
	Age      int
}

func NewParticleData(position Point2D, size, age int) ParticleData {
	return ParticleData{Position: position, Size: size, Age: age}
}
