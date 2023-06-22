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

func New(tileRif *tile.Tile, role Role,
	tilePosX int, tilePosY int, points int) Meeple {

	return Meeple{
		tileRif,
		tilePosition{tilePosX, tilePosY},
		role,
		points,
	}
}
