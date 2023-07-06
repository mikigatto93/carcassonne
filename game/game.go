package game

import (
	"carcassonne/entity"
)

const MaxNumOfPlayers = 6

type Game struct {
	board   *Board
	deck    *Deck
	turn    int
	players map[*entity.Client]*entity.Player
}

func New() *Game {
	g := Game{
		NewBoard(),
		NewDeck(),
		0,
		make(map[*entity.Client]*entity.Player, MaxNumOfPlayers),
	}

	return &g
}

func (g *Game) getPlayerById(id string) *entity.Player {
	for _, v := range g.players {
		if v.Id == id {
			return v
		}
	}
	return nil
}

func (g *Game) AddNewPlayer(player *entity.Player, client *entity.Client) {
	if len(g.players) < MaxNumOfPlayers {
		g.players[client] = player
	}
}

func (g *Game) RemovePlayerById(id string) {
	//TODO
}

func (g *Game) RemovePlayerByClient(client *entity.Client) {
	// Client connection is not closed
	delete(g.players, client)
}
