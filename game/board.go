package game

import (
	"carcassonne/meeple"
	"carcassonne/tile"
	"fmt"
)

type Board struct {
	tiles             map[string]*tile.Tile
	meeples           []*meeple.Meeple
	possibleNeighbors map[string][]int8
}

func NewBoard() *Board {
	b := Board{
		tiles:             make(map[string]*tile.Tile),
		meeples:           make([]*meeple.Meeple, 24),
		possibleNeighbors: make(map[string][]int8),
	}
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
		t.SetPosition(x, y)
		b.tiles[fmt.Sprintf("%d;%d", x, y)] = t
		b.calculatePossibleNeighbors(t, x, y)
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

func (b *Board) getPossibleNeighborsBounds(x int, y int) ([]int8, bool) {
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

func (b *Board) updatePossibleNeighbors(x int, y int, pos int, t *tile.Tile) {
	oppositePos := (pos + 2) % 4
	if elem, err := b.getPossibleNeighborsBounds(x, y); err {
		elem[oppositePos] = t.GetBoundariesBitField()[pos]
	} else {
		boundBitField := []int8{0, 0, 0, 0}
		boundBitField[oppositePos] = t.GetBoundariesBitField()[pos]
		b.possibleNeighbors[fmt.Sprintf("%d;%d", x, y)] = boundBitField
	}

}

func (b *Board) calculatePossibleNeighbors(t *tile.Tile, x int, y int) {
	//remove the possible neighbor associated with the placed tile
	delete(b.possibleNeighbors, fmt.Sprintf("%d;%d", x, y))

	neighbors := b.getNeighbors(x, y)

	if neighbors[0] == nil {
		b.updatePossibleNeighbors(x, y+1, 0, t)
	}

	if neighbors[1] == nil {
		b.updatePossibleNeighbors(x+1, y, 1, t)
	}

	if neighbors[2] == nil {
		b.updatePossibleNeighbors(x, y-1, 2, t)
	}

	if neighbors[3] == nil {
		b.updatePossibleNeighbors(x-1, y, 3, t)
	}

}
