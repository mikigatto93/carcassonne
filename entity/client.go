package entity

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Connection *websocket.Conn
	Output     chan Event
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn,
		make(chan Event),
	}
}

func (c *Client) ReadMessages(
	onError func(*Client),
	onMessage func(Event, *Client) error,
) {

	defer onError(c)

	for {
		_, payload, err := c.Connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("error reading message: ", err)
			}
			break
		}

		var req Event
		if err := json.Unmarshal(payload, &req); err != nil {
			log.Println("error marshalling message: ", err)
			break
		}

		//log.Println(req.Type, string(req.Payload))
		onMessage(req, c)

	}
}

func (c *Client) WriteMessages(onError func(*Client)) {
	defer onError(c)

	for {
		select {
		case message, ok := <-c.Output:

			if !ok {
				// cnannel closed, need to communicate it to the client
				if err := c.Connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					// Log that the connection is closed and the reason
					log.Println("connection closed: ", err)
				}

				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				log.Println("cannot send message: ", err)
				return
			}

			if err := c.Connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println(err)
			}
			log.Println("sent message")

		}
	}

}
