package negodin

import (
	"testing"
)

func TestMod(t *testing.T) {
	if Mod(5, 0) != 0 {
		t.Errorf("Expected 1, got %d", Mod(5, 2))
	}
}

func TestLeft_Shift(t *testing.T) {
	if Left_Shift(5, 2) != 20 {
		t.Errorf("Expected 20, got %d", Left_Shift(5, 2))
	}
}

func TestRight_Shift(t *testing.T) {
	res := Right_Shift(5, 2)
	if res != 1 {
		t.Errorf("Expected 0, got %d", res)
	}
}
