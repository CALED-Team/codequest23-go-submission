package main

type ObjectTypes int

const (
	TANK ObjectTypes = 1 + iota
	BULLET
	WALL
	DESTRUCTIBLE_WALL
	BOUNDARY
	CLOSING_BOUNDARY
	POWERUP
)
