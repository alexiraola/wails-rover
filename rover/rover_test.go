package rover

import "testing"

func TestCreatesRoverAtOrigin(t *testing.T) {
	rover := CreateRover()

	if rover.Location() != "N:0:0" {
		t.Fatalf("Expected %s, got %s", "N:0:0", rover.Location())
	}
}

func TestTurnsRight(t *testing.T) {
	rover := CreateRover()

	rover.Execute([]Command{Right})

	if rover.Location() != "E:0:0" {
		t.Fatalf("Expected %s, got %s", "E:0:0", rover.Location())
	}
}

func TestTurnsLeft(t *testing.T) {
	rover := CreateRover()

	rover.Execute([]Command{Left})

	if rover.Location() != "W:0:0" {
		t.Fatalf("Expected %s, got %s", "W:0:0", rover.Location())
	}
}

func TestMovesForward(t *testing.T) {
	rover := CreateRover()

	rover.Execute([]Command{Forward})

	if rover.Location() != "N:0:1" {
		t.Fatalf("Expected %s, got %s", "W:0:0", rover.Location())
	}
}

func TestTurnsRightTwice(t *testing.T) {
	rover := CreateRover()

	rover.Execute([]Command{Right, Right})

	if rover.Location() != "S:0:0" {
		t.Fatalf("Expected %s, got %s", "S:0:0", rover.Location())
	}
}

func TestMovesToTheEast(t *testing.T) {
	rover := CreateRover()

	rover.Execute([]Command{Right, Forward, Forward})

	if rover.Location() != "E:2:0" {
		t.Fatalf("Expected %s, got %s", "E:2:0", rover.Location())
	}
}
