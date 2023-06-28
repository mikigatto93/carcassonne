package game

import "carcassonne/tile"

const MaxNumOfPlayers = 6

type Game struct {
	board   *Board
	deck    *Deck
	players []*Player
	turn    int
}

func New() *Game {
	g := Game{
		NewBoard(),
		NewDeck(),
		make([]*Player, MaxNumOfPlayers),
		0,
	}

	return &g
}

func (g *Game) getPlayerById(id string) (int, *Player) {
	for i, v := range g.players {
		if v.id == id {
			return i, v
		}
	}
	return -1, nil
}

func (g *Game) AddNewPlayer(player *Player) {
	if len(g.players) < MaxNumOfPlayers {
		g.players = append(g.players, player)
	}
}

func (g *Game) RemovePlayer(id string) {
	if i, _ := g.getPlayerById(id); i != -1 {
		g.players = append(g.players[:i], g.players[i+1:]...)
	}
}

func (g *Game) NextTurn() (*Player, *tile.Tile) {

}
