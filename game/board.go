package game

import (
	"carcassonne/meeple"
	"carcassonne/tile"
	"fmt"
)

type Board struct {
	tiles             map[string]*tile.Tile
	meeples           []*meeple.Meeple
	possibleNeighbors map[string][]byte
}

func NewBoard() *Board {
	b := Board{}

	b.tiles = make(map[string]*tile.Tile)
	b.meeples = make([]*meeple.Meeple, 24)

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

func (b *Board) PlaceMeeple(m *meeple.Meeple) bool {
	b.meeples = append(b.meeples, m)
	return true
}

func (b *Board) getTile(x int, y int) (*tile.Tile, bool) {
	key := fmt.Sprintf("%d;%d", x, y)
	elem, ok := b.tiles[key]
	return elem, ok
}

func (b *Board) getPossibleNeighborsBounds(x int, y int) ([]byte, bool) {
	key := fmt.Sprintf("%d;%d", x, y)
	elem, ok := b.possibleNeighbors[key]
	return elem, ok
}

func (b *Board) getNeighbors(x int, y int) [4]*tile.Tile {
	// 0: top, 1: right, 2: bottom, 3: left
	// nil if it does not exists
	arr := [4]*tile.Tile{}

	if elem, err := b.getTile(x, y+1); err {
		arr[0] = elem
	}

	if elem, err := b.getTile(x+1, y); err {
		arr[1] = elem
	}

	if elem, err := b.getTile(x, y-1); err {
		arr[2] = elem
	}

	if elem, err := b.getTile(x-1, y); err {
		arr[3] = elem
	}

	return arr
}

func (b *Board) updatePossibleNeighbors(x int, y int, pos int, tile *tile.Tile) {
	if elem, err := b.getPossibleNeighborsBounds(x, y); err {
		elem[(pos+2)%4] = tile.GetBoundariesBitField()[pos]
	} else {
		boundBitField := []byte{0, 0, 0, 0}
		boundBitField[(pos+2)%4] = tile.GetBoundariesBitField()[pos]
		b.possibleNeighbors[fmt.Sprintf("%d;%d", x, y)] = boundBitField
	}

}

func (b *Board) CalculatePossibleNeighbors(tile *tile.Tile, x int, y int) {
	//remove the possible neighbor associated with the placed tile
	delete(b.possibleNeighbors, fmt.Sprintf("%d;%d", x, y))

	neighbors := b.getNeighbors(x, y)

	if neighbors[0] == nil {
		b.updatePossibleNeighbors(x, y+1, 0, tile)
	}

	if neighbors[1] == nil {
		b.updatePossibleNeighbors(x+1, y, 1, tile)
	}

	if neighbors[2] == nil {
		b.updatePossibleNeighbors(x, y-1, 2, tile)
	}

	if neighbors[3] == nil {
		b.updatePossibleNeighbors(x-1, y, 3, tile)
	}

}
