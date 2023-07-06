package main

import (
	"carcassonne/server"
	"log"
)

func main() {

	s := server.New("localhost", "8000")
	s.SetupRoutes()

	log.Println("Server started")
	err := s.StartServing()

	if err != nil {
		log.Fatal("An error occurred: ", err)
	}
}
