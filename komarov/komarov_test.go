package komarov

import (
	"testing"
)

func TestCos(t *testing.T) {
	res := Cosinus(0)
	if int(res) != 1 {
		t.Errorf(" Cos(0)=1, your result %d", int(res))
	}
}

func TestSquare(t *testing.T) {
	res := Square(3.0)
	if res != 9.0 {
		t.Errorf("Square(3)=9, your result %f", res)
	}
}
