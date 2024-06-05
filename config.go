package main

import "wails_rover/rover"

type Config struct{}

func (p *Config) WorldSize() (uint8, uint8) {
	return rover.WORLD_WIDTH, rover.WORLD_HEIGHT
}
