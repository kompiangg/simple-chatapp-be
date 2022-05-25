package websocket

import (
	"fmt"
	"log"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan *Message
}

func NewPool() *Pool {
	return &Pool{
		Register: make(chan *Client),
		Unregister: make(chan *Client),
		Clients: make(map[*Client]bool),
		Broadcast: make(chan *Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			log.Println("Size pool:", len(pool.Clients))
			for client := range pool.Clients {
				message := fmt.Sprintf("New User Joined...\n%d User on Group", len(pool.Clients))
				client.Conn.WriteJSON(message)
			}
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			log.Println("Size pool:", len(pool.Clients))
			for client := range pool.Clients {
				message := fmt.Sprintf("User Left...\n%d User on Group", len(pool.Clients))
				client.Conn.WriteJSON(message)
			}
		case message := <-pool.Broadcast:
			log.Println("Sending message to all user")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message.Payload); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}