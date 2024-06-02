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
	x, y        int
}

func (r *Rover) Location() string {
	return fmt.Sprintf("%s:%d:%d", r.orientation, r.x, r.y)
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
		r.y++
	case East:
		r.x++
	case South:
		r.y--
	case West:
		r.x--
	}
}
