package main

type ObjectTypes int

// Constants representing the object types.
const (
	TANK ObjectTypes = 1 + iota
	BULLET
	WALL
	DESTRUCTIBLE_WALL
	BOUNDARY
	CLOSING_BOUNDARY
	POWERUP
)
