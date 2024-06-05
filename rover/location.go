package rover

import "fmt"

const (
	WORLD_WIDTH  = 5
	WORLD_HEIGHT = 10
)

type Location struct {
	x int
	y int
}

func (l *Location) String() string {
	return fmt.Sprintf("%d:%d", l.x, l.y)
}

func (l *Location) increaseX() Location {
	return Location{(l.x + 1) % WORLD_WIDTH, l.y}
}

func (l *Location) decreaseX() Location {
	return Location{(WORLD_WIDTH + l.x - 1) % WORLD_WIDTH, l.y}
}

func (l *Location) increaseY() Location {
	return Location{l.x, (l.y + 1) % WORLD_HEIGHT}
}

func (l *Location) decreaseY() Location {
	return Location{l.x, (WORLD_HEIGHT + l.y - 1) % WORLD_HEIGHT}
}
