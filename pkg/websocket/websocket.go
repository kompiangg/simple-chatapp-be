package websocket

import (
	"io"
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

func Reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func Writer(conn *websocket.Conn) {
	for {
		log.Println("Sending...")
		messageType, r, err := conn.NextReader()
		if err != nil {
			log.Println(err)
			return
		}

		w, err := conn.NextWriter(messageType)
		if err != nil {
			log.Println(err)
			return
		}

		if _, err := io.Copy(w, r); err != nil {
			log.Println(err)
			return
		}

		if err := w.Close(); err != nil {
			log.Println(err)
			return
		}
	}
}