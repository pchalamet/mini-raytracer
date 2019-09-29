package foundation

import "math"

// Vector structure
type Vector struct {
	X float64
	Y float64
	Z float64
	W float64
}


// NewPoint creates a new point
func NewPoint(x, y, z float64) Vector {
	return Vector { x, y, z, 1}
}

// NewVector creates a new vector
func NewVector(x, y, z float64) Vector {
	return Vector { x, y, z, 0}
}

// Add adds 2 vectors
func (v1 Vector) Add(v2 Vector) Vector {
	return Vector { v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z, v1.W + v2.W }
}

// Sub substracts 2 vectors
func (v1 Vector) Sub(v2 Vector) Vector {
	return Vector { v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z, v1.W - v2.W }
}

// Neg negates a vector
func (v Vector) Neg() Vector {
	return Vector { -v.X, -v.Y, -v.Z, -v.W }
}

// Mult multiplies a vector by a scale
func (v Vector) Mult(s float64) Vector {
	return Vector { v.X*s, v.Y*s, v.Z*s, v.W*s }
}

// Mult multiplies a vector by a scale
func (v Vector) Div(s float64) Vector {
	return Vector { v.X/s, v.Y/s, v.Z/s, v.W/s }
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z * v.Z + v.W *v.W)
}
func (v Vector) Normalize() Vector {
	return v.Div(v.Magnitude())
}

func (v1 Vector) Dot(v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z * v2.Z + v1.W *v2.W
}

func (a Vector) Cross(b Vector) Vector {
	return NewVector(a.Y * b.Z - a.Z * b.Y, 
					 a.Z * b.X - a.X * b.Z, 
					 a.X * b.Y - a.Y * b.X)
}


