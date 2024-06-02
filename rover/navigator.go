package rover

import "fmt"

type Orientation string

const (
	North Orientation = "N"
	East  Orientation = "E"
	West  Orientation = "W"
	South Orientation = "S"
)

type Navigator struct {
	orientation Orientation
	location    Location
}

func (n *Navigator) String() string {
	return fmt.Sprintf("%s:%s", n.orientation, n.location.String())
}

func (n *Navigator) turnRight() {
	switch n.orientation {
	case North:
		n.orientation = East
	case East:
		n.orientation = South
	case South:
		n.orientation = West
	case West:
		n.orientation = North
	}
}

func (n *Navigator) turnLeft() {
	switch n.orientation {
	case North:
		n.orientation = West
	case East:
		n.orientation = North
	case South:
		n.orientation = East
	case West:
		n.orientation = South
	}
}

func (n *Navigator) moveForward() {
	switch n.orientation {
	case North:
		n.location = n.location.increaseY()
	case East:
		n.location = n.location.increaseX()
	case South:
		n.location = n.location.decreaseY()
	case West:
		n.location = n.location.decreaseX()
	}
}
