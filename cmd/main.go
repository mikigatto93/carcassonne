package main

import (
	"carcassonne/server"
	"carcassonne/tile"
	"log"
)

func main() {

	tile.LoadAllTilesets()
	//fmt.Println(tile.TileLayouts)

	s := server.New("localhost", "8000")
	s.SetupRoutes()

	log.Println("Server started")
	err := s.StartServing()

	if err != nil {
		log.Fatal("An error occurred: ", err)
	}

}
