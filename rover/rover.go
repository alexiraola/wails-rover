package rover

type Command string

const (
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
