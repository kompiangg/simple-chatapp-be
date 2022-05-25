package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {return true},
}

func Upgrader(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return ws, nil
}