package board

import (
	"testing"
)

func TestControlPositionRight(t *testing.T) {
	b := Board{}
	b.Init()

	b.ControlPosition("right")

	if b.CurrentX != 1 {
		t.Errorf("Expected 1, got %d", b.CurrentX)
	}
}

func TestControlPositionDown(t *testing.T) {
	b := Board{}
	b.Init()

	b.ControlPosition("down")

	if b.CurrentY != 1 {
		t.Errorf("Expected 1, got %d", b.CurrentY)
	}
}

func TestMakePreValues(t *testing.T) {
	b := Board{}
	b.Init()

	b.MakePreValues()

	if b.PreValues[0][0] != "@" {
		t.Errorf("Expected @, got %s", b.PreValues[0][0])
	}
}

func TestFlip(t *testing.T) {
	b := Board{}
	b.Init()

	b.Flip()

	if b.CurrentPattern != "X" {
		t.Errorf("Expected X, got %s", b.CurrentPattern)
	}
}

func TestCheckHasWinnerTrue(t *testing.T) {
	b := Board{}
	b.Init()

	b.Flip()
	b.ControlPosition("right")
	b.Flip()
	b.ControlPosition("right")
	b.Flip()

	if !b.CheckHasWinner() {
		t.Error("Expected true, got false")
	}
}

func TestCheckHasWinnerFalse(t *testing.T) {
	b := Board{}
	b.Init()

	b.Flip()
	b.ControlPosition("right")
	b.Flip()
	b.ControlPosition("right")

	if b.CheckHasWinner() {
		t.Error("Expected false, got true")
	}
}
