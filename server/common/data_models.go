package common

type ParticleData struct {
	Position Point2D
	Size     int
	Age      int
	ToLive   int
}

func NewParticleData(position Point2D, size, age, max_age int) ParticleData {
	return ParticleData{Position: position, Size: size, Age: age, ToLive: max_age - age}
}
