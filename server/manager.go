package server

import (
	"carcassonne/entity"
	"carcassonne/game"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	Clients map[*entity.Client]bool
	*sync.RWMutex
	game *game.Game
}

func NewManager() *Manager {
	return &Manager{
		make(map[*entity.Client]bool),
		&sync.RWMutex{},
		game.New(),
	}
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	// Begin by upgrading the HTTP request
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := entity.NewClient(conn)

	log.Println("New connection2")

	m.addClient(client)
	go client.ReadMessages(m.RemoveClient, m.RouteEvent)
	go client.WriteMessages(m.RemoveClient)
}

func (m *Manager) addClient(client *entity.Client) {
	m.Lock()
	defer m.Unlock()

	m.Clients[client] = true
}

func (m *Manager) RemoveClient(client *entity.Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.Clients[client]; ok {
		client.Connection.Close()
		delete(m.Clients, client)
	}
}

func (m *Manager) RouteEvent(event entity.Event, client *entity.Client) error {
	if handler, ok := game.GameEventHandlerMap[event.Type]; ok {
		if err := handler(m.game, event, client); err != nil {
			log.Printf("Error in event handler %s: %v", event.Type, err)
		}
		return nil
	} else {
		return fmt.Errorf("event with key %s is not supported", event.Type)
	}
}
