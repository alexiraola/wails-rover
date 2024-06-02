package rover

import "testing"

func TestCreatesLocation(t *testing.T) {
	location := Location{0, 0}

	if location.String() != "0:0" {
		t.Fatalf("Expected %s, got %s", "0:0", location.String())
	}
}

func TestIncreasesX(t *testing.T) {
	location := Location{0, 0}
	newLocation := location.increaseX()

	if newLocation.String() != "1:0" {
		t.Fatalf("Expected %s, got %s", "1:0", newLocation.String())
	}
}

func TestIncreasesY(t *testing.T) {
	location := Location{0, 0}
	newLocation := location.increaseY()

	if newLocation.String() != "0:1" {
		t.Fatalf("Expected %s, got %s", "0:1", newLocation.String())
	}
}

func TestDecreasesX(t *testing.T) {
	location := Location{1, 0}
	newLocation := location.decreaseX()

	if newLocation.String() != "0:0" {
		t.Fatalf("Expected %s, got %s", "0:0", newLocation.String())
	}
}

func TestDecreasesY(t *testing.T) {
	location := Location{0, 1}
	newLocation := location.decreaseY()

	if newLocation.String() != "0:0" {
		t.Fatalf("Expected %s, got %s", "0:0", newLocation.String())
	}
}

func TestRoundsXOnLeftEdge(t *testing.T) {
	location := Location{0, 0}
	newLocation := location.decreaseX()

	if newLocation.String() != "9:0" {
		t.Fatalf("Expected %s, got %s", "9:0", newLocation.String())
	}
}

func TestRoundsXOnRightEdge(t *testing.T) {
	location := Location{9, 0}
	newLocation := location.increaseX()

	if newLocation.String() != "0:0" {
		t.Fatalf("Expected %s, got %s", "0:0", newLocation.String())
	}
}

func TestRoundsYOnBottomEdge(t *testing.T) {
	location := Location{0, 0}
	newLocation := location.decreaseY()

	if newLocation.String() != "0:9" {
		t.Fatalf("Expected %s, got %s", "0:9", newLocation.String())
	}
}

func TestRoundsYOnTopEdge(t *testing.T) {
	location := Location{0, 9}
	newLocation := location.increaseY()

	if newLocation.String() != "0:0" {
		t.Fatalf("Expected %s, got %s", "0:0", newLocation.String())
	}
}
