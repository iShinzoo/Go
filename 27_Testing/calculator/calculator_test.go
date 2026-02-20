package calculator

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Expected %d, got %d", 3, result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 2)
	if result != 3 {
		t.Errorf("Expected %d, got %d", 3, result)
	}
}
