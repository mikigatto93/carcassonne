package game

import (
	"carcassonne/entity"
	"carcassonne/tile"
	"encoding/json"
	"fmt"
)

type GameEventHandler func(g *Game, event entity.RequestEvent, client *entity.Client) error

type GameEventDescriptor struct {
}

var GameEventHandlerMap map[string]GameEventHandler = map[string]GameEventHandler{
	"TEST":          TestHandler,
	"NEW_PLAYER_ID": NewPlayerId,
	"PLAYER_READY":  PlayerReady,
	"START_GAME":    StartGame,
	"PASS_TURN":     PassTurn,
}

func TestHandler(g *Game, event entity.RequestEvent, client *entity.Client) error {
	fmt.Println("TEST EVENT RECEIVED")
	return nil
}

func NewPlayerId(g *Game, event entity.RequestEvent, client *entity.Client) error {
	var request struct {
		Name string `json:"player_name"`
		Id   string `json:"player_id"`
	}

	json.Unmarshal(event.Payload, &request)

	order := uint8(len(g.players))
	g.AddNewPlayer(
		entity.NewPlayer(request.Name, request.Id, order),
		client)

	isPlayerMaster := order == 0
	_, masterPlayer := g.GetPlayerByOrder(0)
	fmt.Println(order)
	fmt.Println(g.players)

	response := struct {
		IsMaster         bool   `json:"is_master"`
		MasterPlayerName string `json:"master_player_name"`
		PlayerId         string `json:"player_id"`
	}{
		IsMaster:         isPlayerMaster,
		MasterPlayerName: masterPlayer.Name,
		PlayerId:         masterPlayer.Id,
	}

	client.SendEvent(entity.ResponseEvent{
		Type: "PLAYER_CREATED", Payload: response},
	)
	return nil
}

func PlayerReady(g *Game, event entity.RequestEvent, client *entity.Client) error {
	player := g.players[client]
	player.Ready = true

	response := struct {
		PlayerName string `json:"player_name"`
		PlayerId   string `json:"player_id"`
	}{
		PlayerName: player.Name,
		PlayerId:   player.Id,
	}

	g.BroadcastEvent(entity.ResponseEvent{
		Type: "PLAYER_READY_OK", Payload: response},
	)
	return nil
}

func StartGame(g *Game, event entity.RequestEvent, client *entity.Client) error {

	g.PlaceStartingTile()

	startingTile, _ := g.board.getTile(0, 0)

	response := struct {
		StartingTile tile.Tile `json:"starting_tile"`
	}{
		StartingTile: *startingTile,
	}

	g.BroadcastEvent(entity.ResponseEvent{
		Type: "PLACE_STARTING_TILE", Payload: response},
	)

	PassTurn(g, event, client)
	return nil
}

func PassTurn(g *Game, event entity.RequestEvent, client *entity.Client) error {
	_, nextPlayer := g.GetNextTurnPlayer()
	nextTile := *g.deck.DrawNext()

	response := struct {
		PlayerName        string            `json:"player_name"`
		PlayerId          string            `json:"player_id"`
		NextTile          tile.Tile         `json:"next_tile"`
		TileBounds        []int8            `json:"tile_bounds"`
		PossibleNeighbors map[string][]int8 `json:"possible_neighbors"`
	}{
		PlayerName:        nextPlayer.Name,
		PlayerId:          nextPlayer.Id,
		NextTile:          nextTile,
		TileBounds:        nextTile.GetBoundariesBitField(),
		PossibleNeighbors: g.board.possibleNeighbors,
	}

	g.BroadcastEvent(entity.ResponseEvent{
		Type: "NEXT_TURN", Payload: response},
	)
	return nil
}
