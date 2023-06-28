package game

import (
	"carcassonne/tile"
	"math/rand"
)

type Deck struct {
	tiles        []*tile.Tile
	CurrentIndex int
}

func NewDeck() *Deck {
	// TODO: FILL THE DECK
	length := 100
	d := Deck{
		make([]*tile.Tile, length),
		length - 1,
	}
	return &d
}

func (d *Deck) Shuffle() {
	for i := 0; i < 1000; i++ {
		n1, n2 := rand.Intn(100), rand.Intn(100)
		d.tiles[n1], d.tiles[n2] = d.tiles[n2], d.tiles[n1]
	}
}

func (d *Deck) DrawNext() *tile.Tile {
	t := d.tiles[d.CurrentIndex]
	d.CurrentIndex--
	return t
}
