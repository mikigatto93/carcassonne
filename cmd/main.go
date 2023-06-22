package main

import (
	"carcassonne/tile"
	"fmt"
)

func main() {
	tile1 := tile.New(0, 0, 0, tile.Base)

	tile2 := tile.New(0, -1, 1, tile.Base)

	tile1.Rotate(1)
	tile2.Rotate(3)

	fmt.Println(tile1)

	fmt.Println(tile2)

	res, _ := tile1.CanAttach(tile2)

	fmt.Println(res)

}
