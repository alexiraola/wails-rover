package rover

import "testing"

func TestNavigatorCreates(t *testing.T) {
	navigator := Navigator{North, Location{0, 0}}

	if navigator.String() != "N:0:0" {
		t.Fatalf("Expected %s, got %s", "N:0:0", navigator.String())
	}
}

func TestNavigatorTurnsRight(t *testing.T) {
	navigator := Navigator{North, Location{0, 0}}

	navigator.turnRight()

	if navigator.String() != "E:0:0" {
		t.Fatalf("Expected %s, got %s", "E:0:0", navigator.String())
	}
}

func TestNavigatorTurnsLeft(t *testing.T) {
	navigator := Navigator{North, Location{0, 0}}

	navigator.turnLeft()

	if navigator.String() != "W:0:0" {
		t.Fatalf("Expected %s, got %s", "W:0:0", navigator.String())
	}
}

func TestNavigatorMovesForward(t *testing.T) {
	navigator := Navigator{North, Location{0, 0}}

	navigator.moveForward()

	if navigator.String() != "N:0:1" {
		t.Fatalf("Expected %s, got %s", "N:0:1", navigator.String())
	}
}

func TestNavigatorMovesRight(t *testing.T) {
	navigator := Navigator{North, Location{0, 0}}

	navigator.turnRight()
	navigator.moveForward()

	if navigator.String() != "E:1:0" {
		t.Fatalf("Expected %s, got %s", "E:1:0", navigator.String())
	}
}

func TestNavigatorRoundsRightEdge(t *testing.T) {
	navigator := Navigator{North, Location{0, 9}}

	navigator.moveForward()

	if navigator.String() != "N:0:0" {
		t.Fatalf("Expected %s, got %s", "N:0:0", navigator.String())
	}
}

func TestNavigatorRoundsLeftEdge(t *testing.T) {
	navigator := Navigator{West, Location{0, 0}}

	navigator.moveForward()

	if navigator.String() != "W:9:0" {
		t.Fatalf("Expected %s, got %s", "W:9:0", navigator.String())
	}
}

func TestNavigatorRoundsTopEdge(t *testing.T) {
	navigator := Navigator{North, Location{0, 9}}

	navigator.moveForward()

	if navigator.String() != "N:0:0" {
		t.Fatalf("Expected %s, got %s", "N:0:0", navigator.String())
	}
}

func TestNavigatorRoundsBottomEdge(t *testing.T) {
	navigator := Navigator{South, Location{0, 0}}

	navigator.moveForward()

	if navigator.String() != "S:0:9" {
		t.Fatalf("Expected %s, got %s", "S:0:9", navigator.String())
	}
}

func TestNavigatorTurnsRightManyTimes(t *testing.T) {
	navigator := Navigator{North, Location{0, 0}}

	navigator.turnRight()
	navigator.turnRight()
	navigator.turnRight()

	if navigator.String() != "W:0:0" {
		t.Fatalf("Expected %s, got %s", "W:0:0", navigator.String())
	}
}

func TestNavigatorTurnsLeftManyTimes(t *testing.T) {
	navigator := Navigator{North, Location{0, 0}}

	navigator.turnLeft()
	navigator.turnLeft()
	navigator.turnLeft()

	if navigator.String() != "E:0:0" {
		t.Fatalf("Expected %s, got %s", "E:0:0", navigator.String())
	}
}
