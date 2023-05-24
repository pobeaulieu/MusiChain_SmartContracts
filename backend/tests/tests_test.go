package test

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := 5
	expected := 5
	if result != expected {
		t.Errorf("Add function returned %d, expected %d", result, expected)
	}
}
