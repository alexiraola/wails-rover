package rover

import (
	"strings"
)

type Command string

const (
	Unknown Command = "U"
	Right   Command = "R"
	Left    Command = "L"
	Forward Command = "F"
)

type Rover struct {
	navigator Navigator
}

func CreateRover() Rover {
	return Rover{Navigator{North, Location{0, 0}}}
}

func ParseCommand(command string) Command {
	switch strings.TrimSpace(command) {
	case "R":
		return Right
	case "L":
		return Left
	case "F":
		return Forward
	}
	return Unknown
}

func (r *Rover) Location() string {
	return r.navigator.String()
}

func (r *Rover) Execute(commands []Command) {
	for _, command := range commands {
		switch command {
		case Right:
			r.navigator.turnRight()
		case Left:
			r.navigator.turnLeft()
		case Forward:
			r.navigator.moveForward()
		}
	}
}
