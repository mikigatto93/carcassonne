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

	g.deck.Fill()

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

func (g *Game) GetPlayerByOrder(order uint8) (*entity.Client, *entity.Player) {
	for client, player := range g.players {
		if player.Order == order {
			return client, player
		}
	}
	return nil, nil
}

func (g *Game) PlaceStartingTile() {
	t := g.deck.DrawById("River1-0")
	g.board.PlaceTile(0, 0, t)
}

func (g *Game) BroadcastEvent(event entity.ResponseEvent) {
	for client, _ := range g.players {
		client.SendEvent(event)
	}
}

func (g *Game) GetNextTurnPlayer() (*entity.Client, *entity.Player) {
	client, player := g.GetPlayerByOrder(uint8(g.turn % len(g.players)))
	g.turn++
	return client, player
}
