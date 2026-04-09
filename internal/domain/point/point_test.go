package point

import "testing"

func TestNewPoint(t *testing.T) {
	t.Parallel()
	p := NewPoint(5, 10)
	if p.X() != 5 {
		t.Errorf("Expected X=5, got %d", p.X())
	}
	if p.Y() != 10 {
		t.Errorf("Expected Y=10, got %d", p.Y())
	}
}

func TestPointZeroValues(t *testing.T) {
	t.Parallel()
	p := NewPoint(0, 0)
	if p.X() != 0 || p.Y() != 0 {
		t.Errorf("Expected (0,0), got (%d,%d)", p.X(), p.Y())
	}
}

func TestPointEquality(t *testing.T) {
	t.Parallel()
	p1 := NewPoint(3, 7)
	p2 := NewPoint(3, 7)
	p3 := NewPoint(3, 8)

	if p1 != p2 {
		t.Error("Points with same coordinates should be equal")
	}
	if p1 == p3 {
		t.Error("Points with different coordinates should not be equal")
	}
}
