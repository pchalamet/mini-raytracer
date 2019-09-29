package foundation

import "testing"
import "math"


func TestNewVector(t *testing.T) {
	a := NewVector(4.3, -4.2, 3.1)
	if a.X != 4.3 || a.Y != -4.2 || a.Z != 3.1 || a.W != 0.0 {
		t.Error("NewVector built vector incorrectly")
	}
}

func TestNewPoint(t *testing.T) {
	a := NewPoint(4.3, -4.2, 3.1)
	if a.X != 4.3 || a.Y != -4.2 || a.Z != 3.1 || a.W != 1.0 {
		t.Error("NewPoint built vector incorrectly")
	}
}

func TestAddVector(t *testing.T) {
	// add 2 vectors
	v1 := Vector { 3, -2, 5, 1 }
	v2 := Vector { -2, 3, 1, 0 }
	r := v1.Add(v2)

	e := Vector { 1, 1, 6, 1 }
	if r != e {
		t.Error("Add is wrong")
	}

	// add a vector to a point
	p1 := NewPoint(3, 2, 1)
	v1 = NewVector(3, -2, 5)
	r = p1.Add(v1)
	e = NewPoint(6, 0, 6)
	if r != e {
		t.Error("Add is wrong")
	}
}

func TestSubVector(t *testing.T) {
	// sub 2 points
	p1 := NewPoint(3, 2, 1)
	p2 := NewPoint(5, 6, 7)
	r := p1.Sub(p2)
	e := NewVector(-2, -4, -6)
	if r != e {
		t.Error("Sub is wrong")
	}

	// sub a vector to a point
	v := NewPoint(5, 6, 7)
	p := NewVector(3, 2, 1)
	r = v.Sub(p)
	e = NewPoint(2, 4, 6)
	if r != e {
		t.Error("Sub is wrong")
	}

	// sub 2 vectors
	v1 := NewVector(3, 2, 1)
	v2 := NewVector(5, 6, 7)
	r = v1.Sub(v2)
	e = NewVector(-2, -4, -6)
	if r != e {
		t.Error("Sub is wrong")
	}
}


func TestNegVector(t *testing.T) {
	v := Vector {1, -2, 3, -4}
	r := v.Neg()
	e := Vector {-1, 2, -3, 4 }
	if r != e {
		t.Error("Neg is wrong")
	}
}

func TestMultVector(t *testing.T) {
	v := Vector { 1, -2, 3, -4 }
	r := v.Mult(3.5)
	e := Vector { 3.5, -7, 10.5, -14 }
	if r != e {
		t.Error("Mult is wrong")
	}
}

func TestDivVector(t *testing.T) {
	v := Vector { 1, -2, 3, -4 }
	r := v.Div(2)
	e := Vector { 0.5, -1, 1.5, -2 }
	if r != e {
		t.Error("Div is wrong")
	}
}

func TestMagnitudeVector(t *testing.T) {
	v := NewVector(1, 0, 0)
	if v.Magnitude() != 1 {
		t.Error("Magnitude is wrong")
	}

	v = NewVector(0, 1, 0)
	if v.Magnitude() != 1 {
		t.Error("Magnitude is wrong")
	}

	v = NewVector(0, 0, 1)
	if v.Magnitude() != 1 {
		t.Error("Magnitude is wrong")
	}

	v = NewVector(1, 2, 3)
	if ! Equals(v.Magnitude(), math.Sqrt(14)) {
		t.Error("Magnitude is wrong")
	}

	v = NewVector(-1, -2, -3)
	if ! Equals(v.Magnitude(), math.Sqrt(14)) {
		t.Error("Magnitude is wrong")
	}
}


func TestNormalizeVector(t *testing.T) {
	v := NewVector(4, 0, 0)
	r := v.Normalize()
	e := NewVector(1, 0, 0)
	if r != e {
		t.Error("Normalize is wrong")
	}

	v = NewVector(1, 2, 3)
	r = v.Normalize()
	e = NewVector(1 / math.Sqrt(14), 2 / math.Sqrt(14), 3/math.Sqrt(14))
	if ! (Equals(r.X, e.X) && Equals(r.Y, e.Y) && Equals(r.Z, e.Z) && Equals(r.W, e.W)) {
		t.Error("Normalize is wrong")
	}
}


func TestDotVector(t * testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)
	r := a.Dot(b)
	if r != 20 {
		t.Error("Dot is wrong")
	}
}

func TestCrossVector(t * testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)

	c1 := a.Cross(b)
	r1 := Vector {-1, 2, -1, 0}	
	if c1 != r1 {
		t.Error()
	}

	c2 := b.Cross(a)
	r2 := Vector {1, -2, 1, 0}
	if c2 != r2 {
		t.Error()
	}

}