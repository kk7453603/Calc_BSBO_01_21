package fasulyan

import (
	"testing"
)

func TestDivision(t *testing.T) {

	if Division(3, 2) != 1.5 {
		t.Errorf("Expected 1.5, got %f", Division(3, 2))
	}
}

func TestModule(t *testing.T) {

	if Module(3, 2) != 1 {
		t.Errorf("Expected 1, got %f", Module(3, 2))
	}
}

func TestSin(t *testing.T) {

	if Sin(3) != 0.1411200080598672 {
		t.Errorf("Expected 0.1411200080598672, got %f", Sin(3))
	}
}
