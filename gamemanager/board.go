package gamemanager

import (
	"carcassonne/meeple"
	"carcassonne/tile"
	"fmt"
)

type Board struct {
	tiles   map[string]*tile.Tile
	meeples map[string]*meeple.Meeple
}

func New() *Board {
	b := Board{}

	b.tiles = make(map[string]*tile.Tile)
	b.meeples = make(map[string]*meeple.Meeple)

	b.tiles["0;0"] = tile.New(0, 0, 0, tile.RiverStart)
	return &b
}

func (b *Board) PlaceTile(x int, y int, t *tile.Tile) bool {
	neighbors := b.getNeighbors(x, y)
	ok := true
	for _, v := range neighbors {
		res, _ := t.CanAttach(v)
		ok = ok && res
	}

	if ok {
		b.tiles[fmt.Sprintf("%d;%d", x, y)] = t
	}

	return ok
}

func (b *Board) PlaceMeeple() bool {

}

func (b *Board) getTile(x int, y int) (*tile.Tile, bool) {
	key := fmt.Sprintf("%d;%d", x, y)
	elem, ok := b.tiles[key]
	return elem, ok
}

func (b *Board) getNeighbors(x int, y int) [4]*tile.Tile {
	// 0: top, 1: right, 2: bottom, 3: left
	// nil if it does not exists
	arr := [4]*tile.Tile{}
	for i := 0; i < 4; i++ {
		if elem, err := b.getTile(x, y); err {
			arr[i] = elem
		}
	}
	return arr
}
