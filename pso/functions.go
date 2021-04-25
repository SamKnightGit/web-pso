package pso

import "math"

func beale(pos []float64) float64 {
	return math.Pow(1.5-pos[0]+pos[0]*pos[1], 2) +
		math.Pow(2.25-pos[0]+pos[0]*math.Pow(pos[1], 2), 2) +
		math.Pow(2.625-pos[0]+pos[0]*math.Pow(pos[1], 3), 2)
}
