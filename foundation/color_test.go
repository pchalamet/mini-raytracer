package foundation

import "testing"

func TestAddColor(t * testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	r := c1.Add(c2)
	e := NewColor(1.6, 0.7, 1.0)
	if ! (Equals(r.Red, e.Red) && Equals(r.Green, e.Green) && Equals(r.Blue, e.Blue)) {
		t.Error("Add is wrong")
	}
}

func TestSubColor(t * testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	r := c1.Sub(c2)
	e := NewColor(0.2, 0.5, 0.5)
	if ! (Equals(r.Red, e.Red) && Equals(r.Green, e.Green) && Equals(r.Blue, e.Blue)) {
		t.Error("Sub is wrong")
	}
}

