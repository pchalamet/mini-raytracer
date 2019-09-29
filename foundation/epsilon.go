package foundation

import "math"


const epsilon = 0.00001

// Equals compare 2 floats and tell if they are equals
func Equals(a, b float64) bool {
	return math.Abs(a - b) < epsilon
}
