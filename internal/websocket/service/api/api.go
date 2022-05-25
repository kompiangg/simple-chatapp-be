package api

import "github.com/gorilla/websocket"

type WebSocketServiceAPI interface {
	CreateClient(conn *websocket.Conn)
}