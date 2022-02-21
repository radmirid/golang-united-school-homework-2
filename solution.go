package square

import (
	"math"
)

type length int

const (
	Pi            float64 = 3.14
	SidesCircle   length  = 0
	SidesTriangle length  = 3
	SidesSquare   length  = 4
)

func CalcSquare(sideLen float64, sidesNum length) float64 {
	switch sidesNum {
	case SidesCircle:
		return Pi * sideLen * sideLen
	case SidesTriangle:
		return (sideLen * sideLen * math.Sqrt(3)) / 4
	case SidesSquare:
		return sideLen * 4
	default:
		return 0
	}
}
