package square

import (
	"math"
)

type length int

const (
	Pi    float64 = 3.14
	zero  length  = 0
	three length  = 3
	four  length  = 4
)

func CalcSquare(sideLen float64, sidesNum length) float64 {
	switch sidesNum {
	case zero:
		return Pi * sideLen * sideLen
	case three:
		return (sideLen * sideLen * math.Sqrt(3)) / 4
	case four:
		return sideLen * 4
	default:
		return 0
	}
}
