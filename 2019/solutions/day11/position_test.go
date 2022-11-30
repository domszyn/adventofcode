package day11

import "testing"

func TestPosition(t *testing.T) {
	p := Position{X: 0, Y: 0}

	if p = p.Up(); p.X != 0 || p.Y != -1 {
		t.Errorf("Expected {X:0, Y:-1}, got %v", p)
	}

	if p = p.Down(); p.X != 0 || p.Y != 0 {
		t.Errorf("Expected {X:0, Y:0}, got %v", p)
	}

	if p = p.Left(); p.X != -1 || p.Y != 0 {
		t.Errorf("Expected {X:-1, Y:0}, got %v", p)
	}

	if p = p.Right(); p.X != 0 || p.Y != 0 {
		t.Errorf("Expected {X:0, Y:0}, got %v", p)
	}
}
