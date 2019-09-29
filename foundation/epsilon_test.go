package foundation

import "testing"

func TestInequality(t *testing.T) {
	a := 0.0
	b := a + epsilon

	if Equals(a, b) {
		t.Errorf("Number should not be equal")
	}

	if Equals(b, a) {
		t.Errorf("Number should not be equal")
	}
}


func TestEqualitySame(t *testing.T) {
	a := 0.0
	b := a

	if ! Equals(a, b) {
		t.Errorf("Number should be equal")
	}

	if ! Equals(b, a) {
		t.Errorf("Number should be equal")
	}
}

func TestEqualityLessThanEpsilon(t *testing.T) {
	a := 0.0
	b := a - (epsilon / 2.0)

	if ! Equals(a, b) {
		t.Errorf("Number should be equal")
	}

	if ! Equals(b, a) {
		t.Errorf("Number should be equal")
	}
}
