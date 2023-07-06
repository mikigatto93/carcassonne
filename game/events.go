package game

import (
	"carcassonne/entity"
	"fmt"
)

type GameEventHandler func(g *Game, event entity.Event, client *entity.Client) error

var GameEventHandlerMap map[string]GameEventHandler = map[string]GameEventHandler{
	"TEST": TestHandler,
}

func TestHandler(g *Game, event entity.Event, client *entity.Client) error {
	fmt.Println("TEST EVENT RECEIVED")
	return nil
}
