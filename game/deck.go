package game

import (
	"carcassonne/tile"
	"fmt"
	"math/rand"
)

type Deck struct {
	tiles []*tile.Tile
	//CurrentIndex int
}

func NewDeck() *Deck {
	// TODO: FILL THE DECK
	//length := 12
	d := Deck{
		make([]*tile.Tile, 0),
	}
	return &d
}

func (d *Deck) Shuffle() {
	for i := 0; i < 1000; i++ {
		n1, n2 := rand.Intn(len(d.tiles)), rand.Intn(len(d.tiles))
		d.tiles[n1], d.tiles[n2] = d.tiles[n2], d.tiles[n1]
	}
}

func (d *Deck) removeTile(index int) {
	d.tiles = append(d.tiles[:index], d.tiles[index+1:]...)
}

func (d *Deck) DrawNext() *tile.Tile {
	t := d.tiles[0]
	d.removeTile(0)
	return t
}

func (d *Deck) DrawById(id string) *tile.Tile {
	i := 0
	for ; i < len(d.tiles); i++ {
		if d.tiles[i].Id == id {
			break
		}
	}

	t := d.tiles[i]
	d.removeTile(i)
	return t
}

func (d *Deck) Fill() {
	for k, _ := range tile.TileData {
		v := tile.TileData[k]
		fmt.Println(k, v)
		for i := 0; i < int(v.Quantity); i++ {
			d.tiles = append(d.tiles, tile.New(0, 0, k, v.Layout))
		}

	}

}
