package day11

import "testing"

func TestRotateLeft(t *testing.T) {
	d := Direction(0)

	if d = d.RotateLeft(); d != Direction(270) {
		t.Errorf("Expected 270, got %v", d)
	}

	if d = d.RotateLeft(); d != Direction(180) {
		t.Errorf("Expected 180, got %v", d)
	}

	if d = d.RotateLeft(); d != Direction(90) {
		t.Errorf("Expected 90, got %v", d)
	}

	if d = d.RotateLeft(); d != Direction(0) {
		t.Errorf("Expected 0, got %v", d)
	}
}

func TestRotateRight(t *testing.T) {
	d := Direction(0)

	if d = d.RotateRight(); d != Direction(90) {
		t.Errorf("Expected 270, got %v", d)
	}

	if d = d.RotateRight(); d != Direction(180) {
		t.Errorf("Expected 180, got %v", d)
	}

	if d = d.RotateRight(); d != Direction(270) {
		t.Errorf("Expected 90, got %v", d)
	}

	if d = d.RotateRight(); d != Direction(0) {
		t.Errorf("Expected 0, got %v", d)
	}
}
