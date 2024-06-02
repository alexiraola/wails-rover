package rover

import "fmt"

type Orientation string

const (
	North Orientation = "N"
	East  Orientation = "E"
	West  Orientation = "W"
	South Orientation = "S"
)

type Command string

const (
	Right   Command = "R"
	Left    Command = "L"
	Forward Command = "F"
)

type Rover struct {
	orientation Orientation
	location    Location
}

func CreateRover() Rover {
	return Rover{North, Location{0, 0}}
}

func (r *Rover) Location() string {
	return fmt.Sprintf("%s:%s", r.orientation, r.location.String())
}

func (r *Rover) Execute(commands []Command) {
	for _, command := range commands {
		switch command {
		case Right:
			r.turnRight()
		case Left:
			r.turnLeft()
		case Forward:
			r.goForward()
		}
	}
}

func (r *Rover) turnRight() {
	switch r.orientation {
	case North:
		r.orientation = East
	case East:
		r.orientation = South
	case South:
		r.orientation = West
	case West:
		r.orientation = North
	}
}

func (r *Rover) turnLeft() {
	switch r.orientation {
	case North:
		r.orientation = West
	case East:
		r.orientation = North
	case South:
		r.orientation = East
	case West:
		r.orientation = South
	}
}

func (r *Rover) goForward() {
	switch r.orientation {
	case North:
		r.location = r.location.increaseY()
	case East:
		r.location = r.location.increaseX()
	case South:
		r.location = r.location.decreaseY()
	case West:
		r.location = r.location.decreaseX()
	}
}
