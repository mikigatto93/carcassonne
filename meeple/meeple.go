package meeple

import (
	"carcassonne/tile"
)

type Role int

const (
	Thief Role = iota
	Knight
	Farmer
	Monk
)

type tilePosition struct {
	x int
	y int
}

type Meeple struct {
	tileRif *tile.Tile
	tilePos tilePosition
	role    Role
	points  int
}

func New(role Role) Meeple {

	return Meeple{
		nil,
		tilePosition{-1, -1},
		role,
		1,
	}
}
