package main

import "testing"

func TestCreateHand(t *testing.T) {
	testZestaw := createHand()

	if len(testZestaw) != 9 {
		t.Errorf("Expected 9 cards in hands, but got %v", len(testZestaw))
	}
}
